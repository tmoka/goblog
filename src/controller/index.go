// controller/index.go
package controller

import (
	"net/http"

	"github.com/tmoka/goblog/src/lib"
)

func IndexHandler(w http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodGet {
		lib.IndexRender(w)
	}
}
