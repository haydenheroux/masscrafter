package main

import (
	"reflect"
	"testing"
)

func TestCraftPlanks(t *testing.T) {
	gotMaterials := Materials("woodPlank", 64)
	wantMaterials := ItemList{
		"woodLog": 16.0,
	}
	if !reflect.DeepEqual(gotMaterials, wantMaterials) {
		t.Errorf("got: %v, want: %v", gotMaterials, wantMaterials)
	}
}

func TestCraftHoppers(t *testing.T) {
	gotMaterials := Materials("hopper", 64.0)
	wantMaterials := ItemList{
		"ironIngot": 5 * 64.0,
		"chest":     64.0,
	}
	if !reflect.DeepEqual(gotMaterials, wantMaterials) {
		t.Errorf("got: %v, want: %v", gotMaterials, wantMaterials)
	}
}

func TestSimplifyCraftIronIngots(t *testing.T) {
	gotMaterials := Materials("ironIngot", 9*64)
	done := false
	for !done {
		gotMaterials, done = Simplify(gotMaterials)
	}
	var wantMaterials = ItemList{
		"ironBlock": 64.0,
	}
	if !reflect.DeepEqual(gotMaterials, wantMaterials) {
		t.Errorf("got: %v, want: %v", gotMaterials, wantMaterials)
	}
}

func TestSimplifyCraftHoppers(t *testing.T) {
	gotMaterials := Materials("hopper", 64)
	done := false
	for !done {
		gotMaterials, done = Simplify(gotMaterials)
	}
	var wantMaterials = ItemList{
		"ironBlock": (5.0 * 64.0) / 9.0,
		"woodLog":   (8.0 * 64.0) / 4.0,
	}
	if !reflect.DeepEqual(gotMaterials, wantMaterials) {
		t.Errorf("got: %v, want: %v", gotMaterials, wantMaterials)
	}
}

func TestSimplifyCraftBookshelfs(t *testing.T) {
	var gotMaterials ItemList = Materials("bookshelf", 1)
	done := false
	for !done {
		gotMaterials, done = Simplify(gotMaterials)
	}
	var wantMaterials = ItemList{
		"leather":   3.0,
		"sugarcane": 9.0,
		"woodLog":   (6.0 / 4.0),
	}
	if !reflect.DeepEqual(gotMaterials, wantMaterials) {
		t.Errorf("got: %v, want: %v", gotMaterials, wantMaterials)
	}
}
