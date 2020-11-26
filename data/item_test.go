package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemMissingNameReturnsErr(t *testing.T) {
	it := &Item{
		SKU:         "abcdefg2AD23",
		VendorCode:  "h28920AcT543",
		Description: "Unhealthy Food Item",
		Price:       2.5,
	}
	v := NewValidation()
	err := v.Validate(it)
	assert.Len(t, err, 1)
}

func TestMissingDescriptionReturnsErr(t *testing.T) {
	it := &Item{
		SKU:        "abcdefg2AD23",
		VendorCode: "h28920AcT543",
		Name:       "Burger",
		Price:      2.5,
	}
	v := NewValidation()
	err := v.Validate(it)
	assert.Len(t, err, 1)
}

func TestInvalidPriceReturnsErr(t *testing.T) {
	it := &Item{
		SKU:         "abcdefg2AD23",
		VendorCode:  "h28920AcT543",
		Name:        "Burger",
		Description: "Unhealthy Food Item",
		Price:       -1,
	}
	v := NewValidation()
	err := v.Validate(it)
	assert.Len(t, err, 1)
}

func TestInvalidSKUReturnsErr(t *testing.T) {
	it := &Item{
		SKU:         "abcd",
		VendorCode:  "h28920AcT543",
		Name:        "Burger",
		Description: "Unhealthy Food Item",
		Price:       2.5,
	}
	v := NewValidation()
	err := v.Validate(it)
	assert.Len(t, err, 1)
}

func TestInvalidVendorCodeeturnsErr(t *testing.T) {
	it := &Item{
		SKU:         "abcdefg2AD23",
		VendorCode:  "T543",
		Name:        "Burger",
		Description: "Unhealthy Food Item",
		Price:       2.5,
	}
	v := NewValidation()
	err := v.Validate(it)
	assert.Len(t, err, 1)
}

func TestValidItemDoesNotReturnErr(t *testing.T) {
	it := &Item{
		SKU:         "abcdefg2AD23",
		VendorCode:  "h28920AcT543",
		Name:        "Burger",
		Description: "Unhealthy Food Item",
		Price:       2.5,
	}
	v := NewValidation()
	err := v.Validate(it)
	assert.Len(t, err, 0)
}

func TestItemsToJSON(t *testing.T) {
	items := []*Item{
		{
			SKU:         "abcdefg2AD23",
			VendorCode:  "h28920AcT543",
			Name:        "Burger",
			Description: "Unhealthy Food Item",
			Price:       2.5,
		},
	}
	b := bytes.NewBufferString("")
	err := ToJSON(items, b)
	assert.NoError(t, err)
}
