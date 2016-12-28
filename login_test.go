package funimation

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c.Client.Jar == nil {
		t.Errorf("A new client should have a non nil cookie jar!")
	}
}
