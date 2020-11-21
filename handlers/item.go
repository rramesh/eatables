package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/rramesh/eatables/data"
)

// Items handler struct contains logger reference
type Items struct {
	l *log.Logger
}

//NewItems creates a new handler
func NewItems(l *log.Logger) *Items {
	return &Items{l}
}

func (items *Items) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		items.getItems(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		items.addItem(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		rgx := regexp.MustCompile(`([0-9]+)`)
		nos := rgx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(nos) != 1 {
			items.l.Println("Invalid URI, more than 1 ID")
			http.Error(rw, "Invalid URL, must specifiy a proper ID", http.StatusBadRequest)
			return
		}

		if len(nos[0]) != 2 {
			items.l.Println("Invalid URI, more than 1 Capture Group")
			http.Error(rw, "Invalid URL, must specifiy a proper ID", http.StatusBadRequest)
			return
		}
		idString := nos[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			items.l.Println("Invalid URI, ID not Parseable")
			http.Error(rw, "Unable to extract ID", http.StatusBadRequest)
			return
		}
		items.updateItem(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (items *Items) getItems(rw http.ResponseWriter, r *http.Request) {
	items.l.Println("Handle Get Items")
	itemList := data.GetItems()
	err := itemList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to parse list of items as JSON", http.StatusInternalServerError)
	}
}

func (items *Items) addItem(rw http.ResponseWriter, r *http.Request) {
	items.l.Println("Handle POST Item")
	it := &data.Item{}
	err := it.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall JSON", http.StatusBadRequest)
	}
	data.AddItem(it)
	items.l.Printf("Prod: %#v", *it)
}

func (items *Items) updateItem(id int, rw http.ResponseWriter, r *http.Request) {
	items.l.Println("Handle PUT Item")
	it := &data.Item{}
	err := it.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall JSON", http.StatusBadRequest)
	}
	err = data.UpdateItem(id, it)
	if err == data.ErrItemNotFound {
		http.Error(rw, "Item Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Item Not Found", http.StatusInternalServerError)
		return
	}
	items.l.Printf("Prod: %#v", *it)
}
