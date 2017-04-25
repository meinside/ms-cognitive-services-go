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

func TestVideoFaceDetectTrack(t *testing.T) {
	// test with a video file
	if vidBytes, err := ioutil.ReadFile(testKeys["face-video"]); err == nil {
		if detection, err := VideoFaceDetectTrack(
			testKeys["video-subscription-key"],
			vidBytes,
			func(status string, progress float32) {
				fmt.Printf("[%s] %.2f%%...\n", status, progress)
			},
		); err == nil {
			fmt.Printf("VideoFaceDetectTrack() => %+v\n", detection)
		} else {
			t.Errorf("VideoFaceDetectTrack() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestVideoMotionDetect(t *testing.T) {
	// test with a video file
	if vidBytes, err := ioutil.ReadFile(testKeys["face-video"]); err == nil {
		if detection, err := VideoMotionDetect(
			testKeys["video-subscription-key"],
			vidBytes,
			"",
			1,
			nil,
			false,
			0.0,
			func(status string, progress float32) {
				fmt.Printf("[%s] %.2f%%...\n", status, progress)
			},
		); err == nil {
			fmt.Printf("VideoMotionDetect() => %+v\n", detection)
		} else {
			t.Errorf("VideoMotionDetect() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestVideoStabilize(t *testing.T) {
	// test with a video file
	if vidBytes, err := ioutil.ReadFile(testKeys["face-video"]); err == nil {
		if url, err := VideoStabilize(
			testKeys["video-subscription-key"],
			vidBytes,
			func(status string, progress float32) {
				fmt.Printf("[%s] %.2f%%...\n", status, progress)
			},
		); err == nil {
			fmt.Printf("VideoStabilize() => %s\n", url)

			// test download
			if err := Download(url, testKeys["video-subscription-key"], testKeys["generated-video"]); err != nil {
				t.Errorf("VideoStabilize() result download failed: %s\n", err)
			}
		} else {
			t.Errorf("VideoStabilize() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestVideoThumbnail(t *testing.T) {
	// test with a video file
	if vidBytes, err := ioutil.ReadFile(testKeys["video"]); err == nil {
		if url, err := VideoThumbnail(
			testKeys["video-subscription-key"],
			vidBytes,
			0,
			true,
			true,
			func(status string, progress float32) {
				fmt.Printf("[%s] %.2f%%...\n", status, progress)
			},
		); err == nil {
			fmt.Printf("VideoThumbnail() => %s\n", url)

			// test download
			if err := Download(url, testKeys["video-subscription-key"], testKeys["generated-video"]); err != nil {
				t.Errorf("VideoThumbnail() result download failed: %s\n", err)
			}
		} else {
			t.Errorf("VideoThumbnail() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}
