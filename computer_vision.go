package cognitive

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type ComputerVisionImageAnalyzeResult struct {
	Categories []struct {
		Name   string  `json:"name"`
		Score  float32 `json:"score"`
		Detail struct {
			Celebrities []struct {
				Name          string    `json:"name"`
				FaceRectangle Rectangle `json:"faceRectangle"`
				Confidence    float32   `json:"confidence"`
			} `json:"celebrities"`
			Landmarks []struct {
				Name       string  `json:"name"`
				Confidence float32 `json:"confidence"`
			} `json:"landmarks"`
		} `json:"detail"`
	} `json:"categories"`
	Adult struct {
		IsAdultContent bool    `json:"isAdultContent"`
		IsRacyContent  bool    `json:"isRacyContent"`
		AdultScore     float32 `json:"adultScore"`
		RacyScore      float32 `json:"racyScore"`
	} `json:"adult"`
	Tags []struct {
		Name       string  `json:"name"`
		Confidence float32 `json:"confidence"`
	} `json:"tags"`
	Description struct {
		Tags     []string `json:"tags"`
		Captions []struct {
			Text       string  `json:"text"`
			Confidence float32 `json:"confidence"`
		} `json:"captions"`
	}
	RequestId string `json:"requestId"`
	Metadata  struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Format string `json:"format"`
	} `json:"metadata"`
	Faces []struct {
		Age           int       `json:"age"`
		Gender        string    `json:"gender"`
		FaceRectangle Rectangle `json:"faceRectangle"`
	} `json:"faces"`
	Color struct {
		DominantColorForeground string   `json:"dominantColorForeground"`
		DominantColorBackground string   `json:"dominantColorBackground"`
		DominantColors          []string `json:"dominantColors"`
		AccentColor             string   `json:"accentColor"`
		IsBWImg                 bool     `json:"isBWImg"`
	} `json:"color"`
	ImageType struct {
		ClipArtType     int `json:"clipArtType"`     // non-clipart: 0, ambiguous: 1, normal-clipart: 2, good-clipart: 3
		LineDrawingType int `json:"lineDrawingType"` // none-linedrawing: 0, linedrawing: 1
	} `json:"imageType"`
}

type ComputerVisionImageDescribeResult struct {
	Description struct {
		Tags     []string `json:"tags"`
		Captions []struct {
			Text       string  `json:"text"`
			Confidence float32 `json:"confidence"`
		} `json:"captions"`
	} `json:"description"`
	RequestId string `json:"requestId"`
	Metadata  struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Format string `json:"format"`
	} `json:"metadata"`
}

type ComputerVisionDomainSpecificModelsResult struct {
	Models []struct {
		Name       string   `json:"name"`
		Categories []string `json:"categories"`
	} `json:"models"`
}

type ComputerVisionDomainSpecificResult struct {
	RequestId string `json:"requestId"`
	Metadata  struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Format string `json:"format"`
	} `json:"metadata"`
	Result map[string][]struct {
		Name          string    `json:"name"`
		FaceRectangle Rectangle `json:"faceRectangle"`
		Confidence    float32   `json:"confidence"`
	} `json:"result"`
}

type ComputerVisionOcrResult struct {
	Language    string  `json:"language"`
	TextAngle   float32 `json:"textAngle"`
	Orientation string  `json:"orientation"`
	Regions     []struct {
		BoundingBox string `json:"boundingBox"`
		Lines       []struct {
			BoundingBox string `json:"boundingBox"`
			Words       []struct {
				BoundingBox string `json:"boundingBox"`
				Text        string `json:"text"`
			} `json:"words"`
		} `json:"lines"`
	} `json:"regions"`
}

type ComputerVisionHandwrittenProcessingResult struct {
	Lines []struct {
		BoundingBox []int  `json:"boundingBox"`
		Text        string `json:"text"`
		Words       []struct {
			BoundingBox []int  `json:"boundingBox"`
			Text        string `json:"text"`
		} `json:"words"`
	} `json:"lines"`
}

type ComputerVisionTagImageResult struct {
	Tags []struct {
		Name       string  `json:"name"`
		Confidence float32 `json:"confidence"`
	} `json:"tags"`
	RequestId string `json:"requestId"`
	Metadata  struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Format string `json:"format"`
	} `json:"metadata"`
}

// Computer Vision API: Analyze Image
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fa
//
// location       : API location
// key            : subscription key for this API
// image          : string(image url) or []byte(image bytes array)
// visualFeatures : "Categories", "Tags", "Description", "Faces", "ImageType", "Color", "Adult"
// details        : "Celebrities", "Landmarks"
// language       : "en" or "zh" (default: "en")
func ComputerVisionAnalyzeImage(
	location ApiLocation,
	key string,
	image interface{},
	visualFeatures []string,
	details []string,
	language string,
) (processResult ComputerVisionImageAnalyzeResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/analyze"

	// params
	params := map[string]string{}
	if len(visualFeatures) > 0 {
		params["visualFeatures"] = strings.Join(visualFeatures, ",")
	}
	if len(details) > 0 {
		params["details"] = strings.Join(details, ",")
	}
	if language != "" {
		params["language"] = language
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return ComputerVisionImageAnalyzeResult{}, err
}

// Computer Vision API: Describe Image
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fe
//
// location      : API location
// key           : subscription key for this API
// image         : string(image url) or []byte(image bytes array)
// maxCandidates : default 1
func ComputerVisionDescribeImage(
	location ApiLocation,
	key string,
	image interface{},
	maxCandidates int,
) (processResult ComputerVisionImageDescribeResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/describe"

	// params
	params := map[string]string{}
	if maxCandidates > 1 {
		params["maxCandidates"] = strconv.Itoa(maxCandidates)
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return ComputerVisionImageDescribeResult{}, err
}

// Computer Vision API: Get Thumbnail
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fb
//
// location      : API location
// key           : subscription key for this API
// image         : string(image url) or []byte(image bytes array)
// width         : 1 ~ 1024 (recommended: minimum 50)
// height        : 1 ~ 1024 (recommended: minimum 50)
// smartCropping :
func ComputerVisionGetThumbnail(
	location ApiLocation,
	key string,
	image interface{},
	width int,
	height int,
	smartCropping bool,
) (processResult []byte, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/generateThumbnail"

	// params
	params := map[string]string{}
	if width < 1 || width > 1024 || height < 1 || height > 1024 {
		return []byte{}, fmt.Errorf("Parameter width and height should be in range: 1 - 1024")
	}
	params["width"] = strconv.Itoa(width)
	params["height"] = strconv.Itoa(height)
	if smartCropping {
		params["smartCropping"] = "true"
	}

	var result []byte
	if result, err = postArg(apiUrl, key, params, image); err == nil {
		return result, nil
	}
	return []byte{}, err
}

// Computer Vision API: List Domain Specific Models
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fd
//
// location      : API location
// key           : subscription key for this API
// image         : string(image url) or []byte(image bytes array)
// maxCandidates : default 1
func ComputerVisionGetModels(
	location ApiLocation,
	key string,
) (processResult ComputerVisionDomainSpecificModelsResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/models"

	var result []byte
	result, err = httpGet(apiUrl, key, nil)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return ComputerVisionDomainSpecificModelsResult{}, err
}

// Computer Vision API: OCR
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1fc
//
// location          : API location
// key               : subscription key for this API
// image             : string(image url) or []byte(image bytes array)
// language          : "unk", "zh-Hans", "zh-Hant", "cs", "da", "nl", "en", "fi", "fr", "de", "el", "hu", "it", "ja", "ko", "nb", "pl", "pt", "ru", "es", "sv", "tr" (default: "unk")
// detectOrientation :
func ComputerVisionOcr(
	location ApiLocation,
	key string,
	image interface{},
	language string,
	detectOrientation bool,
) (processResult ComputerVisionOcrResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/ocr"

	// params
	params := map[string]string{}
	if language != "" {
		params["language"] = language
	}
	if detectOrientation {
		params["detectOrientation"] = "true"
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return ComputerVisionOcrResult{}, err
}

// Computer Vision API: Recognize Domain Specific Content
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e200
//
// location : API location
// key      : subscription key for this API
// model    :
func ComputerVisionDomainSpecificRecognize(
	location ApiLocation,
	key string,
	image interface{},
	model string,
) (processResult ComputerVisionDomainSpecificResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/models/" + model + "/analyze"

	// params
	params := map[string]string{
		"model": model,
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return ComputerVisionDomainSpecificResult{}, err
}

// Computer Vision API: Recognize Handwritten Text
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/587f2c6a154055056008f200
//
// key              : subscription key for this API
// image            : string(image url) or []byte(image bytes array)
// handwriting      : (default: false)
// progressNotifier : can be nil
func ComputerVisionRecognizeHandwritten(
	location ApiLocation,
	key string,
	image interface{},
	handwriting bool,
	progressNotifier func(status string, progress float32),
) (processResult ComputerVisionHandwrittenProcessingResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/recognizeText"

	params := map[string]string{}
	if handwriting {
		params["handwriting"] = "true"
	}

	var result []byte
	result, err = postArg(apiUrl, key, params, image)

	// get operation result
	//
	// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/587f2cf1154055056008f201
	if err == nil {
		opLocation := string(result)

		if IsVerbose {
			log.Printf("<< Checking operation location...: %s", opLocation)
		}

		var lastProgress float32 = -1.0

		for i := 0; i < NumTries; i++ {
			time.Sleep(WaitSeconds * time.Second)

			var status struct {
				Status            string                                    `json:"status"`
				Progress          float32                                   `json:"progress"`
				Message           string                                    `json:"message"`
				RecognitionResult ComputerVisionHandwrittenProcessingResult `json:"recognitionResult"`
			}
			if result, err = httpGet(opLocation, key, nil); err == nil {
				if err = json.Unmarshal(result, &status); err == nil {
					if status.Status == "Succeeded" {
						return status.RecognitionResult, nil
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

	return ComputerVisionHandwrittenProcessingResult{}, err
}

// Computer Vision API: Tag Image
//
// https://westus.dev.cognitive.microsoft.com/docs/services/56f91f2d778daf23d8ec6739/operations/56f91f2e778daf14a499e1ff
//
// location : API location
// key      : subscription key for this API
func ComputerVisionTagImage(
	location ApiLocation,
	key string,
	image interface{},
) (processResult ComputerVisionTagImageResult, err error) {
	apiUrl := "https://" + string(location) + ".api.cognitive.microsoft.com/vision/v1.0/tag"

	var result []byte
	result, err = postArg(apiUrl, key, nil, image)

	if err == nil {
		if err = json.Unmarshal(result, &processResult); err == nil {
			return processResult, nil
		}
	}
	return ComputerVisionTagImageResult{}, err
}
