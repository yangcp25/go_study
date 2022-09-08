package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

func main() {
	//defaultFunc()
	// 绑定结构体
	//bindStruct()
	//bindFormCheckBox()
	//bindPerson()
	//bindUrl()
	//bindCustomize()
	//createLog()
	//bindFormValidate()
	bindGo()
}

func bindGo() {
	router := gin.Default()
	// 使用协程需要拷贝 因为是全局变量要避免污染
	router.GET("/test1", func(context *gin.Context) {
		cgo := context.Copy()

		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(cgo.Request.URL.Path)
		}()
	})
	router.GET("/test2", func(context *gin.Context) {
		cgo := context.Copy()
		fmt.Println(cgo.Request.URL.Path)
	})
	router.Run()
}

type Form struct {
	CheckIn  time.Time `form:"check_in" binding:"required,validateCheckIn" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var validateCheckIn validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)

	if ok {
		now := time.Now()
		if now.After(date) {
			return false
		}
	}
	return true
}

func getDataForm(context *gin.Context) {
	var f Form
	if err := context.ShouldBindWith(&f, binding.Query); err == nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "输入有效！",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "输入无效！",
			"err":     err.Error(),
		})
	}
}

func bindFormValidate() {
	router := gin.Default()

	// 注册自定义验证规则
	validatorObj, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		validatorObj.RegisterValidation("validateCheckIn", validateCheckIn)
	}
	router.GET("/validate", getDataForm)
	router.Run()
}

func createLog() {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义日志输出格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// 使用 recovery 中间件
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8080")
}

func bindCustomize() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

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
