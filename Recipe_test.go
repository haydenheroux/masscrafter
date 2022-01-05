package main

import (
	"testing"
)

func TestCraftPlanks(t* testing.T) {
	var materials map[Item]float64 = GetMaterialAmounts(WoodPlank, WoodPlankRecipe, 64)
	var logs = materials[WoodLog]
	if logs != 16.0 {
		t.Errorf("got: %f, want: %f", logs, 16.0)
	}
}

func TestCraftHoppers(t* testing.T) {
	var materials map[Item]float64 = GetMaterialAmounts(Hopper, HopperRecipe, 64.0)
	var ironIngots = materials[IronIngot]
	var chests = materials[Chest]
	if ironIngots != 5 * 64.0 || chests != 1 * 64.0 {
		t.Errorf("got: (iron ingots %f, chests %f), want: (iron ingots %f, chests %f)", ironIngots, chests, 5 * 64.0, 1 * 64.0)
	}
}

func TestCompactCraftHoppers(t* testing.T) {
	var materials map[Item]float64 = GetMaterialAmounts(Hopper, HopperRecipe, 64)
	var isFullyCompacted bool
	for {
		isFullyCompacted, materials = GetMaterialAmountsCompact(materials)
		if isFullyCompacted {
			break
		}
	}
	var expectedMaterials = map[Item]float64 {
		IronBlock: (5.0 * 64.0) / 9.0,
		WoodLog: (8.0 * 64.0) / 4.0,
	}
	if materials[IronBlock] != expectedMaterials[IronBlock] || materials[WoodLog] != materials[WoodLog] {
		t.Errorf("got: (iron blocks %f wood logs %f), want (iron blocks %f wood logs %f)", materials[IronBlock], expectedMaterials[IronBlock], materials[WoodLog], expectedMaterials[WoodLog])
	}
}
