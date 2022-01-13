package main

import (
	"testing"
)

func TestLoad(t *testing.T) {
	allItems := Load()	

	woodPlank := allItems["woodPlank"]
	if woodPlank.Repr != "Wood Plank" || woodPlank.CraftRecipe.Outputs["woodPlank"] != 4.0 {
		t.Errorf("got: (repr: %s, outputs: %f), want: (repr: %s, outputs: %f)", woodPlank.Repr, woodPlank.CraftRecipe.Outputs["woodPlank"], "Wood Plank", 4.0)
	}
}
