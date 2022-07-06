package main

type Recipe struct {
	Outputs ItemList `json:"outputs"`
	Inputs  ItemList `json:"inputs"`
}

type Item struct {
	Repr         string `json:"repr"`
	IsSimplified bool   `json:"isSimplified"`
	CraftRecipe  Recipe `json:"recipe"`
}

type ItemList map[string]float64
