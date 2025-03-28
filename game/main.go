package main

import (
	"fmt"
)

type Item struct {
	X int
	Y int
}

func main() {

	var i1 Item
	fmt.Printf("i1: %#v \n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v \n", i2)

	i3 := Item{
		Y: 10,
	}
	fmt.Printf("i3: %#v \n", i3)
}

func NewItem1(x, y int) Item {
	return Item{X: x, Y: y}
}

func NewItem2(x, y int) (Item, error) {
	var result Item

	if x < 0 || y < 0 {
		return result, fmt.Errorf("x and y must be greater than 0")
	}

	result = Item{X: x, Y: y}
	return result, nil
}

func NewItem3(x, y int) *Item {
	i := Item{X: x, Y: y}
	return &i
}

func NewItem4(x, y int) (*Item, error) {
	var i Item
	if x < 0 || y < 0 {
		return nil, fmt.Errorf("x and y must be greater than 0")
	}

	i = Item{X: x, Y: y}
	return &i, nil
}
 
type Player struct {
	Name string
	Item // Here its an embedded. This allows for the fields in Item to be lifted up in to the root level.
} 

func moveAll(ms []mover, x, y int) {
	for _, m := range
}