package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/*func main() {
	//defaultFunc()
	// 绑定结构体
	//bindStruct()
	//bindFormCheckBox()
	//bindPerson()
	bindUrl()
}*/

type Student struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func bindUrl() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var student Student
		// 将路由参数绑定到结构体中
		if err := c.ShouldBindUri(&student); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": student.Name, "uuid": student.ID})
	})
	route.Run()
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}
func bindPerson() {
	route := gin.Default()
	// GET 请求
	route.GET("/testing", startPage)
	// POST 请求
	route.POST("/testing", startPage)
	route.Run()
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

type myForm struct {
	Color []string `form:"colors[]"`
}

func bindFormCheckBox() {
	ginWeb := gin.Default()

	ginWeb.LoadHTMLGlob("templates/**/*")
	ginWeb.GET("/colors", renderCheckBox)
	ginWeb.POST("/colors", getColor)

	ginWeb.Run()
}

func getColor(context *gin.Context) {
	var fakeForm myForm
	// ShouldBind 和 Bind 类似，不过会在出错时退出而不是返回400状态码
	context.ShouldBind(&fakeForm)
	context.JSON(200, gin.H{"color": fakeForm.Color})
}

func renderCheckBox(context *gin.Context) {
	context.HTML(http.StatusOK, "checkbox/color.tmpl", gin.H{
		"title": "xxx",
	})
}

type StructA struct {
	FieldA string
}

type StructB struct {
	nestedStruct StructA
	FieldB       string
}

type StructC struct {
	nestedStruct *StructA
	FieldC       string
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldNest string
	}
	FieldD string
}

func GetDataD(ctx *gin.Context) {
	var d StructD
	ctx.Bind(&d)
	ctx.JSON(200, gin.H{
		"a": d.NestedAnonyStruct,
		"b": d.FieldD,
	})
}

func GetDataC(ctx *gin.Context) {
	var c StructC
	ctx.Bind(&c)
	ctx.JSON(200, gin.H{
		"a": c.nestedStruct,
		"b": c.FieldC,
	})
}

func GetDataB(ctx *gin.Context) {
	var b StructB
	ctx.Bind(&b)
	ctx.JSON(200, gin.H{
		"a": b.nestedStruct,
		"b": b.FieldB,
	})
}

func bindStruct() {
	ginWeb := gin.Default()

	ginWeb.GET("/getB", GetDataB)
	ginWeb.GET("/getC", GetDataC)
	ginWeb.GET("/getD", GetDataD)
	ginWeb.Run()
}

func defaultFunc() {
	ginWeb := gin.Default()

	ginWeb.GET("/test", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "golang",
			"tag":  "<br>",
		}

		context.AsciiJSON(http.StatusOK, data)
	})
	ginWeb.Run(":8080")
}
