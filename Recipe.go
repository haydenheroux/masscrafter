package main

type Item struct {
	repr string;
	isCompact bool;
}

type Recipe struct {
	outputs ItemList;
	inputs ItemList;
}

type ItemList map[Item]float64;

func GetMaterialAmounts(output Item, recipe Recipe, total float64) ItemList {
	var numOfCrafts float64
	var inputsScaled ItemList

	numOfCrafts = float64(total) / float64(recipe.outputs[output])
	inputsScaled = make(ItemList, len(recipe.inputs))
	for item, count := range recipe.inputs {
		inputsScaled[item] = float64(float64(count) * numOfCrafts)
	}

	return inputsScaled
}

func getMaterialAmountsCompact(input Item, output Item, recipe Recipe, total float64) ItemList {
	var compactionFactor float64
	var compactedMaterials ItemList

	compactedMaterials = make(ItemList)

	compactionFactor = float64(recipe.inputs[output]) / float64(recipe.outputs[input])

	compactedMaterials[output] = float64(float64(total) * compactionFactor)

	return compactedMaterials
}

func GetMaterialAmountsCompact(inputs ItemList) (bool, map[Item]float64) {
	var toBeCompacted, compacted ItemList
	var isFullyCompacted bool

	toBeCompacted = make(ItemList)
	compacted = make(ItemList)

	for item, count := range inputs {
		if item.isCompact == false {
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
		var itemCompact Item
		var itemRecipeCompact Recipe

		itemCompact = GetItemCompact(item)
		itemRecipeCompact = GetItemRecipeCompact(item)

		inputsForItem = getMaterialAmountsCompact(item, itemCompact, itemRecipeCompact, count)
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
		if item.isCompact == false {
			isFullyCompacted = false
		}
	}

	return isFullyCompacted, compacted
}

