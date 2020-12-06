// view関連の汎用関数やインターフェイスの定義
package main

import (
	"net/http"
)

type View interface {
	Validate(r *http.Request) ApplicationError
	WriteResponse(w http.ResponseWriter, r *http.Request)
}
