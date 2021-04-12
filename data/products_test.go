package data

import "testing"

func TestChecksValidator(t *testing.T) {
	p := &Product{
		Name:  "Test",
		Price: 1.00,
		SKU:   "avs-ad-vdf",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
