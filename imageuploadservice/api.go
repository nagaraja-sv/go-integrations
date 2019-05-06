package cloudstorageexample

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
	//http.HandleFunc("/api/image/", handleUpload)
	router := httprouter.New()

	router.POST("/api/image", handleUpload)
	http.Handle("/", router)

}

/*
//blobs
	blobs, _, err := blobstore.ParseUpload(r)
	log.Errorf(ctx, "blobs blobs", handleUpload)

	if err != nil {
		serveError(ctx, w, err)
		return
	}
	file := blobs["myFile"]
	log.Errorf(ctx, "rrrrrrrrrrrrrrrrr", file)
	if len(file) == 0 {
		log.Errorf(ctx, "no file uploaded")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)

	func init() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/serve/", handleServe)
	//http.HandleFunc("/upload", handleUpload)
}

*/
