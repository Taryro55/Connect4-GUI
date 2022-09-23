package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	w, h := int32(800), int32(450)
	rl.InitWindow(w , h, "Connect 4!")
	rl.SetTargetFPS(60)

	x, y := int32(0),int32(0)

	for !rl.WindowShouldClose() {

		if (rl.IsKeyPressed(257) && (rl.IsKeyDown(342) || rl.IsKeyDown(346))) {
			display := rl.GetCurrentMonitor()

			if rl.IsWindowFullscreen() {
				rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
			} else {
				rl.SetWindowSize(int(w), int(h))
			}
			rl.ToggleFullscreen()
		}

		x, y = rl.GetMouseX(), rl.GetMouseY()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText("Alt + Enter to get in fullscreen mode!", w*24/100, h/2, 20, rl.LightGray)
		rl.DrawRectangleLines(w*2/100, h*2/100, w-w*4/100, h-h*4/100, rl.LightGray)
		rl.DrawCircle(x, y, 30, rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}