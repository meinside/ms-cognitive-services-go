package cognitive

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Emotion struct {
	FaceRectangle Rectangle          `json:"faceRectangle"`
	Scores        map[string]float64 `json:"scores"`
}

// https://docs.microsoft.com/en-us/azure/cognitive-services/emotion/emotion-api-how-to-topics/howtocallemotionforvideo
type EmotionProcessingResult struct {
	Version   int   `json:"version"` // version = 1
	Timescale int64 `json:"timescale"`
	Offset    int64 `json:"offset"`
	Framerate int   `json:"framerate"`
	Width     int   `json:"width"`
	Height    int   `json:"height"`
	Fragments []struct {
		Start    int64 `json:"start"`
		Duration int64 `json:"duration"`
		Interval int64 `json:"interval"`
		Events   [][]struct {
			WindowFaceDistribution map[string]float32 `json:"windowFaceDistribution"`
			WindowMeanScores       map[string]float32 `json:"windowMeanScores"`
		} `json:"events"`
	} `json:"fragments"`
}

// Emotion API: Emotion Recognition
//
// https://westus.dev.cognitive.microsoft.com/docs/services/5639d931ca73072154c1ce89/operations/563b31ea778daf121cc3a5fa
// https://westus.dev.cognitive.microsoft.com/docs/services/5639d931ca73072154c1ce89/operations/56f23eb019845524ec61c4d7
//
// key   : subscription key for this API
// image : string(image url) or []byte(image bytes array)
// rects : rectangles of faces (can be nil if none)
func EmotionRecognizeImage(key string, image interface{}, rects []Rectangle) (emotions []Emotion, err error) {
	const apiUrl = "https://westus.api.cognitive.microsoft.com/emotion/v1.0/recognize"

	var result []byte

	var params map[string]string = nil
	if len(rects) > 0 {
		faceRects := []string{}
		for _, rect := range rects {
			faceRects = append(faceRects, fmt.Sprintf("%d,%d,%d,%d", rect.Left, rect.Top, rect.Width, rect.Height))
		}

		params = map[string]string{"faceRectangles": strings.Join(faceRects, ";")}
	}

	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &emotions); err == nil {
			return emotions, nil
		}
	}
	return []Emotion{}, err
}

// Emotion API: Emotion Recognition in Video
//
// https://westus.dev.cognitive.microsoft.com/docs/services/5639d931ca73072154c1ce89/operations/56f8d40e1984551ec0a0984e
//
// key              : subscription key for this API
// video            : string(video url) or []byte(video bytes array)
// outputStyle      : "aggregate" (default) or "perFrame"
// progressNotifier : can be nil
func EmotionRecognizeVideo(
	key string,
	video interface{},
	outputStyle string,
	progressNotifier func(status string, progress float32),
) (processResult EmotionProcessingResult, err error) {
	const apiUrl = "https://westus.api.cognitive.microsoft.com/emotion/v1.0/recognizeinvideo"

	var result []byte

	// params
	params := map[string]string{}
	if outputStyle != "" {
		params["outputStyle"] = outputStyle // "aggregate" (default) or "perFrame"
	}

	result, err = postArg(apiUrl, key, params, video)

	// get recognition in video operation result
	//
	// https://westus.dev.cognitive.microsoft.com/docs/services/5639d931ca73072154c1ce89/operations/56f8d4471984551ec0a0984f
	if err == nil {
		opLocation := string(result)

		if IsVerbose {
			log.Printf("<< Checking operation location...: %s", opLocation)
		}

		var lastProgress float32 = -1.0

		for i := 0; i < NumTries; i++ {
			time.Sleep(WaitSeconds * time.Second)

			var status OperationStatus
			if result, err = httpGet(opLocation, key, nil); err == nil {
				if err = json.Unmarshal(result, &status); err == nil {
					if status.Status == "Succeeded" {
						if err = json.Unmarshal([]byte(status.ProcessingResultJson), &processResult); err == nil {
							return processResult, nil
						} else {
							break
						}
					} else if status.Status == "Failed" {
						err = fmt.Errorf(status.Message)
						break
					} else {
						if progressNotifier != nil && lastProgress != status.Progress {
							progressNotifier(status.Status, status.Progress)
						}

						if IsVerbose {
							log.Printf(">> %s (%.2f%%)", status.Status, status.Progress)
						}

						lastProgress = status.Progress
					}
				}
			}

			if i+1 == NumTries {
				err = fmt.Errorf("Reached the limit of tries: %d", NumTries)
			}
		}
	}

	return EmotionProcessingResult{}, err
}
