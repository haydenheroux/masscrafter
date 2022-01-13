package main

import (
	"encoding/json"
	"os"
)

func Load() map[string]Item {
	allItems := make(map[string]Item)

	itemsFile, err := os.Open("items.json")
	defer itemsFile.Close()
	if err != nil {
		// TODO: Explain why error occured
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(itemsFile)
	if err = jsonParser.Decode(&allItems); err != nil {
		// TODO: Explain why error occured
		os.Exit(1)
	}

	return allItems
}
