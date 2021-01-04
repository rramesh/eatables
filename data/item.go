package data

import (
	"time"
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
	NonVegetarian bool `pg:",default:false" json:"nonVegetarian"`

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
	// Range provided as Array of map consisiting of from and to times as integers represented in minutes
	// example: [{From: 360, To: 660}, {From: 1020, To: 1350}]
	AvailableTimes []TimeRange `pg:"type:jsonb" json:"availableTimes" validate:"dive,required"`

	// Tags to be associated with this Item.
	// Helpful as search keywords
	// example: ["Yummy", "South Indian", "Dosa", "Special Dosa"]
	Tags []string `pg:",array" json:"tags"`

	// Is the Item still made? Active?
	// Defaults to False, meaning Item is still being made and active
	// example: false
	DontMakeItAnymore bool `pg:",default:false" json:"dontMakeItAnymore"`

	CreatedAt time.Time `pg:",default:now()" json:"-"`
	UpdatedAt time.Time `pg:",default:now()" json:"-"`
}

// TimeRange holds a starting and ending time
type TimeRange struct {
	From uint32 `json:"from" validate:"required,gte=0,lte=1440"`
	To   uint32 `json:"to" validate:"required,gte=0,lte=1440,gtfield=From"`
}

// Items is a collection of Item
type Items []*Item
