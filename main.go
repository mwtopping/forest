package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	tree := NewTree("./data/heart.csv", 16)

	fmt.Println(tree.get_max_depth(*tree.Root, 1))

	rl.InitWindow(1600, 1600, "raylib [core] example - basic window")
	rl.SetLineWidth(5.0)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		// rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		tree.display()

		rl.EndDrawing()
	}

}
