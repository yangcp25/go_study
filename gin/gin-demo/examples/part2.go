package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
)

func part2() {
	// 绑定到不同实体
	//bindDifferentStruct()
	//上传文件
	//uploadFiles()
	//BasicAuth 中间件
	//secretsFunc()
	//BasicAuth 中间件
	//httpFunc()
	//middlewareFunc()
	//notMiddlewareFunc()
	dataFormate()
}

func dataFormate() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output: {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8088
	r.Run(":8088")
}

func notMiddlewareFunc() {
	//r := gin.New()
	// 这种情况下默认会使用 Logger 和 Recovery 中间件
	r := gin.Default()
	r.Run()
}

func middlewareFunc() {
	// 创建一个默认不使用任何中间件的路由器
	r := gin.New()

	// 设置全局中间件
	// Logger 中间件会记录日志到 gin.DefaultWriter，即使你设置了 GIN_MODE=release 这个环境变量
	// 默认情况下 gin.DefaultWriter = os.Stdout（控制台标准输出）
	r.Use(gin.Logger())

	// Recovery 中间件会从任意 panics 中恢复并且返回 500 响应（服务端错误）
	r.Use(gin.Recovery())

	// 设置路由中间件
	// 路由中间件可以被添加到指定路由上，并且不限数量
	r.GET("/benchmark", MyBenchLogger, benchEndpoint)

	// 为指定路由分组设置中间件
	// 认证组
	// authorized := r.Group("/", AuthRequired())
	// 上面这行代码等同于下面下的代码:
	authorized := r.Group("/")
	// 下面为 `/` 前缀的路由分组设置中间件 AuthRequired，表示该分组下的路由用户认证后才能访问
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func MyBenchLogger(context *gin.Context) {

}

func analyticsEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("sss")
	}
}

func benchEndpoint(context *gin.Context) {

}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "HTTP GET Method")
}

func posting(c *gin.Context) {
	c.String(http.StatusOK, "HTTP POST Method")
}

func putting(c *gin.Context) {
	c.String(http.StatusOK, "HTTP PUT Method")
}

func deleting(c *gin.Context) {
	c.String(http.StatusOK, "HTTP DELETE Method")
}

func patching(c *gin.Context) {
	c.String(http.StatusOK, "HTTP PATCH Method")
}

func head(c *gin.Context) {
	c.String(http.StatusOK, "HTTP HEAD Method")
}

func options(c *gin.Context) {
	c.String(http.StatusOK, "HTTP OPTIONS Method")
}
func httpFunc() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/get", getting)
	router.POST("/post", posting)
	router.PUT("/put", putting)
	router.DELETE("/delete", deleting)
	router.PATCH("/patch", patching)
	router.HEAD("/head", head)
	router.OPTIONS("/options", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}

func secretsFunc() {
	r := gin.Default()

	// 在 /admin 分组中使用 gin.BasicAuth() 中间件
	// 通过 gin.Accounts 来初始化一些测试用户信息（用户名/密码键值对）
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets
	// 匹配 "localhost:8088/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，通过 BasicAuth 中间件设置
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 监听 0.0.0.0:8080，等待客户端请求
	r.Run(":8088")
}

// 模拟一些私有数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func uploadFiles() {
	router := gin.Default()
	// 为 multipart forms 设置更小的内存限制 (默认是 32 MB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// 为 multipart forms 设置更小的内存限制 (默认是 32 MB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MB
	router.POST("/uploads", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8088")
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func bindDifferentStruct() {
	r := gin.Default()
	r.POST("/bindBodyToStruct", SomeHandler)
	r.POST("/bindBodyToStruct2", SomeHandler2)
	r.Run(":8088")
}

func SomeHandler2(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用 c.Request.Body 并且不能被复用
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这里总是会报错，因为 c.Request.Body 现在是 EOF
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		// 其他结构体
		c.String(http.StatusOK, `the body should be other form`)
	}
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用 c.Request.Body 并且不能被复用
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 这里总是会报错，因为 c.Request.Body 现在是 EOF
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		// 其他结构体
		c.String(http.StatusOK, `the body should be other form`)
	}
}
