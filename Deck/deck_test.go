package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 28 {
		t.Errorf("Expected length of 28, but got", len(d))
	}

}
