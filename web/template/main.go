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
	http.ListenAndServe(":8089", nil)
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
