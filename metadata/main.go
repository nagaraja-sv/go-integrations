package forcego

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type DynEnt map[string]interface{}

func (d *DynEnt) Load(props []datastore.Property) error {
	// Note: you might want to clear current values from the map or create a new map
	for _, p := range props {
		(*d)[p.Name] = p.Value
	}
	return nil
}

func (d *DynEnt) Save() (props []datastore.Property, err error) {
	for k, v := range *d {
		props = append(props, datastore.Property{Name: k, Value: v})
	}
	return
}

/* type DynEnt map[string]interface{}

func (d *DynEnt) Load(ch <-chan datastore.Property) error {
	// Note: you might want to clear current values from the map or create a new map
	for p := range ch { // Read until channel is closed
		(*d)[p.Name] = p.Value
	}
	return nil
}

func (d *DynEnt) Save(ch chan<- datastore.Property) {
	defer close(ch) // Channel must be closed
	for k, v := range *d {
		ch <- datastore.Property{Name: k, Value: v}
	}
	return
} */

//GetAccountHandler is to
func GetAccountHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	c := appengine.NewContext(r)

	//d := DynEnt{"email": "me@myhost.com", "time": time.Now()}

	d := &DynEnt{}
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		log.Printf("err ", err)

		return
	}

	log.Println("d=>", d)
	log.Println("&d=>", &d)
	k := datastore.NewIncompleteKey(c, "DynEntity", nil)
	key, err := datastore.Put(c, k, d)
	//	log.Fatalf("%v %v", key, err)
	log.Printf("Key ", key)
	log.Printf("err ", err)

}
