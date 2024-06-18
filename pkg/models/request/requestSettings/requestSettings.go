package requestSettings

import "net/http"

type RequestSettings struct {
	Url     string
	Method  Method
	Body    string
	Headers http.Header
	Cookies []http.Cookie
}
