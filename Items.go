package main

var Null = Item {
	repr: "Null",
	isCompact: true,
}

var NullRecipe = Recipe {
	outputs: map[Item]float64 {
		Null: 1,
	},
	inputs: map[Item]float64 {
		Null: 1,
	},
}

var WoodLog = Item {
	repr: "Wood Log",
	isCompact: true,
}

var WoodPlank = Item {
	repr: "Wood Plank",
	isCompact: false,
}

var WoodPlankRecipe = Recipe {
	outputs: map[Item]float64 {
		WoodPlank: 4,
	},
	inputs: map[Item]float64 {
		WoodLog: 1,
	},
}

var Chest = Item {
	repr: "Chest",
	isCompact: false,
}

var ChestRecipe = Recipe {
	outputs: map[Item]float64 {
		Chest: 1,
	},
	inputs: map[Item]float64 {
		WoodPlank: 8,
	},
}

var RedstoneBlock = Item {
	repr: "Redstone Block",
	isCompact: true,
}

var RedstoneDust = Item {
	repr: "Redstone Dust",
	isCompact: false,
}

var RedstoneDustRecipe = Recipe {
	outputs: map[Item]float64 {
		RedstoneDust: 9,
	},
	inputs: map[Item]float64 {
		RedstoneBlock: 1,
	},
}

var IronBlock = Item {
	repr: "Iron Block",
	isCompact: true,
}

var IronIngot = Item {
	repr: "Iron Ingot",
	isCompact: false,
}

var IronIngotRecipe = Recipe {
	outputs: map[Item]float64 {
		IronIngot: 9,
	},
	inputs: map[Item]float64 {
		IronBlock: 1,
	},
}

var Hopper = Item {
	repr: "Hopper",
	isCompact: false,
}

var HopperRecipe = Recipe {
	outputs: map[Item]float64 {
		Hopper: 1,
	},
	inputs: map[Item]float64 {
		IronIngot: 5,
		Chest: 1,
	},
}
