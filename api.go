package forcego

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// callwithParams function is helping us to call controller from middleware having access to URL params
func callwithParams(router *httprouter.Router, handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := getURLParams(router, r)
		handler(w, r, params)
	}
}

// getUrlParams function is extracting URL parameters
func getURLParams(router *httprouter.Router, req *http.Request) httprouter.Params {
	_, params, _ := router.Lookup(req.Method, req.URL.Path)
	return params
}
func init() {

	initKeys()
	router := httprouter.New()

	router.POST("/api/organization", GetAccountHandler)

	http.Handle("/", router)
}
