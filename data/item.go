package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Item structure holds basic details of the Food Item
type Item struct {
	ID             int              `json:"id"`
	SKU            string           `json:"sku"`
	VendorCode     string           `json:"vendorCode"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Price          string           `json:"price"`
	NonVegetarian  bool             `json:"nonVegetarian"`
	Cuisine        string           `json:"cuisine"`
	Category       string           `json:"category"`
	Customizable   bool             `json:"customizable"`
	AvailableTimes []([2]time.Time) `json:"availableTimes"`
	Tags           []string         `json:"tags"`
}

// FromJSON transforms a JSON based item to Item structure
func (it *Item) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(it)
}

// Items holds the current Item list
type Items []*Item

// ToJSON converts a Item list structure to a JSON
func (it *Items) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(it)
}

// GetItems returns static list of Items
func GetItems() Items {
	return itemList
}

// AddItem adds a new Item to the Item list
func AddItem(it *Item) {
	it.ID = getNextID()
	itemList = append(itemList, it)
}

// UpdateItem updates a Item in the list
func UpdateItem(id int, it *Item) error {
	_, pos, err := findItem(id)

	if err != nil {
		return err
	}
	it.ID = id
	itemList[pos] = it
	return nil
}

// ErrItemNotFound is custom error message when Item not found in DB
var ErrItemNotFound = fmt.Errorf("Item Not Found")

func findItem(id int) (*Item, int, error) {
	for i, p := range itemList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrItemNotFound
}

// GetNextID picks up the last element in the Item list and adds 1 to the ID value
func getNextID() int {
	lp := itemList[len(itemList)-1]
	return lp.ID + 1
}

var itemList = []*Item{}
