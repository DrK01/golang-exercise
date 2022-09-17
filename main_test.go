package main

import (
	"testing"
)

func TestIsAP(t *testing.T) {

	values := []*Product{
		{ID: 1, Name: "shoes", Price: 8},
		{ID: 2, Name: "shirt", Price: 4},
		{ID: 3, Name: "socks", Price: 5},
		{ID: 4, Name: "sweater", Price: 2},
		{ID: 5, Name: "jacket", Price: 6},
		{ID: 6, Name: "pants", Price: 1},
	}

	result := IsAPS(values)

	if !result {
		t.Error("IsAPS should return true but returned", result)
	}

}

func TestISNotAP(t *testing.T) {
	values := []*Product{
		{ID: 1, Name: "shoes", Price: 8},
		{ID: 2, Name: "shirt", Price: 5},
		{ID: 3, Name: "socks", Price: 10},
		{ID: 4, Name: "sweater", Price: 12},
		{ID: 5, Name: "jacket", Price: 14},
		{ID: 6, Name: "pants", Price: 6},
	}

	result := IsAPS(values)

	if result {
		t.Error("IsAPS should return false but returned", result)
	}
}

func TestLongestAPSSubsequenceLength(t *testing.T) {

	values := []*Product{
		{ID: 1, Name: "shoes", Price: 8},
		{ID: 2, Name: "shirt", Price: 5},
		{ID: 3, Name: "socks", Price: 10},
		{ID: 4, Name: "sweater", Price: 12},
		{ID: 5, Name: "jacket", Price: 14},
		{ID: 6, Name: "pants", Price: 6},
	}

	result := longestAPSSubsequenceLength(values)

	if result != 3 {
		t.Error("longestAPSSubsequenceLength should be 3 but returned", result)
	}

}
