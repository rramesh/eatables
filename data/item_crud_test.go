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

func TestGetItemsReturnsEmpty(t *testing.T) {
	items := idb.GetItems()
	assert.NotEmpty(t, items)
	assert.Equal(t, "Uthappam", items[0].Name)
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
	dbh.Init(true)
	idb := NewItemDB(l, dbh.DB)
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
	_, err := idb.db.Model(item1).Insert()
	if err != nil {
		idb.l.Error("Could not seed Test DB for running tests", "error", err)
		panic("Could not seed Test DB for running tests")
	}
}

func teardown(idb *ItemDB) {
	err := idb.db.Model(&Item{}).DropTable(&orm.DropTableOptions{IfExists: true})
	if err != nil {
		idb.l.Error("Could not truncate Test DB after running tests", "error", err)
		panic("Could not truncate Test DB after running tests")
	}
}
