package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"slices"
	"strconv"

	"atomicgo.dev/keyboard/keys"
	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfacedetect [camera ID] [classifier XML file]")
		return
	}
	exitKeys := []keys.KeyCode{keys.Enter, keys.Escape}

	// parse args
	deviceID, _ := strconv.Atoi(os.Args[1])
	xmlFile := os.Args[2]

	// open webcam
	webcam, err := gocv.VideoCaptureDeviceWithAPIParams(int(deviceID), gocv.VideoCaptureV4L2, []gocv.VideoCaptureProperties{gocv.VideoCaptureFPS, 15})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}
	// numFacesDetected:= 0
	facesDetected := []image.Rectangle{}
	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		if len(rects) > 0 {
			facesDetected = slices.Clone(rects)
		}
		// fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range facesDetected {
			rect := image.Rect(r.Min.X-10, r.Min.Y-10, r.Max.X+10, r.Max.Y+10)
			gocv.Rectangle(&img, rect, blue, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(rect.Min.X+2, rect.Min.Y-(size.Y/2))
			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)

		if slices.Contains(exitKeys, keys.KeyCode(window.WaitKey(1))) {
			break
		}
	}
}
