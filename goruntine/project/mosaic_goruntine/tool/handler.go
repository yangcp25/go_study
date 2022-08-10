package tool

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strconv"
	"sync"
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

	db := CloneTilesDB()

	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)

	c := combine(bounds, c1, c2, c3, c4)

	bf1 := new(bytes.Buffer)
	jpeg.Encode(bf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(bf1.Bytes())

	t1 := time.Now()

	imagesRes := map[string]string{
		"original": originalStr,
		"newImage": <-c,
		"time":     fmt.Sprintf("%v ", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("views/results.html")

	t.Execute(writer, imagesRes)
}

func combine(r image.Rectangle, c1, c2, c3, c4 <-chan image.Image) <-chan string {
	c := make(chan string)
	go func() {
		wg := sync.WaitGroup{}
		newImage := image.NewNRGBA(r)
		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			wg.Done()
		}

		wg.Add(4)
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool

		for {
			select {
			case s1, ok1 = <-c1:
				go copy(newImage, s1.Bounds(), s1, image.Point{r.Min.X, r.Min.Y})
			case s2, ok2 = <-c2:
				go copy(newImage, s2.Bounds(), s2, image.Point{r.Max.X / 2, r.Min.Y})
			case s3, ok3 = <-c3:
				go copy(newImage, s3.Bounds(), s3, image.Point{r.Min.X, r.Max.Y / 2})
			case s4, ok4 = <-c4:
				go copy(newImage, s4.Bounds(), s4, image.Point{r.Max.X / 2, r.Max.Y / 2})
			}
			if ok1 && ok2 && ok3 && ok4 {
				break
			}
		}
		wg.Wait()
		buf2 := new(bytes.Buffer)
		// 将合并后的最终马赛克图片进行 base64 编码并写入 c 通道返回
		jpeg.Encode(buf2, newImage, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()
	return c
}

func cut(original image.Image, db *DB, tileSize int, x1, y1, x2, y2 int) <-chan image.Image {
	c := make(chan image.Image)
	// 从原始图标将图片按tileSize分成几个区域
	sp := image.Point{0, 0}
	go func() {
		newImage := image.NewRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}
				nearest := db.Nearest(color)

				file, err := os.Open(nearest)

				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						t := Resize(img, tileSize)
						tile := t.SubImage(t.Bounds())
						tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
						draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
					} else {
						fmt.Println("加载图片出错2", err, file)
					}
				} else {
					fmt.Println("载入图片出错0", err, file, nearest)
				}
				file.Close()
			}
		}
		c <- newImage.SubImage(newImage.Rect)
	}()
	return c
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("views/upload.html")
	t.Execute(writer, nil)
}
