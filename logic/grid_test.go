package logic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchForSequence_AllAtOnce(t *testing.T) {
	g := &Grid{
		Size: 8,
		Data: []int{
			0, 0, 0, 0, 1, 0, 0, 0,
			1, 34, 0, 0, 0, 0, 0, 0,
			0, 21, 3, 5, 8, 13, 21, 34,
			0, 13, 0, 0, 0, 0, 2, 0,
			13, 8, 5, 3, 2, 0, 3, 1,
			0, 5, 0, 0, 0, 0, 5, 0,
			2, 1, 0, 0, 0, 0, 8, 0,
			0, 0, 0, 0, 0, 0, 13, 0,
		},
	}
	paths := g.SearchForSequence()
	fmt.Println(paths)
	if len(paths) != 4 {
		t.Fatal("exepected 4 paths")
	}
	expectedPaths := [][]int{
		[]int{19, 20, 21, 22, 23},
		[]int{30, 38, 46, 54, 62},
		[]int{36, 35, 34, 33, 32},
		[]int{41, 33, 25, 17, 9},
	}
	for i, expected := range expectedPaths {
		if !reflect.DeepEqual(paths[i], expected) {
			t.Fatalf("expected %v to be %v", paths[0], expected)
		}
	}
}
