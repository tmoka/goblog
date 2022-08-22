package main

import (
	"net/http"

	"github.com/tmoka/goblog/src/controller"
)

// setRouter ルーティングをセット
func setRouter() {
	// ユーザのルーティング
	http.HandleFunc("/", controller.IndexHandler)
}
