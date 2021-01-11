package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
)

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
var ErrItemNotFound = fmt.Errorf("Could not find Item to process")

// ErrSKUInCreate is custom error message when Item not found in DB
var ErrSKUInCreate = fmt.Errorf("SKU provided to add new Item. Provide data without SKU")

// ErrInvalidUUID is custom error message when UUID passed is invalid
var ErrInvalidUUID = fmt.Errorf("Invalid UUID format or value")

// GetItems returns static collection of Items
func (i *ItemDB) GetItems() Items {
	var items Items
	i.db.Model(&items).Select()
	return items
}

// GetItemByID returns a particular Item identified by ID
// This can be used for internal calls where record ID of the item is known
// Returns ErrItemNotFound when no item with given ID is found
func (i *ItemDB) GetItemByID(id int) (*Item, error) {
	item := new(Item)
	err := i.db.Model(item).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, ErrItemNotFound
		}
		i.l.Error("Error finding item by ID", "error", err)
		return nil, err
	}
	return item, nil
}

// GetItemBySKU returns a particular Item identified by SKU
// This can be used for other services or UI to call as
// SKU of item alone is exposed and not record ID
// Returns ErrItemNotFound when no item with given ID is found
func (i *ItemDB) GetItemBySKU(sku string) (*Item, error) {
	_, err := uuid.Parse(sku)
	if err != nil {
		i.l.Error("Error validating UUID", "error", err)
		return nil, ErrInvalidUUID
	}
	item := new(Item)
	err = i.db.Model(item).Where("sku = ?", sku).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, ErrItemNotFound
		}
		i.l.Error("Error finding item by SKU", "error", err)
		return nil, err
	}
	return item, nil
}

// GetItemByVendorCode returns list of Items identified by Vendor UUID
// Returns ErrItemNotFound when no item with given ID is found
func (i *ItemDB) GetItemByVendorCode(vc string) (Items, error) {
	_, err := uuid.Parse(vc)
	if err != nil {
		i.l.Error("Error validating UUID", "error", err)
		return nil, ErrInvalidUUID
	}
	var items Items
	err = i.db.Model(&items).Where("vendor_code = ?", vc).Select()
	return items, err
}

// AddNewItem creates a new Item to the Item DB
func (i *ItemDB) AddNewItem(it Item) (string, error) {
	if it.SKU != "" {
		return "", ErrSKUInCreate
	}
	ctx := context.Background()
	tx, err := i.db.Begin()
	defer tx.Close()
	if err = func(tx *pg.Tx, ctx context.Context) error {
		it.SKU = uuid.New().String()
		_, err := i.db.Model(&it).Insert()
		return err
	}(tx, ctx); err != nil {
		i.l.Error("Error inserting item into DB", "error", err)
		_ = tx.Rollback()
		return "", err
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return it.SKU, err
}

// UpdateItem updates an Item with the given ID
func (i *ItemDB) UpdateItem(it Item) error {
	it.UpdatedAt = time.Now()
	ctx := context.Background()
	tx, err := i.db.Begin()
	defer tx.Close()
	if err = func(tx *pg.Tx, ctx context.Context) error {
		res, err := i.db.Model(&it).Where("sku = ?", it.SKU).UpdateNotZero()
		if res.RowsAffected() == 0 {
			err = ErrItemNotFound
		}
		return err
	}(tx, ctx); err != nil {
		i.l.Error("Error updating Item to DB", "error", err)
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return err
}

// DeleteItem removes an Item from the DB
func (i *ItemDB) DeleteItem(sku string) error {
	ctx := context.Background()
	tx, err := i.db.Begin()
	defer tx.Close()
	if err = func(tx *pg.Tx, ctx context.Context) error {
		res, err := i.db.Model(&Item{}).Where("sku = ?", sku).Delete()
		if res.RowsAffected() == 0 {
			err = ErrItemNotFound
		}
		return err
	}(tx, ctx); err != nil {
		i.l.Error("Error deleting Item from DB", "error", err)
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return err
}
