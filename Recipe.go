package main

type Item struct {
	repr string;
	isCompact bool;
}

type Recipe struct {
	outputs map[Item]float64;
	inputs map[Item]float64;
}

func GetRequiredInputs(output Item, recipe Recipe, total float64) map[Item]float64 {
	var numOfCrafts float64
	var inputsScaled map[Item]float64

	numOfCrafts = float64(total) / float64(recipe.outputs[output])
	inputsScaled = make(map[Item]float64, len(recipe.inputs))
	for item, count := range recipe.inputs {
		inputsScaled[item] = float64(float64(count) * numOfCrafts)
	}

	return inputsScaled
}

func GetCompactedInputs(input Item, output Item, recipe Recipe, total float64) map[Item]float64 {
	var compactionFactor float64
	var compactedInputs map[Item]float64

	compactedInputs = make(map[Item]float64)

	compactionFactor = float64(recipe.inputs[output]) / float64(recipe.outputs[input])

	compactedInputs[output] = float64(float64(total) * compactionFactor)

	return compactedInputs
}

func GetItemCompact(item Item) Item {
	if item == WoodPlank {
		return WoodLog
	} else if item == Chest {
		return WoodPlank
	} else if item == RedstoneDust {
		return RedstoneBlock
	} else if item == IronIngot {
		return IronBlock
	}
	return Null
}

func GetCompactItemRecipe(item Item) Recipe {
	if item == WoodPlank {
		return WoodPlankRecipe
	} else if item == Chest {
		return ChestRecipe
	} else if item == RedstoneDust {
		return RedstoneDustRecipe
	} else if item == IronIngot {
		return IronIngotRecipe
	}
	return NullRecipe
}

func GetCompactInputs(inputs map[Item]float64) (bool, map[Item]float64) {
	var toBeCompacted, compacted map[Item]float64
	var isFullyCompacted bool

	toBeCompacted = make(map[Item]float64)
	compacted = make(map[Item]float64)

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
		var inputsForItem map[Item]float64
		var itemCompact Item
		var itemCompactRecipe Recipe

		itemCompact = GetItemCompact(item)
		itemCompactRecipe = GetCompactItemRecipe(item)

		inputsForItem = GetCompactedInputs(item, itemCompact, itemCompactRecipe, count)
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

