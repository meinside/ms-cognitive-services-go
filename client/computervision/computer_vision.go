package computervision

// Wrapper Client for ComputerVision functions

import (
	"github.com/meinside/ms-cognitive-services-go"
)

type Client struct {
	Location cognitive.ApiLocation
	ApiKey   string
}

func NewClient(apiKey string) *Client {
	return &Client{
		Location: cognitive.WestUS,
		ApiKey:   apiKey,
	}
}

func NewClientWithLocation(location cognitive.ApiLocation, apiKey string) *Client {
	return &Client{
		Location: location,
		ApiKey:   apiKey,
	}
}

// Analyze Image
//
// image          : string(image url) or []byte(image bytes array)
// visualFeatures : "Categories", "Tags", "Description", "Faces", "ImageType", "Color", "Adult"
// details        : "Celebrities", "Landmarks"
// language       : "en" or "zh" (default: "en")
func (c *Client) AnalyzeImage(
	image interface{},
	visualFeatures []string,
	details []string,
	language string,
) (processResult cognitive.ComputerVisionImageAnalyzeResult, err error) {
	return cognitive.ComputerVisionAnalyzeImage(
		c.Location,
		c.ApiKey,
		image,
		visualFeatures,
		details,
		language,
	)

}

// Describe Image
//
// image         : string(image url) or []byte(image bytes array)
// maxCandidates : default 1
func (c *Client) DescribeImage(
	image interface{},
	maxCandidates int,
) (processResult cognitive.ComputerVisionImageDescribeResult, err error) {
	return cognitive.ComputerVisionDescribeImage(
		c.Location,
		c.ApiKey,
		image,
		maxCandidates,
	)
}

// Get Thumbnail
//
// image         : string(image url) or []byte(image bytes array)
// width         : 1 ~ 1024 (recommended: minimum 50)
// height        : 1 ~ 1024 (recommended: minimum 50)
// smartCropping :
func (c *Client) GetThumbnail(
	image interface{},
	width int,
	height int,
	smartCropping bool,
) (processResult []byte, err error) {
	return cognitive.ComputerVisionGetThumbnail(
		c.Location,
		c.ApiKey,
		image,
		width,
		height,
		smartCropping,
	)
}

// List Domain Specific Models
func (c *Client) GetModels() (processResult cognitive.ComputerVisionDomainSpecificModelsResult, err error) {
	return cognitive.ComputerVisionGetModels(
		c.Location,
		c.ApiKey,
	)
}

// OCR
//
// image             : string(image url) or []byte(image bytes array)
// language          : "unk", "zh-Hans", "zh-Hant", "cs", "da", "nl", "en", "fi", "fr", "de", "el", "hu", "it", "ja", "ko", "nb", "pl", "pt", "ru", "es", "sv", "tr" (default: "unk")
// detectOrientation :
func (c *Client) Ocr(
	image interface{},
	language string,
	detectOrientation bool,
) (processResult cognitive.ComputerVisionOcrResult, err error) {
	return cognitive.ComputerVisionOcr(
		c.Location,
		c.ApiKey,
		image,
		language,
		detectOrientation,
	)
}

// Recognize Domain Specific Content
//
// image : string(image url) or []byte(image bytes array)
// model :
func (c *Client) DomainSpecificRecognize(
	image interface{},
	model string,
) (processResult cognitive.ComputerVisionDomainSpecificResult, err error) {
	return cognitive.ComputerVisionDomainSpecificRecognize(
		c.Location,
		c.ApiKey,
		image,
		model,
	)
}

// Recognize Handwritten Text
//
// image            : string(image url) or []byte(image bytes array)
// handwriting      : (default: false)
// progressNotifier : can be nil
func (c *Client) RecognizeHandwritten(
	image interface{},
	handwriting bool,
	progressNotifier func(status string, progress float32),
) (processResult cognitive.ComputerVisionHandwrittenProcessingResult, err error) {
	return cognitive.ComputerVisionRecognizeHandwritten(
		c.Location,
		c.ApiKey,
		image,
		handwriting,
		progressNotifier,
	)
}

// Tag Image
//
// image : string(image url) or []byte(image bytes array)
func (c *Client) TagImage(
	image interface{},
) (processResult cognitive.ComputerVisionTagImageResult, err error) {
	return cognitive.ComputerVisionTagImage(
		c.Location,
		c.ApiKey,
		image,
	)
}
