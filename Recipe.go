package main

var AllItems = Load()

// GetMaterialAmounts computes the amount of items required to craft total
// number of outputItems.
func GetMaterialAmounts(outputItem string, total float64) ItemList {
	output := AllItems[outputItem]
	recipe := *output.CraftRecipe

	// Compute the number of crafting interactions
	numOfCrafts := float64(total) / float64(recipe.Outputs[outputItem])

	// Scale vector of items to compute the total number per item
	inputsScaled := make(ItemList, len(recipe.Inputs))
	for item, count := range recipe.Inputs {
		inputsScaled[item] = float64(float64(count) * numOfCrafts)
	}

	return inputsScaled
}

// getMaterialAmountsCompact transforms total number of inputItems into
// some number of the more compact item type outputItem.
func getMaterialAmountsCompact(inputItem string, outputItem string, total float64) ItemList {
	input := AllItems[inputItem]
	recipe := *input.CraftRecipe

	compactedMaterials := make(ItemList)

	// TODO: Possibly pre-calculate the compaction factor
	compactionFactor := float64(recipe.Inputs[outputItem]) / float64(recipe.Outputs[inputItem])

	compactedMaterials[outputItem] = float64(float64(total) * compactionFactor)

	return compactedMaterials
}

func GetMaterialAmountsCompact(inputs ItemList) (bool, ItemList) {
	toBeCompacted := make(ItemList)
	compacted := make(ItemList)

	// Separate item(s) by need for compaction
	for item, count := range inputs {
		if AllItems[item].IsCompact == false {
			_, exists := toBeCompacted[item]
			if exists {
				toBeCompacted[item] += count
			} else {
				toBeCompacted[item] = count
			}
		} else {
			_, exists := compacted[item]
			if exists {
				compacted[item] += count
			} else {
				compacted[item] = count
			}
		}
	}

	// Perform compaction procedure on all item(s) which need compaction
	for item, count := range toBeCompacted {
		var inputsForItem ItemList
		var itemCompact string

		// Find which item(s) the recipe originates from
		compactOutputs := AllItems[item].CraftRecipe.Inputs
		for item := range compactOutputs {
			// TODO: "Compact"-ness could be computed at runtime
			// as whether or not an item meets the following criteria:
			// 1. The item, as input to a recipe, produces a greater amount of another item
			// 2. The item exists in only one recipe
			itemCompact = item
		}

		inputsForItem = getMaterialAmountsCompact(item, itemCompact, count)
		for input, inputCount := range inputsForItem {
			_, exists := compacted[input]
			if exists {
				compacted[input] += inputCount
			} else {
				compacted[input] = inputCount
			}
		}
	}

	// Test whether all items are fully compacted by iteration
	isFullyCompacted := true
	for item, _ := range compacted {
		if AllItems[item].IsCompact == false {
			isFullyCompacted = false
		}
	}

	return isFullyCompacted, compacted
}
