package emotion

import (
	"github.com/meinside/ms-cognitive-services-go"
)

type Client struct {
	ApiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}

// Emotion Recognition in Image
//
// image : string(image url) or []byte(image bytes array)
// rects : rectangles of faces (can be nil if none)
func (c *Client) RecognizeImage(
	image interface{},
	rects []cognitive.Rectangle,
) (emotions []cognitive.Emotion, err error) {
	return cognitive.EmotionRecognizeImage(
		c.ApiKey,
		image,
		rects,
	)
}

// Emotion Recognition in Video
//
// video            : string(video url) or []byte(video bytes array)
// outputStyle      : "aggregate" (default) or "perFrame"
// progressNotifier : can be nil
func (c *Client) RecognizeVideo(
	video interface{},
	outputStyle string,
	progressNotifier func(status string, progress float32),
) (processResult cognitive.EmotionProcessingResult, err error) {
	return cognitive.EmotionRecognizeVideo(
		c.ApiKey,
		video,
		outputStyle,
		progressNotifier,
	)
}
