package data

import (
	"fmt"
	"time"
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

// Items is a collection of Item
type Items []*Item

// ErrItemNotFound is custom error message when Item not found in DB
var ErrItemNotFound = fmt.Errorf("Item Not Found")

// GetItems returns static collection of Items
func GetItems() Items {
	return itemList
}

// GetItemByID returns a particular Item identified by ID
// Returns ErrItemNotFound when no item with given ID is found
func GetItemByID(id int) (*Item, error) {
	idx, item := findIndexAndItemByID(id)
	if idx == -1 {
		return nil, ErrItemNotFound
	}
	return item, nil
}

// AddNewItem creates a new Item to the Item DB
func AddNewItem(it Item) {
	it.ID = getNextID()
	it.CreatedAt = time.Now().UTC().String()
	it.UpdatedAt = time.Now().UTC().String()
	itemList = append(itemList, &it)
}

// UpdateItem updates an Item with the given ID
func UpdateItem(it Item) error {
	idx, itWas := findIndexAndItemByID(it.ID)

	if idx == -1 {
		return ErrItemNotFound
	}

	it.CreatedAt = itWas.CreatedAt
	it.UpdatedAt = time.Now().UTC().String()
	itemList[idx] = &it
	return nil
}

// DeleteItem removes an Item from the Item DB
func DeleteItem(id int) error {
	idx, _ := findIndexAndItemByID(id)
	if idx == -1 {
		return ErrItemNotFound
	}
	itemList = itemList[:idx+copy(itemList[idx:], itemList[idx+1:])]
	return nil
}

func findIndexAndItemByID(id int) (int, *Item) {
	for idx, item := range itemList {
		if item.ID == id {
			return idx, item
		}
	}
	return -1, nil
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
