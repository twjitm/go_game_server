package main

import (
	"fmt"
	"go_game_server/photoweb/server"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", server.UploadHandler)
	http.HandleFunc("/view", server.View)
	http.HandleFunc("/", server.FileList)
	e := http.ListenAndServe("127.0.0.1:8909", nil)
	if e != nil {
		fmt.Println("start photo server error", e)
	}
}
