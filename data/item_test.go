package data

import "testing"

func TestCheckValidation(t *testing.T) {
	it := &Item{
		SKU:         "abcdefg2AD23",
		VendorCode:  "h28920AcT543",
		Name:        "Burger",
		Description: "Unhealthy Food Item",
		Price:       2.5,
	}
	err := it.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
