package cognitive

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func init() {
	//IsVerbose = true // XXX - if you wanna see verbose messages, uncomment it

	// read keys
	initTestKeys()
}

func TestComputerVisionAnalyzeImage(t *testing.T) {
	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["celebrity-face-image"]); err == nil {
		if result, err := ComputerVisionAnalyzeImage(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
			[]string{"Categories", "Tags", "Description", "Faces", "ImageType", "Color", "Adult"},
			[]string{"Celebrities", "Landmarks"},
			"en",
		); err == nil {
			fmt.Printf("ComputerVisionAnalyzeImage() => %+v\n", result)
		} else {
			t.Errorf("ComputerVisionAnalyzeImage() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestComputerVisionDescribeImage(t *testing.T) {
	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["celebrity-face-image"]); err == nil {
		if result, err := ComputerVisionDescribeImage(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
			2,
		); err == nil {
			fmt.Printf("ComputerVisionDescribeImage() => %+v\n", result)
		} else {
			t.Errorf("ComputerVisionDescribeImage() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestComputerVisionGetThumbnail(t *testing.T) {
	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["celebrity-face-image"]); err == nil {
		if result, err := ComputerVisionGetThumbnail(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
			64,
			64,
			true,
		); err == nil {
			fmt.Printf("ComputerVisionGetThumbnail() => %d bytes\n", len(result))

			if err := ioutil.WriteFile(testKeys["generated-image"], result, 0644); err != nil {
				t.Errorf("ComputerVisionGetThumbnail() result download failed: %s\n", err)
			}
		} else {
			t.Errorf("ComputerVisionGetThumbnail() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestComputerVisionGetModels(t *testing.T) {
	if result, err := ComputerVisionGetModels(
		WestUS,
		testKeys["computervision-subscription-key"],
	); err == nil {
		fmt.Printf("ComputerVisionGetModels() => %+v\n", result)

		if len(result.Models) <= 0 {
			t.Errorf("ComputerVisionGetModels() got no domain specific models")
		}
	} else {
		t.Errorf("ComputerVisionGetThumbnail() failed: %s\n", err)
	}
}

func TestComputerVisionDomainSpecificRecognize(t *testing.T) {
	// test with a celebrity image file
	if imgBytes, err := ioutil.ReadFile(testKeys["celebrity-face-image"]); err == nil {
		if result, err := ComputerVisionDomainSpecificRecognize(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
			"celebrities",
		); err == nil {
			fmt.Printf("ComputerVisionDomainSpecificRecognize() => %+v\n", result)
		} else {
			t.Errorf("ComputerVisionDomainSpecificRecognize() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}

	// test with a landmark image file
	if imgBytes, err := ioutil.ReadFile(testKeys["landmark-image"]); err == nil {
		if result, err := ComputerVisionDomainSpecificRecognize(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
			"landmarks",
		); err == nil {
			fmt.Printf("ComputerVisionDomainSpecificRecognize() => %+v\n", result)
		} else {
			t.Errorf("ComputerVisionDomainSpecificRecognize() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestComputerVisionOcr(t *testing.T) {
	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["text-image"]); err == nil {
		if result, err := ComputerVisionOcr(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
			"unk",
			true,
		); err == nil {
			fmt.Printf("ComputerVisionOcr() => %+v\n", result)
		} else {
			t.Errorf("ComputerVisionOcr() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestComputerVisionRecognizeHandwritten(t *testing.T) {
	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["handwritten-image"]); err == nil {
		if result, err := ComputerVisionRecognizeHandwritten(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
			true,
			func(status string, progress float32) {
				fmt.Printf("[%s] %.2f%%...\n", status, progress)
			},
		); err == nil {
			fmt.Printf("ComputerVisionRecognizeHandwritten() => %+v\n", result)
		} else {
			t.Errorf("ComputerVisionRecognizeHandwritten() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}

func TestComputerVisionTagImage(t *testing.T) {
	// test with an image file
	if imgBytes, err := ioutil.ReadFile(testKeys["celebrity-face-image"]); err == nil {
		if result, err := ComputerVisionTagImage(
			WestUS,
			testKeys["computervision-subscription-key"],
			imgBytes,
		); err == nil {
			fmt.Printf("ComputerVisionTagImage() => %+v\n", result)
		} else {
			t.Errorf("ComputerVisionTagImage() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}
