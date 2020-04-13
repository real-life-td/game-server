package main

import (
	"fmt"
	"github.com/real-life-td/game-core/world"
)

func main() {
	test := world.NewRoad(23, nil, nil)
	fmt.Print(test.Id())
}