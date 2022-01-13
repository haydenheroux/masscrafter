package main

import (
	"testing"
)

func TestCraftPlanks(t* testing.T) {
	var materials ItemList = GetMaterialAmounts("woodPlank", 64)
	var logs = materials["woodLog"]
	if logs != 16.0 {
		t.Errorf("got: %f, want: %f", logs, 16.0)
	}
}

func TestCraftHoppers(t* testing.T) {
	var materials ItemList = GetMaterialAmounts("hopper", 64.0)
	var ironIngots = materials["ironIngot"]
	var chests = materials["chest"]
	if ironIngots != 5 * 64.0 || chests != 1 * 64.0 {
		t.Errorf("got: (iron ingots %f, chests %f), want: (iron ingots %f, chests %f)", ironIngots, chests, 5 * 64.0, 1 * 64.0)
	}
}

func TestCompactCraftHoppers(t* testing.T) {
	var materials ItemList = GetMaterialAmounts("hopper", 64)
	var isFullyCompacted bool
	for {
		isFullyCompacted, materials = GetMaterialAmountsCompact(materials)
		if isFullyCompacted {
			break
		}
	}
	var expectedMaterials = ItemList {
		"ironBlock": (5.0 * 64.0) / 9.0,
		"woodLog": (8.0 * 64.0) / 4.0,
	}
	if materials["ironBlock"] != expectedMaterials["ironBlock"] || materials["woodLog"] != materials["woodLog"] {
		t.Errorf("got: (iron blocks %f wood logs %f), want (iron blocks %f wood logs %f)", materials["ironBlock"], expectedMaterials["ironBlock"], materials["woodLog"], expectedMaterials["woodLog"])
	}
}
