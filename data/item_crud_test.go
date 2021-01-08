package data

import (
	"os"
	"testing"

	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/rramesh/eatables/config"
	"github.com/stretchr/testify/assert"
)

var idb *ItemDB

func TestMain(m *testing.M) {
	idb = setup()
	seed(idb)
	code := m.Run()
	teardown(idb)
	os.Exit(code)
}

func TestGetItemsReturnsSeedData(t *testing.T) {
	items := idb.GetItems()
	assert.NotEmpty(t, items)
	assert.Equal(t, "Uthappam", items[0].Name)
}

func TestGetItemByIDNotPresent(t *testing.T) {
	item, err := idb.GetItemByID(1200)
	assert.Nil(t, item)
	assert.Error(t, ErrItemNotFound, err)
}

func TestGetItemByIDSuccess(t *testing.T) {
	item, err := idb.GetItemByID(1)
	assert.Equal(t, "Uthappam", item.Name)
	assert.Nil(t, err)
}

func TestGetItemBySKUNotPresent(t *testing.T) {
	item, err := idb.GetItemBySKU(uuid.New().String())
	assert.Nil(t, item)
	assert.Error(t, ErrItemNotFound, err)
}

func TestGetItemBySKUInvalidUUIDt(t *testing.T) {
	item, err := idb.GetItemBySKU("x2db1163-763a-44b3-a77a-ccd91234ef5b")
	assert.Nil(t, item)
	assert.Error(t, ErrInvalidUUID, err)
}

func TestGetItemBySKUSuccess(t *testing.T) {
	itemExpected, err := idb.GetItemByID(1)
	itemGot, err := idb.GetItemBySKU(itemExpected.SKU)
	assert.Equal(t, itemExpected, itemGot)
	assert.Nil(t, err)
}

func TestGetItemsByVendorCodeNotPresent(t *testing.T) {
	items, err := idb.GetItemByVendorCode(uuid.New().String())
	assert.Empty(t, items)
	assert.Nil(t, err)
}

func TestGetItemsByVendorCodeInvalidUUIDt(t *testing.T) {
	items, err := idb.GetItemByVendorCode("x2db1163-763a-44b3-a77a-ccd91234ef5b")
	assert.Nil(t, items)
	assert.Error(t, ErrInvalidUUID, err)
}

func TestGetItemsByVendorCodeSuccess(t *testing.T) {
	itemExpected, err := idb.GetItemByID(1)
	itemGot, err := idb.GetItemByVendorCode(itemExpected.VendorCode)
	assert.Equal(t, itemExpected, itemGot[0])
	assert.Nil(t, err)
}

func TestAddNewItemSKUPassedError(t *testing.T) {
	item := new(Item)
	item.SKU = uuid.New().String()
	err := idb.AddNewItem(*item)
	assert.Error(t, ErrSKUInCreate, err)
}

func TestAddNewItemDBError(t *testing.T) {
	item := &Item{
		// Name & Vendor Code deliberately left unfilled
		Description:       "South Indian version of Pan Cake",
		Price:             45.85,
		NonVegetarian:     false,
		Cuisine:           "South Indian",
		Category:          []string{"Breakfast", "Dinner"},
		Customizable:      false,
		AvailableTimes:    []TimeRange{{From: 390, To: 690}},
		Tags:              []string{"Dosa"},
		DontMakeItAnymore: false,
	}
	err := idb.AddNewItem(*item)
	assert.NotNil(t, err)
}

func TestAddNewItemSuccess(t *testing.T) {
	item := &Item{
		VendorCode:        uuid.New().String(),
		Name:              "Salad",
		Description:       "Vegetable Salad",
		Price:             25.45,
		NonVegetarian:     false,
		Cuisine:           "World",
		Category:          []string{"Breakfast", "Dinner"},
		Customizable:      false,
		AvailableTimes:    []TimeRange{{From: 390, To: 690}},
		Tags:              []string{"Salad", "Vegetable Salad"},
		DontMakeItAnymore: false,
	}
	err := idb.AddNewItem(*item)
	assert.Nil(t, err)
}

func TestUpdateItemSKUNotFound(t *testing.T) {
	item := &Item{
		SKU: uuid.New().String(),
	}
	err := idb.UpdateItem(*item)
	assert.Error(t, ErrItemNotFound, err)
}

func TestUpdateItemSuccess(t *testing.T) {
	item, _ := idb.GetItemByID(1)
	item.Price = 56.75
	err := idb.UpdateItem(*item)
	assert.Nil(t, err)
	item, _ = idb.GetItemByID(1)
	assert.Equal(t, 56.75, item.Price)
}

func TestDeleteItemSKUNotFound(t *testing.T) {
	sku := uuid.New().String()
	err := idb.DeleteItem(sku)
	assert.Error(t, ErrItemNotFound, err)
}

func TestDeleteItemSuccess(t *testing.T) {
	vc := uuid.New().String()
	item := &Item{
		VendorCode:        vc,
		Name:              "Salad",
		Description:       "Vegetable Salad",
		Price:             25.45,
		NonVegetarian:     false,
		Cuisine:           "World",
		Category:          []string{"Breakfast", "Dinner"},
		Customizable:      false,
		AvailableTimes:    []TimeRange{{From: 390, To: 690}},
		Tags:              []string{"Salad", "Vegetable Salad"},
		DontMakeItAnymore: false,
	}
	idb.AddNewItem(*item)
	items, _ := idb.GetItemByVendorCode(vc)
	sku := items[0].SKU
	err := idb.DeleteItem(sku)
	assert.Nil(t, err)
	item, err = idb.GetItemBySKU(sku)
	assert.Error(t, ErrItemNotFound, err)
}

func setup() *ItemDB {
	l := hclog.Default()
	dbh := NewDBHandle(l)
	conf := new(config.Config)
	conf.Logger = l
	conf.BindAddress = ":9999"
	conf.DB.Hostname = "localhost"
	conf.DB.Port = 5432
	conf.DB.Username = "ramesh"
	conf.DB.Password = "Think@123"
	conf.DB.Database = "eatables_test"
	conf.DB.MaxConns = 1
	conf.Log.Level = "debug"
	if err := dbh.Connect(conf); err != nil {
		panic(err)
	}
	dbh.Init()
	idb := NewItemDB(l, dbh.DB)
	l.Info("Test DB Initialized")
	return idb
}

func seed(idb *ItemDB) {
	item1 := &Item{
		SKU:               uuid.New().String(),
		VendorCode:        uuid.New().String(),
		Name:              "Uthappam",
		Description:       "South Indian version of Pan Cake",
		Price:             45.85,
		NonVegetarian:     false,
		Cuisine:           "South Indian",
		Category:          []string{"Breakfast", "Dinner"},
		Customizable:      false,
		AvailableTimes:    []TimeRange{{From: 390, To: 690}},
		Tags:              []string{"Uthappam", "Dosa"},
		DontMakeItAnymore: false,
	}

	item2 := &Item{
		SKU:               uuid.New().String(),
		VendorCode:        uuid.New().String(),
		Name:              "Rava Dosa",
		Description:       "Dosa made from Rava",
		Price:             40.00,
		NonVegetarian:     false,
		Cuisine:           "South Indian",
		Category:          []string{"Breakfast", "Dinner"},
		Customizable:      false,
		AvailableTimes:    []TimeRange{{From: 390, To: 1350}},
		Tags:              []string{"Rava", "Dosa"},
		DontMakeItAnymore: false,
	}
	items := []*Item{item1, item2}
	_, err := idb.db.Model(&items).Insert()
	if err != nil {
		idb.l.Error("Could not seed Test DB for running tests", "error", err)
		panic("Could not seed Test DB for running tests")
	}
	idb.l.Info("Seed Data Initialized")
}

func teardown(idb *ItemDB) {
	err := idb.db.Model(&Item{}).DropTable(&orm.DropTableOptions{IfExists: true})
	if err != nil {
		idb.l.Error("Could not truncate Test DB after running tests", "error", err)
		panic("Could not truncate Test DB after running tests")
	}
	idb.l.Info("Test DB teardown success")
}
