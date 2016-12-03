package main

import "fmt"
import "time"
import ui "github.com/gizak/termui"

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar(":PRESS CTRL-C TO QUIT DEMO")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Text Box"
	p.BorderFg = ui.ColorCyan

	pi := ui.NewPar("Default value")
	pi.Height = 3
	pi.Width = 20
	pi.TextFgColor = ui.ColorWhite
	pi.SetX(0)
	pi.SetY(3)

	go func() {
		var i int = 0
		for {
			pi.Text = fmt.Sprintf("%d str", i)
			i++
			ui.Render(p, pi)
			time.Sleep(time.Second * 1)
		}
	}()

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}
