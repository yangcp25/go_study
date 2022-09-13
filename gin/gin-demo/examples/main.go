package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	//bindGo()
	// 路由分组
	//ginRouteGroup()
	// 日志文件写入文件
	//logInFile()
	// 模板引擎
	//parseHtml()
	// 自定义模板
	//parseCustomHtml()
	//createHttp2()
	//平滑重启或关闭服务器
	//shotDownHttp()
	// jsonp
	//createJsonP()
	//form
	// 获取字典数据
	//formMap()
	// 绑定结构体
	//bindStructData()
	//
	//test2()
	//运行多个服务器
	//test3()
	//下载文件
	//downFile()
	//资源文件
	//staticFile()
	//cookie
	//setCookie()
	// 自定义证书
	//autotlFunc()
	part2()
}

func autotlFunc() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}

func setCookie() {
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("ycp_token")
		if err != nil {
			cookie = "empty"
			c.SetCookie("ycp_token", "123", 3600, "/", "127.0.0.1", false, true)
		}

		fmt.Println(cookie)
	})
	router.Run()
}

func staticFile() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.Run()
}

func downFile() {
	e := gin.Default()
	e.GET("/dataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://laravel.gstatics.cn/storage/uploads/images/cover_page/2020-08/thumbs-850-350/cvBgin-logo.jpg")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	e.GET("/file", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "test")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		c.File("H:\\code\\go\\go_study\\gin\\gin-demo\\examples\\log.File")
	})
	e.Run()
}

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	e.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}
func test3() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func test2() {
	r := gin.Default()

	// HTTP 重定向
	r.GET("/test", func(c *gin.Context) {
		fmt.Println("in test")
		c.Redirect(http.StatusMovedPermanently, "https://www.xueyuanjun.com/")
	})

	// 路由重定向
	r.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test"
		r.HandleContext(c)
	})

	// 查询字符串参数使用已存在的底层请求对象进行解析
	// 示例请求 URL:  /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 底层调用的是 c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")                 // 查询字符串
		page := c.DefaultQuery("page", "0") // 查询字符串（带默认值）
		name := c.PostForm("name")          //  POST 表单数据
		message := c.PostForm("message")    //  同上

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}

type NameLogin struct {
	Name     string `from:"user" json:"user" xml:"user" binding:"required"`
	Password string `from:"password" json:"password" xml:"password" binding:"required"`
}

func bindStructData() {
	ginWeb := gin.Default()

	// This handler will match /user/xueyuanjun but will not match /user/ or /user
	ginWeb.Any("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/xueyuanjun/ and also /user/xueyuanjun/send
	// If no other routers match /user/xueyuanjun, it will redirect to /user/xueyuanjun/
	ginWeb.Any("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	ginWeb.Any("/testing", startPageF)

	ginWeb.POST("login_s", func(c *gin.Context) {
		var nameLogin NameLogin
		err := c.ShouldBindJSON(&nameLogin)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if nameLogin.Name != "ycp" || nameLogin.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "登录失败",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})
	ginWeb.POST("loginByXml", func(c *gin.Context) {
		var nameLogin NameLogin
		err := c.ShouldBindXML(&nameLogin)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if nameLogin.Name != "ycp" || nameLogin.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "登录失败",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	})

	ginWeb.POST("loginByForm", func(c *gin.Context) {
		var nameLogin NameLogin
		err := c.ShouldBind(&nameLogin)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if nameLogin.Name != "ycp" || nameLogin.Password != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "登录失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "登录成功",
			})
		}
	})
	ginWeb.POST("/login", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		var form NameLogin
		// in this case proper binding will be automatically selected
		if c.ShouldBind(&form) == nil {
			if form.Name == "xueyuanjun" && form.Password == "123456" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	ginWeb.Run()
}

func startPageF(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

func formMap() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		queryData := c.QueryMap("ids")
		formData := c.PostFormMap("test2")
		data := map[string]interface{}{
			"name": "ycp",
		}

		fmt.Println(queryData)
		fmt.Println(formData)
		c.JSONP(http.StatusOK, data)
	})

	r.Run()
}

func createJsonP() {
	r := gin.Default()

	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"name": "ycp",
		}

		//如果传递的查询字符串 callback=hello
		//则输出：hello({\"name\":\"学院君\"})
		c.JSONP(http.StatusOK, data)
	})

	r.Run()
}

func shotDownHttp() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	c := <-quit
	fmt.Println(c)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		return
	}
	select {
	case <-ctx.Done():
		log.Println("5秒过后")
	}
	log.Println("服务器已退出")
}

var html = template.Must(template.New("http2-push").Parse(`
    <html>
    <head>
      <title>Https Test</title>
      <script src="/assets/app.js"></script>
    </head>
    <body>
      <h1 style="color:red;">Welcome, Ginner!</h1>
    </body>
    </html>
    `))

func createHttp2() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "http2-push", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in http://127.0.0.1:8085
	r.Run(":8085")
}

func parseCustomHtml() {
	ginWeb := gin.Default()

	ginWeb.Delims("[[", "]]")

	ginWeb.SetFuncMap(template.FuncMap{
		"formateDate": formateDate,
	})
	ginWeb.LoadHTMLGlob("templates/**/*")
	ginWeb.GET("/ping", func(context *gin.Context) {
		context.HTML(http.StatusOK, "file1.html", map[string]interface{}{
			"date": time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC),
		})
	})
	ginWeb.Run(":8080")
}

func formateDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d%02d", year, month, day)
}

func parseHtml() {
	ginWeb := gin.Default()

	ginWeb.LoadHTMLGlob("templates/**/*")

	ginWeb.GET("/html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "html/index.tmpl", gin.H{
			"title": "测试",
		})
	})
	ginWeb.Run(":8080")
}

func logInFile() {
	gin.DisableConsoleColor()

	f, _ := os.Create("log.File")

	gin.DefaultWriter = io.MultiWriter(f)
	//同时写入文件和输出日志
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	ginWeb := gin.Default()

	ginWeb.GET("/ping", func(context *gin.Context) {
		fmt.Println("pong")
	})
	ginWeb.Run(":8080")
}

func ginRouteGroup() {
	ginWeb := gin.New()

	v1 := ginWeb.Group("v1")
	{
		v1.GET("login", doLogin)
		v1.GET("submit", doSubmit)
		v1.GET("loginOut", doLoginOut)
	}

	v2 := ginWeb.Group("v2")
	{
		v2.POST("login", doLoginV2)
		v2.POST("submit", doSubmitV2)
		v2.POST("loginOut", doLoginOutV2)
	}
	ginWeb.Run(":8080")
}

func doLoginOutV2(context *gin.Context) {
	fmt.Println("v2:loginOut")
}

func doSubmitV2(context *gin.Context) {
	fmt.Println("v2:submit")
}

func doLoginV2(context *gin.Context) {
	fmt.Println("v2:login")
}

func doLoginOut(context *gin.Context) {
	fmt.Println("loginOut")
}

func doSubmit(context *gin.Context) {
	fmt.Println("submit")
}

func doLogin(context *gin.Context) {
	fmt.Println("login")
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
