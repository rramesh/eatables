package data

import (
	protos "github.com/rramesh/eatables/protos/items"
)

// ToTimeRange converts protos TimeRange to data TimeRange
func ToTimeRange(tr []*protos.TimeRange) []TimeRange {
	var trange []TimeRange
	for _, v := range tr {
		trg := &TimeRange{
			From: v.From,
			To:   v.To,
		}
		trange = append(trange, *trg)
	}
	return trange
}

// FromTimeRange converts data TimeRange to protos TimeRange
func FromTimeRange(tr []TimeRange) []*protos.TimeRange {
	var trange []*protos.TimeRange
	for _, v := range tr {
		trg := &protos.TimeRange{
			From: v.From,
			To:   v.To,
		}
		trange = append(trange, trg)
	}
	return trange
}

// FromItem converts item data to protos Item Details
func FromItem(it Item) *protos.ItemDetails {
	return &protos.ItemDetails{
		Id:                int32(it.ID),
		Sku:               it.SKU,
		VendorCode:        it.VendorCode,
		Name:              it.Name,
		Description:       it.Description,
		Price:             float32(it.Price),
		NonVegetarian:     it.NonVegetarian,
		Cuisine:           it.Cuisine,
		Category:          it.Category,
		Customizable:      it.Customizable,
		AvailableTimes:    FromTimeRange(it.AvailableTimes),
		Tags:              it.Tags,
		DontMakeItAnymore: it.DontMakeItAnymore,
	}
}

//ToItem converts protos Item Details to Item Data
func ToItem(req *protos.CreateOrUpdateRequest) *Item {
	return &Item{
		SKU:               req.GetSku(),
		VendorCode:        req.GetVendorCode(),
		Name:              req.GetName(),
		Description:       req.GetDescription(),
		Price:             float64(req.GetPrice()),
		NonVegetarian:     req.GetNonVegetarian(),
		Cuisine:           req.GetCuisine(),
		Category:          req.GetCategory(),
		Customizable:      req.GetCustomizable(),
		AvailableTimes:    ToTimeRange(req.GetAvailableTimes()),
		Tags:              req.GetTags(),
		DontMakeItAnymore: false,
	}
}

// FromItems converts Item list data to protos Item Details list
func FromItems(its Items) []*protos.ItemDetails {
	var itr = []*protos.ItemDetails{}
	for _, v := range its {
		itr = append(itr, FromItem(*v))
	}
	return itr
}
