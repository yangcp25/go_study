package tool

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Mosaic(writer http.ResponseWriter, request *http.Request) {
	t0 := time.Now()
	// 最大上传内容大小设置
	request.ParseMultipartForm(10485760)
	file, _, _ := request.FormFile("image")
	defer file.Close()
	// 设置的区块尺寸
	tileSize, _ := strconv.Atoi(request.FormValue("tile_size"))
	// 解码获取原始图片
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	// 克隆图片
	newImage := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	db := CloneDb()

	// 从原始图标将图片按tileSize分成几个区域
	sp := image.Point{}
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; y < bounds.Max.X; x = x + tileSize {
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}
			nearest := nearest(color, &db)

			file, err := os.Open(nearest)

			if err == nil {
				img, _, err := image.Decode(file)
				if err != nil {
					t := Resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("加载图片出错", err, file)
				}
			} else {
				fmt.Println("载入图片出错", err, file)
			}
			file.Close()
		}
	}

	bf1 := new(bytes.Buffer)
	jpeg.Encode(bf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(bf1.Bytes())

	bf2 := new(bytes.Buffer)
	jpeg.Encode(bf2, newImage, nil)
	newImageStr := base64.StdEncoding.EncodeToString(bf2.Bytes())

	t1 := time.Now()

	imagesRes := map[string]string{
		"original": originalStr,
		"newImage": newImageStr,
		"time":     fmt.Sprintf("%v ", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("views/results.html")

	t.Execute(writer, imagesRes)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("views/upload.html")
	t.Execute(writer, nil)
}
