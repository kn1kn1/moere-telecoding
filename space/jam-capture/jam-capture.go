// ## prepare and run (OSX)
// ```sh
// $ brew install homebrew/science/opencv
// $ go get github.com/lazywei/go-opencv
// $ go run jam-capture.go
// ```
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	//"path"
	//"runtime"
	//"strings"

	"github.com/lazywei/go-opencv/opencv"
	//"../opencv" // can be used in forks, comment in real application
)

func main() {
	win := opencv.NewWindow("Go-OpenCV Webcam")
	defer win.Destroy()

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("can not open camera")
	}
	defer cap.Release()

	font := opencv.InitFont(opencv.CV_FONT_HERSHEY_DUPLEX, 1, 1, 0, 1, 8)

	fmt.Println("Press ESC to quit")
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
				ProcessImage(img, win, font)
			} else {
				fmt.Println("Image ins nil")
			}
		}
		key := opencv.WaitKey(10)

		if key == 27 {
			os.Exit(0)
		}
	}
}

func ProcessImage(img *opencv.IplImage, win *opencv.Window, font *opencv.Font) error {
	color := opencv.NewScalar(255/2, 255/2, 255/2, 0)
	// w := img.Width()
	// h := img.Height()
	// pt := opencv.Point{w / 2, h / 2}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	absPath := usr.HomeDir + "/" + ".sonic-pi/store/default/haskap-buffer-192.168.100.139-one.spi"
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	y := 25
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		font.PutText(img, scanner.Text(), opencv.Point{2, y}, color)
		y += 25
	}

	// code := "Hello\ngo-opencv"
	// codes := strings.Split(code, "\n")
	// for _, element := range codes {
	// 		font.PutText(img, element, opencv.Point{2, y}, color)
	// 		y += 22
	// }

	win.ShowImage(img)
	return nil
}
