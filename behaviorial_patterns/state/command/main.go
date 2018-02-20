package main

import "github.com/vg0x00/test/behaviorial_patterns/state"

func main() {
	initState := state.InitState{}
	game := state.Context{}
	game.Next = &initState
	for game.Next.Execute(&game) { // game loop
	}
}
