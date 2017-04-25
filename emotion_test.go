package cognitive

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func init() {
	//IsVerbose = true	// XXX - if you wanna see verbose messages, uncomment it

	// read keys
	initTestKeys()
}

func TestEmotionImage(t *testing.T) {
	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["face-image1"]); err == nil {
		if emotions, err := EmotionRecognizeImage(
			testKeys["emotion-subscription-key"],
			imgBytes,
			nil, // []Rectangle{Rectangle{Left: 482, Top: 210, Width: 306, Height: 306}},
		); err == nil {
			fmt.Printf("EmotionRecognizeImage() => %+v\n", emotions)
		} else {
			t.Errorf("EmotionRecognizeImage() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestEmotionVideo(t *testing.T) {
	// test with a video file
	if vidBytes, err := ioutil.ReadFile(testKeys["face-video"]); err == nil {
		if emotions, err := EmotionRecognizeVideo(
			testKeys["emotion-subscription-key"],
			vidBytes,
			"",
			func(status string, progress float32) {
				fmt.Printf("[%s] %.2f%%...\n", status, progress)
			},
		); err == nil {
			fmt.Printf("EmotionRecognizeVideo() => %+v\n", emotions)
		} else {
			t.Errorf("EmotionRecognizeVideo() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}
