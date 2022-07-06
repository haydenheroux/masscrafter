package main

var allItems = Load()

// Materials returns the materials required to craft a number of other items.
func Materials(item string, amount float64) ItemList {
	itemInfo := allItems[item]

	// Handle trivial case where the item is not craftable
	if itemInfo.IsCompact {
		return ItemList{
			item: amount,
		}
	}

	// The materials needed to craft this item
	materials := itemInfo.CraftRecipe.Inputs

	// The amount of items crafted per craft action
	yielded := itemInfo.CraftRecipe.Outputs[item]
	// ... and the number of times needed to craft the desired amount
	crafts := amount / yielded

	// Compute the materials needed to complete the desired crafts
	for material, perCraft := range materials {
		materials[material] = perCraft * crafts
	}

	return materials
}

// Simplify returns the simplified materials required to craft the items.
func Simplify(items ItemList) (ItemList, bool) {
	return ItemList{}, true
}
