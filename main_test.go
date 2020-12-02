package main

import (
	"testing"

	"github.com/rramesh/eatables/sdk/client"
	"github.com/rramesh/eatables/sdk/client/items"
	"github.com/stretchr/testify/assert"
)

// Client SDK generated through swagger is used for integration testing
// sdk/>swagger generate client -f ../swagger.yml -A eatables
// Run the server and run these tests
func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := items.NewListItemsParams()
	it, err := c.Items.ListItems(params)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Masal Dosai", *it.GetPayload()[0].Name)
	// fmt.Printf("%#v", it.GetPayload()[0])
}
