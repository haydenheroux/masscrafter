package main

var allItems = Load()

// Materials returns the materials required to craft a number of other items.
func Materials(item string, amount float64) ItemList {
	itemInfo := allItems[item]

	// Handle trivial case where the item is not craftable
	if itemInfo.IsSimplified {
		return ItemList{
			item: amount,
		}
	}

	// The materials needed to craft this item
	inputMaterials := itemInfo.CraftRecipe.Inputs

	// The amount of items crafted per craft action
	yielded := itemInfo.CraftRecipe.Outputs[item]
	// ... and the number of times needed to craft the desired amount
	crafts := amount / yielded

	// Store required materials in here
	materials := make(ItemList)

	// Compute the materials needed to complete the desired crafts
	for material, perCraft := range inputMaterials {
		materials[material] = perCraft * crafts
	}

	return materials
}

// Simplify returns the simplified materials required to craft the items.
func Simplify(items ItemList) (ItemList, bool) {
	materials := make(ItemList)

	// Compute the total materials needed to craft the items
	for item, amount := range items {
		required := Materials(item, amount)

		// Add the required materials to the totals
		for rItem, rAmount := range required {
			materials[rItem] += rAmount
		}
	}

	// Flag for if the materials cannot be simplified further
	done := true

	for item, amount := range materials {
		// Check if there is no way to craft the materials
		// Materials will be identical if there is no other way to craft
		required := Materials(item, amount)
		if required[item] != amount {
			// If the materials can be crafted from something else, not done
			done = false
			break
		}
	}

	return materials, done
}
