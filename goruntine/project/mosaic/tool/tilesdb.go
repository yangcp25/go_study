package tool

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
)

var TILESDB map[string][3]float64

// CloneDb 克隆一个图片DB
func CloneDb() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, item := range TILESDB {
		db[k] = item
	}
	return db
}

// TilesDb 构建相识图片库
func TilesDb() map[string][3]float64 {
	fmt.Println("开始构建相识图片库...")
	db := make(map[string][3]float64)

	files, err := ioutil.ReadDir("tiles")
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		name := "tiles/" + file.Name()
		fmt.Println(name)
		f, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(f)
			if err == nil {
				db[name] = AverageColor(img)
			} else {
				fmt.Println("构建嵌入图片数据库出错", err, name)
			}
		} else {
			fmt.Println("载入图片出错1", err, name)
		}
		f.Close()
	}
	fmt.Print(db)
	fmt.Println("完成嵌入图片数据库构建")
	return db
}

// 查找最相似的图片
func nearest(img [3]float64, db *map[string][3]float64) string {
	var file string
	var smallDis = 1000000.0
	for name, rgb := range *db {
		dis := distance(rgb, img)
		fmt.Printf("%v", dis)
		if dis < smallDis {
			file, smallDis = name, dis
		}
	}
	/*if file == "" {
		for name, _ := range *db {
			file = name
			break
		}
	}*/
	delete(*db, file)
	return file
}

func distance(img [3]float64, rgb [3]float64) float64 {
	// 求 3个像素的差集的平方和 再开方
	return math.Sqrt(math.Pow(img[0]-rgb[0], 2) + math.Pow(img[1]-rgb[1], 2) + math.Pow(img[2]-rgb[2], 2))
}
