package main

var AllItems = Load()

func GetMaterialAmounts(outputItem string, total float64) ItemList {
	var numOfCrafts float64
	var inputsScaled ItemList

	output := AllItems[outputItem]
	recipe := *output.CraftRecipe

	numOfCrafts = float64(total) / float64(recipe.Outputs[outputItem])
	inputsScaled = make(ItemList, len(recipe.Inputs))
	for item, count := range recipe.Inputs {
		inputsScaled[item] = float64(float64(count) * numOfCrafts)
	}

	return inputsScaled
}

func getMaterialAmountsCompact(inputItem string, outputItem string, total float64) ItemList {
	var compactionFactor float64
	var compactedMaterials ItemList

	input := AllItems[inputItem]
	recipe := *input.CraftRecipe

	compactedMaterials = make(ItemList)

	compactionFactor = float64(recipe.Inputs[outputItem]) / float64(recipe.Outputs[inputItem])

	compactedMaterials[outputItem] = float64(float64(total) * compactionFactor)

	return compactedMaterials
}

func GetMaterialAmountsCompact(inputs ItemList) (bool, ItemList) {
	var toBeCompacted, compacted ItemList
	var isFullyCompacted bool

	toBeCompacted = make(ItemList)
	compacted = make(ItemList)


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

	for item, count := range toBeCompacted {
		var inputsForItem ItemList
		var itemCompact string

		compactOutputs := AllItems[item].CraftRecipe.Inputs
		for item := range compactOutputs {
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

	isFullyCompacted = true
	for item, _ := range compacted {
		if AllItems[item].IsCompact == false {
			isFullyCompacted = false
		}
	}

	return isFullyCompacted, compacted
}

