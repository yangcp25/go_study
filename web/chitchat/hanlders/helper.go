package handlers

import (
	. "chichat/config"
	"chichat/models"
	"errors"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"
)

var config *Configuration
var localizer *i18n.Localizer

func init() {
	// 获取全局配置实例
	// 获取本地化实例
	//localizer = i18n.NewLocalizer(config.LocaleBundle, config.App.Language)
	localizer = i18n.NewLocalizer(ViperConfig.LocaleBundle, ViperConfig.App.Language)
	os.OpenFile(ViperConfig.App.Log+"/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

// 通过 Cookie 判断用户是否已登录
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 解析 HTML 模板（应对需要传入多个模板文件的情况，避免重复编写模板代码）
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s/%s.html", config.App.Language, file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 生成响应 HTML
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {

	var files []string
	for _, file := range filenames {
		//files = append(files, fmt.Sprintf("views/%s.html", file))
		//files = append(files, fmt.Sprintf("views/%s/%s.html", config.App.Language, file))
		files = append(files, fmt.Sprintf("views/%s/%s.html", ViperConfig.App.Language, file))
	}
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// 返回版本号
func Version() string {
	return "0.1"
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// 日志相关
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// 异常处理统一重定向到错误页面
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 日期格式化
func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}
