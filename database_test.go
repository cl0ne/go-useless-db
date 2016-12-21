package main

import (
	"testing"

	"github.com/cl0ne/go-useless-db/database"
)

func TestInsert(t *testing.T) {
	db_test := database.NewDatabase()
	records := []database.Record{
		{Amount: 1, Name: "data"},
		{Amount: -1, Name: "shot"},
		{Name: "rock", Amount: 40000},
	}
	for k, v := range records {
		ok := db_test.Insert(k, v)
		if !ok {
			t.Error("Insertion should be ok.")
		}
	}
}

func TestRemove(t *testing.T) {
	db_test := database.NewDatabase()
	records := []database.Record{
		{Amount: 1, Name: "data"},
		{Amount: -1, Name: "shot"},
		{Name: "rock", Amount: 40000},
	}
	for k, v := range records {
		db_test.Insert(k, v)
	}
	for i := 0; i < db_test.Length(); i++ {
		ok := db_test.Remove(i)
		if !ok {
			t.Error("Index", i, "is valid, data exist")
		}
	}
	ok := db_test.Remove(4)
	if ok {
		if db_test.Length() < 4 {
			t.Error("Index 4 is out of range")
		} else {
			t.Error("Data doesn't exist at 4")
		}
	}
	ok = db_test.Remove(-1)
	if ok {
		t.Error("Index -1 is out of range")
	}
}

func TestUpdate(t *testing.T) {
	db_test := database.NewDatabase()
	records := []database.Record{
		{Amount: 1, Name: "data"},
		{Amount: -1, Name: "shot"},
		{Name: "rock", Amount: 40000},
	}
	for k, v := range records {
		db_test.Insert(k, v)
	}
	r := database.Record{Amount: 2, Name: "data"}
	ok := db_test.Update(0, r)
	if !ok {
		t.Error("Index 0 is valid, data exist")
	}
	ok = db_test.Update(-1, r)
	if ok {
		t.Error("Index -1 is out of range")
	}
	db_test.Remove(0)
	ok = db_test.Update(db_test.Length(), r)
	if ok {
		t.Error("Data doesn't exist at 0")
	}
}

func TestClear(t *testing.T) {
	db_test := database.NewDatabase()
	records := []database.Record{
		{Amount: 1, Name: "data"},
		{Amount: -1, Name: "shot"},
		{Name: "rock", Amount: 40000},
	}
	for k, v := range records {
		db_test.Insert(k, v)
	}
	db_test.Clear()
	if db_test.Length() != 0 {
		t.Error("db_test should be empty")
	}
}

func TestGet(t *testing.T) {
	db_test := database.NewDatabase()
	records := []database.Record{
		{Amount: 1, Name: "data"},
		{Amount: -1, Name: "shot"},
		{Name: "rock", Amount: 40000},
	}
	for k, v := range records {
		db_test.Insert(k, v)
	}
	_, ok := db_test.Get(0)
	if !ok {
		t.Error("Data exists at 0")
	}
	_, ok = db_test.Get(-1)
	if ok {
		t.Error("Index -1 is out of range")
	}
}
