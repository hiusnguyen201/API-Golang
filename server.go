package main

import (
	"api-golang/routers"
	"net/http"
)

func main() {
	router := routers.InitRouters()
	http.ListenAndServe(":3000", router)
}