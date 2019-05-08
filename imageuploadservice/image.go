package cloudstorageexample

import (
	"fmt"
	"io"
	"net/http"
	"path"

	"cloud.google.com/go/storage"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const bucketID = "ap1-simplyst-health.appspot.com"

func init() {

	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.URL.Path != "/" {

		return
	}

	html := `
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="image">
		<input type="submit">
	    </form>
	`
	if r.Method == "POST" {

		f, fh, err := r.FormFile("image")
		if err == http.ErrMissingFile {
			fmt.Println("Error1=>")
			return
		}
		if err != nil {
			return
		}

		// random filename, retaining existing extension.
		name := uuid.Must(uuid.NewV4()).String() + path.Ext(fh.Filename)

		//ctx := context.Background()
		client, err := storage.NewClient(ctx)
		w := client.Bucket(bucketID).Object(name).NewWriter(ctx)

		// Warning: storage.AllUsers gives public read access to anyone.
		w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
		w.ContentType = fh.Header.Get("Content-Type")

		// Entries are immutable, be aggressive about caching (1 day).
		w.CacheControl = "public, max-age=86400"

		if _, err := io.Copy(w, f); err != nil {
			log.Infof(ctx, "Error2=> ", err)
			return
		}
		if err := w.Close(); err != nil {
			//fmt.Println("")
			log.Infof(ctx, "Error3=> ", err)
			return
		}

		publicURL := "https://storage.googleapis.com/" + bucketID + "/" + name
		log.Infof(ctx, "Public URL ", publicURL)

	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, html)
}
