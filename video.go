package cognitive

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type VideoProcessingResult1 struct {
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
			X      float32 `json:"x"`
			Y      float32 `json:"y"`
			Width  float32 `json:"width"`
			Height float32 `json:"height"`
		} `json:"events"`
	} `json:"fragments"`
}

type VideoProcessingResult2 struct {
	Version   int   `json:"version"` // version = 2
	Timescale int64 `json:"timescale"`
	Offset    int64 `json:"offset"`
	Framerate int   `json:"framerate"`
	Rotation  int   `json:"rotation"`
	Width     int   `json:"width"`
	Height    int   `json:"height"`
	Regions   []struct {
		Id     int     `json:"id"`
		Type   string  `json:"type"`
		Points []Point `json:"points"`
	} `json:"regions"`
	Fragments []struct {
		Start    int64 `json:"start"`
		Duration int64 `json:"duration"`
		Interval int64 `json:"interval"`
		Events   [][]struct {
			Type      int    `json:"type"`
			TypeName  string `json:"typeName"`
			Locations []struct {
				X      float32 `json:"x"`
				Y      float32 `json:"y"`
				Width  float32 `json:"width"`
				Height float32 `json:"height"`
			} `json:"locations"`
			RegionId int `json:"regionId"`
		} `json:"events"`
	} `json:"fragments"`
}

// Video API: Face Detection and Tracking
//
// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/565d6517778daf0978c45e39
//
// key              : subscription key for this API
// video            : string(video url) or []byte(video bytes array)
// progressNotifier : can be nil
func VideoFaceDetectTrack(
	key string,
	video interface{},
	progressNotifier func(status string, progress float32),
) (processResult VideoProcessingResult1, err error) {
	const apiUrl = "https://westus.api.cognitive.microsoft.com/video/v1.0/trackface"

	var result []byte
	result, err = postArg(apiUrl, key, nil, video)

	// get video operation result
	//
	// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/565d6517778daf0978c45e36
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
						if status.ProcessingResultJson != "" {
							if err = json.Unmarshal([]byte(status.ProcessingResultJson), &processResult); err == nil {
								return processResult, nil
							} else {
								break
							}
						} else {
							err = fmt.Errorf("processingResult is empty")
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

	return VideoProcessingResult1{}, err
}

// Video API: Motion Detection
//
// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/565d6517778daf0978c45e3a
//
// key                : subscription key for this API
// video              : string(video url) or []byte(video bytes array)
// sensitivityLevel   : "low", "medium", or "high" (default: "medium")
// frameSamplingValue : 1 ~ 20 (default: 1)
// detectionZones     : can be nil
// detectLightChange  : default false
// mergeTimeThreshold : 0.0 ~ 10.0 (default: 0.0)
// progressNotifier   : can be nil
func VideoMotionDetect(
	key string,
	video interface{},
	sensitivityLevel string,
	frameSamplingValue int,
	detectionZones [][]Point,
	detectLightChange bool,
	mergeTimeThreshold float32,
	progressNotifier func(status string, progress float32),
) (processResult VideoProcessingResult2, err error) {
	const apiUrl = "https://westus.api.cognitive.microsoft.com/video/v1.0/detectmotion"

	// params
	params := map[string]string{}
	if sensitivityLevel != "" {
		params["sensitivityLevel"] = sensitivityLevel
	}
	if frameSamplingValue >= 1 || frameSamplingValue <= 20 {
		params["frameSampleValue"] = strconv.Itoa(frameSamplingValue)
	}
	if len(detectionZones) > 0 {
		zones := []string{}
		for _, zone := range detectionZones {
			points := []string{}
			for _, point := range zone {
				points = append(points, fmt.Sprintf("%.3f,%.3f", point.X, point.Y))
			}

			zones = append(zones, strings.Join(points, ";"))
		}
		params["detectionZones"] = strings.Join(zones, "|")
	}
	if detectLightChange {
		params["detectLightChange"] = "true"
	}
	if mergeTimeThreshold >= 0.0 && mergeTimeThreshold <= 10.0 {
		params["mergeTimeThreshold"] = fmt.Sprintf("%.3f", mergeTimeThreshold)
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, video)

	// get video operation result
	//
	// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/565d6517778daf0978c45e36
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
						if status.ProcessingResultJson != "" {
							if err = json.Unmarshal([]byte(status.ProcessingResultJson), &processResult); err == nil {
								return processResult, nil
							} else {
								break
							}
						} else {
							err = fmt.Errorf("processingResult is empty")
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

	return VideoProcessingResult2{}, err
}

// Video API: Stabilization
//
// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/565d6517778daf0978c45e35
//
// key              : subscription key for this API
// video            : string(video url) or []byte(video bytes array)
// progressNotifier : can be nil
func VideoStabilize(
	key string,
	video interface{},
	progressNotifier func(status string, progress float32),
) (fileUrl string, err error) {
	const apiUrl = "https://westus.api.cognitive.microsoft.com/video/v1.0/stabilize"

	var result []byte
	result, err = postArg(apiUrl, key, nil, video)

	// get video operation result
	//
	// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/565d6517778daf0978c45e36
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
						if status.ResourceLocation != "" {
							return status.ResourceLocation, nil
						} else {
							err = fmt.Errorf("resourceLocation is empty")
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

	return "", err
}

// Video API: Thumbnail
//
// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/56f8acb0778daf23d8ec6738
//
// key                              : subscription key for this API
// video                            : string(video url) or []byte(video bytes array)
// maxMotionThumbnailDurationInSecs : default 0
// outputAudio                      : default true
// fadeInFadeOut                    : default true
// progressNotifier                 : can be nil
func VideoThumbnail(
	key string,
	video interface{},
	maxMotionThumbnailDurationInSecs int,
	outputAudio bool,
	fadeInFadeOut bool,
	progressNotifier func(status string, progress float32),
) (fileUrl string, err error) {
	const apiUrl = "https://westus.api.cognitive.microsoft.com/video/v1.0/generatethumbnail"

	// params
	params := map[string]string{}
	if maxMotionThumbnailDurationInSecs > 0 {
		params["maxMotionThumbnailDurationInSecs"] = strconv.Itoa(maxMotionThumbnailDurationInSecs)
	}
	if !outputAudio {
		params["outputAudio"] = "false"
	}
	if !fadeInFadeOut {
		params["fadeInFadeOut"] = "false"
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, video)

	// get video operation result
	//
	// https://westus.dev.cognitive.microsoft.com/docs/services/565d6516778daf15800928d5/operations/565d6517778daf0978c45e36
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
						if status.ResourceLocation != "" {
							return status.ResourceLocation, nil
						} else {
							err = fmt.Errorf("resourceLocation is empty")
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

	return "", err
}
