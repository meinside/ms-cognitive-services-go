# ms-cognitive-services-go

Go library for MS Cognitive Services

## How to get

```bash
$ git clone https://github.com/meinside/ms-cognitive-services-go.git
```

or

```bash
$ go get -d github.com/meinside/ms-cognitive-services-go/...
```

## How to test

You need API keys for MS Congitive Services.

Copy a sample config file and edit it,

```bash
$ cd ms-cognitive-services-go
$ cp files/testconf.json.sample files/testconf.json
$ vi files/testconf.json
```

like this:

```json
{
	"emotion-subscription-key": "1234567890abcdefghijklmnopqrstuvwxyz",
	"video-subscription-key": "abcdefghijklmnopqrstuvwxyz1234567890",
	"computervision-subscription-key": "0987654321abcdefghijklmnopqrstuvwxyz",
	"face-subscription-key": "abcdefghijklmnopqrstuvwxyz0987654321",
	"celebrity-face-image": "/Users/meinside/Downloads/kimjongun.jpg",
	"landmark-image": "/Users/meinside/Downloads/eiffeltower.jpg",
	"handwritten-image": "/Users/meinside/Downloads/handwritten.jpg",
	"text-image": "/Users/meinside/Downloads/text.png",
	"face-image1": "/Users/meinside/Downloads/face1.jpg",
	"face-image2": "/Users/meinside/Downloads/face2.jpg",
	"face-video": "/Users/meinside/Downloads/face.mp4",
	"video": "/Users/meinside/Downloads/video.mp4",
	"generated-image": "/Users/meinside/Downloads/new.jpg",
	"generated-video": "/Users/meinside/Downloads/new.mp4"
}
```

then run test:

```bash
$ go test
```

## How to use

### Sample code

```go
// sample.go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/meinside/ms-cognitive-services-go"
)

const (
	EmotionApiKey = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func main() {
	// get emotions from a facial image
	if imgBytes, err := ioutil.ReadFile("/tmp/face_image.jpg"); err == nil {
		if emotions, err := cognitive.EmotionRecognizeImage(
			EmotionApiKey,
			imgBytes,
			nil,
		); err == nil {
			fmt.Printf("Emotion: %+v\n", emotions)
		} else {
			fmt.Printf("cognitive.EmotionRecognizeImage() failed: %s\n", err)
		}
	} else {
		fmt.Printf("File read error.\n")
	}
}
```

## License

MIT
