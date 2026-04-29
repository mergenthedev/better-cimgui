package main

import (
	"fmt"

	"github.com/go-gl/glfw/v3.4/glfw"
	bettercimgui "github.com/mergenthedev/better-cimgui"
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Println("Cannot initialize GLFW")
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(800, 600, "Demo", nil, nil)
	if err != nil {
		fmt.Println("Cannot initialize Window")
	}

	window.MakeContextCurrent()

	winC := bettercimgui.GetGLFWWindowCPtr(window)
	bettercimgui.Init(winC)
	defer bettercimgui.Terminate()

	for !window.ShouldClose() {
		bettercimgui.Update()

		bettercimgui.Demo()

		bettercimgui.Render()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
