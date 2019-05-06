package cloudstorageexample

import (
	"context"
	"html/template"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/log"
)

func serveError(ctx context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
	log.Errorf(ctx, "%v", err)
}

var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

const rootTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form></body></html>
`

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "rrrrrrrr", r)
	uploadURL, err := blobstore.UploadURL(ctx, "/upload", nil)
	log.Infof(ctx, "uploadURL uploadURL", uploadURL)
	if err != nil {
		serveError(ctx, w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = rootTemplate.Execute(w, uploadURL)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
}

func handleServe(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}

func handleUpload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := appengine.NewContext(r)

	w.Header().Set("Content-Type", "form-data")

	blobs, _, err := blobstore.ParseUpload(r)
	log.Infof(ctx, "blobs blobs", blobs)
	if err != nil {
		serveError(ctx, w, err)
		return
	}

	file := blobs["file"]
	log.Infof(ctx, "file file", file)
	if len(file) == 0 {
		log.Errorf(ctx, "no file uploaded")
		//http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
}

/*
func init() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/serve/", handleServe)
	http.HandleFunc("/upload", handleUpload)
}*/
