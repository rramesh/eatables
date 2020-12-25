package data

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/hashicorp/go-hclog"
)

// Item defines the structure for an API Food Item
// swagger:model
type Item struct {
	tableName struct{} `pg:"items,alias:item"`

	// The ID of this Item
	// required: false
	// min: 1
	ID int `pg:",pk" json:"id"`

	// The SKU of this Item. UUID format
	// required: true
	// pattern: [a-zA-Z0-9]{36}
	// example: b5113148-d1fc-4c17-8177-519120495b4c
	SKU string `pg:",type:uuid,notnull,unique" json:"sku"`

	// The Vendor Code of this Item. UUID format
	// required: true
	// pattern: [a-zA-Z0-9]{36}
	// example: ee846edd-b2ee-4ab2-bd97-2c4246c56cf5
	VendorCode string `pg:",type:uuid,notnull" json:"vendorCode"`

	// Name of this Item
	// required: true
	// example: Masala Dosa
	Name string `pg:",notnull" json:"name" validate:"required"`

	// Description of this Item
	// required: true
	// example: Made from rice, lentils, potato, fenugreek, and curry leaves, and served with chutneys and sambar.
	Description string `pg:",notnull" json:"description" validate:"required"`

	// Price of this Item
	// required: true
	// min: 0.01
	// example: 75.00
	Price float64 `pg:",notnull" json:"price" validate:"gt=0.0"`

	// Whether this Item is Non-vegetarian
	// Defaults to False if not provided - Item is Vegetarian by default
	// example: false
	NonVegetarian bool `pg:",notnull,default:false" json:"nonVegetarian"`

	// Cuisine this Item belongs to
	// example: South Indian
	//          Chinese
	Cuisine string `json:"cuisine"`

	// Category this Item belongs to, array of Strings. Used for grouping Items under menu
	// example: ["Breakfast", "Dinner"]
	//          ["Snacks", "Anytime"]
	Category []string `pg:",array" json:"category"`

	// Is the Item Customizable. Defaults to False
	// example: false
	Customizable bool `pg:",default:false" json:"customizable"`

	// What times this item is available.
	// Range provided as Array of Array of two strings
	// example: [{from: "7:00", to: "11:30"}, {from: "17:00", to: "22:30"}]
	AvailableTimes []TimeRange `pg:"type:jsonb" json:"availableTimes"`

	// Tags to be associated with this Item.
	// Helpful as search keywords
	// example: ["Yummy", "South Indian", "Dosa", "Special Dosa"]
	Tags []string `pg:",array" json:"tags"`

	// Is the Item still made? Active?
	// Defaults to False, meaning Item is still being made and active
	// example: false
	DontMakeItAnymore bool `pg:",default:false" json:"dontMakeItAnymore"`

	CreatedAt time.Time `pg:",default:now()" json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// TimeRange holds a starting and ending time
type TimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Items is a collection of Item
type Items []*Item

// ItemDB is the interface to DB methods
type ItemDB struct {
	l  hclog.Logger
	db *pg.DB
}

// NewItemDB creates an instance of ItemDB
func NewItemDB(l hclog.Logger, db *pg.DB) *ItemDB {
	return &ItemDB{l, db}
}

// ErrItemNotFound is custom error message when Item not found in DB
var ErrItemNotFound = fmt.Errorf("No Items Found")

// GetItems returns static collection of Items
func (i *ItemDB) GetItems() Items {
	var items []*Item
	i.db.Model(&items).Select()
	return items
	// return itemList
}

// GetItemByID returns a particular Item identified by ID
// This can be used for internal calls where record ID of the item is known
// Returns ErrItemNotFound when no item with given ID is found
func (i *ItemDB) GetItemByID(id int) (*Item, error) {
	idx, item := findIndexAndItemByID(id)
	if idx == -1 {
		return nil, ErrItemNotFound
	}
	return item, nil
}

// GetItemBySKU returns a particular Item identified by SKU
// This can be used for other services or UI to call as
// SKU of item alone is exposed and not record ID
// Returns ErrItemNotFound when no item with given ID is found
func (i *ItemDB) GetItemBySKU(uuid string) (*Item, error) {
	idx, item := findIndexAndItemBySKU(uuid)
	if idx == -1 {
		return nil, ErrItemNotFound
	}
	return item, nil
}

// GetItemByVendorCode returns list of Items identified by Vendor UUID
// Returns ErrItemNotFound when no item with given ID is found
func (i *ItemDB) GetItemByVendorCode(uuid string) ([]*Item, error) {
	items := findItemsByVendorCode(uuid)
	if len(items) == 0 {
		return nil, ErrItemNotFound
	}
	return items, nil
}

// AddNewItem creates a new Item to the Item DB
func (i *ItemDB) AddNewItem(it Item) {
	_, err := i.db.Model(&it).Insert()
	if err != nil {
		i.l.Error("Error inserting item into DB", "error", err)
		return
	}
	// it.ID = getNextID()
	// it.CreatedAt = time.Now().UTC().String()
	// it.UpdatedAt = time.Now().UTC().String()
	// itemList = append(itemList, &it)
}

// UpdateItem updates an Item with the given ID
func (i *ItemDB) UpdateItem(it Item) error {
	idx, itWas := findIndexAndItemBySKU(it.SKU)

	if idx == -1 {
		return ErrItemNotFound
	}
	// Do not allow change of ID and SKU
	it.ID = itWas.ID
	it.SKU = itWas.SKU
	it.CreatedAt = itWas.CreatedAt
	it.UpdatedAt = time.Now()
	itemList[idx] = &it
	return nil
}

// DeleteItem removes an Item from the Item DB
func (i *ItemDB) DeleteItem(id int) error {
	idx, _ := findIndexAndItemByID(id)
	if idx == 0 {
		return ErrItemNotFound
	}
	itemList = itemList[:int(idx)+copy(itemList[idx:], itemList[idx+1:])]
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

func findIndexAndItemBySKU(uuid string) (int, *Item) {
	for idx, item := range itemList {
		if item.SKU == uuid {
			return idx, item
		}
	}
	return -1, nil
}

func findItemsByVendorCode(uuid string) []*Item {
	var items = []*Item{}
	for _, item := range itemList {
		if item.VendorCode == uuid {
			items = append(items, item)
		}
	}
	return items
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
		SKU:            "jj3d8dk3mk9x",
		VendorCode:     "dd33gk988kdx",
		Name:           "Masal Dosai",
		Description:    "Rice batter toasted circular on tava filled with mashed potatoes",
		Price:          43.75,
		NonVegetarian:  false,
		Cuisine:        "South Indian",
		Category:       []string{"Breakfast", "Dinner"},
		Customizable:   false,
		AvailableTimes: []TimeRange{{From: "6:00", To: "11:00"}, {From: "17:00", To: "22:30"}},
		Tags:           []string{"Dosa", "Masal Dosa", "South Indian", "Dosai", "Masala Dosai", "Masala"},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	},
}
