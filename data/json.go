package data

import (
	"encoding/json"
	"io"
)

// FromJSON transforms a JSON based item to Item structure
func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

// ToJSON serializes the contents of the collection to JSON
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
