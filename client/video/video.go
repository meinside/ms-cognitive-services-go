package video

// Wrapper Client for Video functions

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

// Face Detection and Tracking
//
// video            : string(video url) or []byte(video bytes array)
// progressNotifier : can be nil
func (c *Client) FaceDetectTrack(
	video interface{},
	progressNotifier func(status string, progress float32),
) (processResult cognitive.VideoProcessingResult1, err error) {
	return cognitive.VideoFaceDetectTrack(
		c.ApiKey,
		video,
		progressNotifier,
	)
}

// Motion Detection
//
// video              : string(video url) or []byte(video bytes array)
// sensitivityLevel   : "low", "medium", or "high" (default: "medium")
// frameSamplingValue : 1 ~ 20 (default: 1)
// detectionZones     : can be nil
// detectLightChange  : default false
// mergeTimeThreshold : 0.0 ~ 10.0 (default: 0.0)
// progressNotifier   : can be nil
func (c *Client) MotionDetect(
	video interface{},
	sensitivityLevel string,
	frameSamplingValue int,
	detectionZones [][]cognitive.Point,
	detectLightChange bool,
	mergeTimeThreshold float64,
	progressNotifier func(status string, progress float32),
) (processResult cognitive.VideoProcessingResult2, err error) {
	return cognitive.VideoMotionDetect(
		c.ApiKey,
		video,
		sensitivityLevel,
		frameSamplingValue,
		detectionZones,
		detectLightChange,
		mergeTimeThreshold,
		progressNotifier,
	)
}

// Stabilization
//
// video            : string(video url) or []byte(video bytes array)
// progressNotifier : can be nil
func (c *Client) Stabilize(
	video interface{},
	progressNotifier func(status string, progress float32),
) (fileUrl string, err error) {
	return cognitive.VideoStabilize(
		c.ApiKey,
		video,
		progressNotifier,
	)
}

// Thumbnail
//
// video                            : string(video url) or []byte(video bytes array)
// maxMotionThumbnailDurationInSecs : default 0
// outputAudio                      : default true
// fadeInFadeOut                    : default true
// progressNotifier                 : can be nil
func (c *Client) Thumbnail(
	video interface{},
	maxMotionThumbnailDurationInSecs int,
	outputAudio bool,
	fadeInFadeOut bool,
	progressNotifier func(status string, progress float32),
) (fileUrl string, err error) {
	return cognitive.VideoThumbnail(
		c.ApiKey,
		video,
		maxMotionThumbnailDurationInSecs,
		outputAudio,
		fadeInFadeOut,
		progressNotifier,
	)
}
