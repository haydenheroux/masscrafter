package main

var Null = Item {
	repr: "Null",
	isCompact: true,
}

var NullRecipe = Recipe {
	outputs: ItemList {
		Null: 1,
	},
	inputs: ItemList {
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
	outputs: ItemList {
		WoodPlank: 4,
	},
	inputs: ItemList {
		WoodLog: 1,
	},
}

var Chest = Item {
	repr: "Chest",
	isCompact: false,
}

var ChestRecipe = Recipe {
	outputs: ItemList {
		Chest: 1,
	},
	inputs: ItemList {
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
	outputs: ItemList {
		RedstoneDust: 9,
	},
	inputs: ItemList {
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
	outputs: ItemList {
		IronIngot: 9,
	},
	inputs: ItemList {
		IronBlock: 1,
	},
}

var Hopper = Item {
	repr: "Hopper",
	isCompact: false,
}

var HopperRecipe = Recipe {
	outputs: ItemList {
		Hopper: 1,
	},
	inputs: ItemList {
		IronIngot: 5,
		Chest: 1,
	},
}
