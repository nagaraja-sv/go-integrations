package cloudstorageexample

import (
	"net/http"

	"cloud.google.com/go/storage"
)

func upload(w http.ResponseWriter, r *http.Request) {

	wc := storage.NewWriter(ctx, "ap1-simplyst-health.appspot.com", "filename1")
	wc.ContentType = "image/jpg"
	wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
	if _, err := wc.Write(w.Body); err != nil {

	}

}

/*
const maxUploadSize = 2 * 1024 * 1024 * 1024 // 2 mb
const uploadPath = "./tmp"

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		fileType := r.PostFormValue("type")
		log.Infof(ctx, "fileType", fileType)
		file, _, err := r.FormFile("uploadFile")
		log.Infof(ctx, "file file", file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		log.Infof(ctx, "fileBytes fileBytes", fileBytes)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		log.Infof(ctx, "filetype filetype", filetype)
		switch filetype {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		fileName := randToken(12)
		log.Infof(ctx, "fileName fileName", fileName)
		//fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		newPath := filepath.Join(uploadPath, fileName)
		fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/*
import (
	"context"

	"html/template"
	"io"

	"net/http"
	"time"

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
<html>
<body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="myFile"><br>
<input type="submit" name="submit" value="Submit">
</form>
</body>
</html>
`

type BlobInfo struct {
	BlobKey      appengine.BlobKey
	ContentType  string    `datastore:"content_type"`
	CreationTime time.Time `datastore:"creation"`
	Filename     string    `datastore:"filename"`
	Size         int64     `datastore:"size"`
	MD5          string    `datastore:"md5_hash"`

	// ObjectName is the Google Cloud Storage name for this blob.
	ObjectName string `datastore:"gs_object_name"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	uploadURL, err := blobstore.UploadURL(ctx, "/upload", nil)
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

//var testHead = `multipart/related; charset=utf-8; boundary="example-1"; type="text/xml"; start="<a@b.c>"`

func handleUpload(w http.ResponseWriter, r *http.Request) {

	//r.Header.Set("Content-Type", "multipart/form-data")
	ctx := appengine.NewContext(r)
	key, _ := blobstore.BlobKeyForFile(ctx, "file")
	log.Infof(ctx, "key key", key)
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

	ctx := appengine.NewContext(r)

	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		serveError(ctx, w, err)
		return
	}
	file := blobs["file"]
	if len(file) == 0 {
		log.Errorf(ctx, "no file uploaded")
		//http.Redirect(w, r, "/", http.StatusFound)
		return
	}

}*/
