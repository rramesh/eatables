package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Item defines the structure for an API Food Item
type Item struct {
	ID                int           `json:"id"`
	SKU               string        `json:"sku" validate:"required,uuid"`
	VendorCode        string        `json:"vendorCode" validate:"required,uuid"`
	Name              string        `json:"name" validate:"required"`
	Description       string        `json:"description" validate:"required"`
	Price             float32       `json:"price" validate:"gt=0.0"`
	NonVegetarian     bool          `json:"nonVegetarian"`
	Cuisine           string        `json:"cuisine"`
	Category          []string      `json:"category"`
	Customizable      bool          `json:"customizable"`
	AvailableTimes    []([2]string) `json:"availableTimes"`
	Tags              []string      `json:"tags"`
	DontMakeItAnymore bool          `json:"dontMakeItAnymore"`
	CreatedAt         string        `json:"-"`
	UpdatedAt         string        `json:"-"`
}

// FromJSON transforms a JSON based item to Item structure
func (it *Item) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(it)
}

// Validate the Item
func (it *Item) Validate() error {
	v := validator.New()
	v.RegisterValidation("uuid", validateUUID)
	return v.Struct(it)
}

//Validate UUID
func validateUUID(fl validator.FieldLevel) bool {
	//UUID must be 12 length alphanumeric
	re := regexp.MustCompile(`[a-zA-Z0-9]{12}`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

// Items is a collection of Item
type Items []*Item

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (it *Items) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(it)
}

// GetItems returns static collection of Items
func GetItems() Items {
	return itemList
}

// AddItem adds a new Item to the Item collection
func AddItem(it *Item) {
	it.ID = getNextID()
	it.CreatedAt = time.Now().UTC().String()
	it.UpdatedAt = time.Now().UTC().String()
	itemList = append(itemList, it)
}

// UpdateItem updates a Item in the collection
func UpdateItem(id int, it *Item) error {
	itWas, pos, err := findItem(id)

	if err != nil {
		return err
	}
	it.ID = id
	it.CreatedAt = itWas.CreatedAt
	it.UpdatedAt = time.Now().UTC().String()
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

// itemList is a hard coded list of items for this
// example data source
var itemList = []*Item{
	{
		ID:             1,
		SKU:            "jj3d8dk3mk9",
		VendorCode:     "dd33gk988kdx",
		Name:           "Masala Dosai",
		Description:    "Rice batter toasted circular on tava filled with mashed potatoes",
		Price:          43.75,
		NonVegetarian:  false,
		Cuisine:        "South Indian",
		Category:       []string{"Breakfast", "Dinner"},
		Customizable:   false,
		AvailableTimes: [][2]string{{"6:00", "11:00"}, {"17:00", "22:30"}},
		Tags:           []string{"Dosa", "Masal Dosa", "South Indian", "Dosai"},
		CreatedAt:      time.Now().UTC().String(),
		UpdatedAt:      time.Now().UTC().String(),
	},
}
