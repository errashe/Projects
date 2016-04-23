// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3.1 and OpenGL 2.1.
package main

import (
	"log"
	"math"
	"math/rand"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(800, 600, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	window.SetKeyCallback(onKey)

	setupScene()
	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

var buffer [][]float32

func recreate(a float32) {
	buffer = nil
	var x float32
	for x = -10; x < 10; x += 0.1 {
		buffer = append(buffer, []float32{x, a * float32(math.Sin(float64(x)))})
	}
}

var push bool = false

func onKey(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	} else if key == glfw.KeyR {
		if action == glfw.Press {
			push = true
		} else if action == glfw.Release {
			push = false
		}
	}
}

func setupScene() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-1, 1, -1, 1, 1.0, 1000.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	recreate(0)
}

func destroyScene() {
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -10)
	// gl.Rotatef(rotationX, 1, 0, 0)
	// gl.Rotatef(rotationY, 0, 1, 0)

	// rotationX += 0.5
	// rotationY += 0.5
	if push {
		recreate(5 * rand.Float32())
	}

	gl.LineWidth(10)
	gl.Begin(gl.LINE_STRIP)
	for _, i := range buffer {
		gl.Vertex2f(i[0], i[1])
	}
	gl.End()
}
