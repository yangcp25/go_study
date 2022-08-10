package tool

import (
	"image"
	"image/color"
)

// AverageColor 计算平均颜色
func AverageColor(image image.Image) [3]float64 {
	bounds := image.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			// 把所有的x,y每个坐标的像素的rgb 分表求和
			tempR, tempG, tempB, _ := image.At(i, j).RGBA()
			r, g, b = r+float64(tempR), g+float64(tempG), b+float64(tempB)
		}
	}
	// 总的像素个数
	total := float64(bounds.Max.X * bounds.Max.Y)

	return [3]float64{
		r / total,
		g / total,
		b / total,
	}
}

// Resize 将图片重置尺寸, 图片按比例缩放
func Resize(img image.Image, newWidth int) *image.NRGBA {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	// 缩放比例
	rate := width / newWidth

	// 新建一个按比例缩放的图像
	newImage := image.NewNRGBA(image.Rect(bounds.Min.X/rate, bounds.Min.Y/rate, bounds.Max.X/rate, bounds.Max.Y/rate))

	// 按比例重新存放像素值
	for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+rate, i+1 {
		for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+rate, j+1 {
			r, g, b, a := img.At(x, y).RGBA()
			newImage.SetNRGBA(i, j, color.NRGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			})
		}
	}
	return newImage
}
