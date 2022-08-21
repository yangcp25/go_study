package main

import (
	"log"
	"net/http"
	. "routes/routes"
)

func main() {
	initWeb()
}

func initWeb() {
	startWebServe("8091")
}

func startWebServe(port string) {
	r := NewRouter()
	http.Handle("/", r)

	log.Println("start web on port:" + port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("an error occurred on start web at port " + port)
		log.Println("error:" + err.Error())
	}
}
