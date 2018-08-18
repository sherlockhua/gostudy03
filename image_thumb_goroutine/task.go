package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/nfnt/resize"
)


const (
	TaskThumb = iota
)

const (
	 DEFAULT_MAX_WIDTH float64 = 64
	 DEFAULT_MAX_HEIGHT float64 = 64
)

type Task struct  {
	imageFile string
	saveFile string
	taskType int
}

// 计算图片缩放后的尺寸
func (t *Task)calcRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func (t *Task)generateThumb() error {

	file, _ := os.Open(t.imageFile)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := t.calcRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, err := os.Create(t.saveFile)
	if err != nil {
		fmt.Printf("os create file, err:%v\n", err)
		return err
	}
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = png.Encode(imgfile, m)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) Process() (err error) {
	switch (t.taskType) {
	case TaskThumb:
		return t.generateThumb()
	default:
		err = fmt.Errorf("task type:%d is not support\n", t.taskType)
	}
	return
}
