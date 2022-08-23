package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	initTemplate()
}

func initTemplate() {
	http.HandleFunc("/test", parseHtml)
	http.HandleFunc("/string", parseString)
	http.HandleFunc("/mui", parseMui)
	http.HandleFunc("/condition", process)
	http.HandleFunc("/rangeArray", rangeArray)
	http.HandleFunc("/set", setHtml)
	http.HandleFunc("/tpl", tpl)
	http.HandleFunc("/pipeline", pipeline)
	http.HandleFunc("/date_func", date_func)
	http.HandleFunc("/context", context)
	http.HandleFunc("/xss", xssAttackExample)
	http.HandleFunc("/layout", layout)
	http.HandleFunc("/layout_example", layoutExample)
	http.ListenAndServe(":8089", nil)
}

func layoutExample(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t = template.Must(template.ParseFiles("html/layout.html", "html/hello_blue.html"))
	} else {
		t = template.Must(template.ParseFiles("html/layout.html"))
	}
	t.ExecuteTemplate(w, "layout", "")
}

func layout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/layout.html", "html/hello.html")
	t.ExecuteTemplate(w, "layout", "")
}

func xssAttackExample(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/xss.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func xss(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/xss.html"))
	if r.Method == "GET" {
		t.Execute(w, nil)
	} else {
		t.Execute(w, r.FormValue("comment"))
	}
}

func context(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("html/content.html"))
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func formateDate(t time.Time) string {
	layout := "2022年8月23日 20:51:07"
	return t.Format(layout)
}
func date_func(writer http.ResponseWriter, request *http.Request) {
	funcMap := template.FuncMap{
		"fdate": formateDate,
	}

	t := template.New("function.html").Funcs(funcMap)

	t, _ = t.ParseFiles("html/function.html")
	t.Execute(writer, time.Now())
}

func pipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("html/pipiline.html"))

	t.Execute(writer, 12.23)
}

func tpl(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("html/t1.html", "html/t2.html")
	t.Execute(writer, "hello")
	//t.ExecuteTemplate(writer, "html/t2.html", "haha")
}

func setHtml(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("html/set.html"))

	t.Execute(writer, "golang")
}

func rangeArray(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("html/rangeArray.html"))

	stringS := []string{"one", "two", "three", "", "three", "three"}

	t.Execute(writer, stringS)
}

func process(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("html/condition.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(writer, rand.Intn(10) > 5)
}

func parseMui(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("html/t2.html", "html/t1.html")
	//t.Execute(writer, "hello")
	t.ExecuteTemplate(writer, "html/t2.html", "haha")
}

func parseString(writer http.ResponseWriter, request *http.Request) {
	tmpl := `<!DOCTYPE html> <html>
        <head>
            <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
            <title>Go Web Programming</title>
        </head>
        <body>
            {{ . }}
        </body> 
    </html>`
	t := template.New("temp.html")
	t.Parse(tmpl)
	t.Execute(writer, "string")
}

func parseHtml(writer http.ResponseWriter, request *http.Request) {
	//t, _ := template.ParseFiles("html/temp.html")

	t := template.Must(template.ParseFiles("html/temp.html"))
	t.Execute(writer, "哈喽")
}
