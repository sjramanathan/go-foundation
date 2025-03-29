package main

import (
	"fmt"
	"slices"
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

	k := Jade
	fmt.Printf("k: %v \n", k)
	fmt.Printf("k: %#v \n", k)
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
	Keys []Key
}

type Key byte

const (
	Jade Key = iota + 1 // First value starts with one and then each value after is an increment. This structure is used for enums.
	Copper
	Crystal
	invalidKey
)

// The reason why enabling this immediately causes "fmt" to print the string is that here I'm inheriting an interface. This notifies the "fmt" package that this is available for use.
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("This is an unknown key")
	}

	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}

	return nil
}

func containsKey(keys []Key, k Key) bool {
	return slices.Contains(keys, k)
}
