package forcego

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router := httprouter.New()

	router.POST("/api/organization", GetAccountHandler)

	http.Handle("/", router)
}
