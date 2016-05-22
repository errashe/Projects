package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"net/http"
	"strconv"
)

var pi float64 = 3.14

func clear(img *image.RGBA) {
	for i := 0; i < (*img).Bounds().Dx(); i++ {
		for j := 0; j < (*img).Bounds().Dy(); j++ {
			img.Set(i, j, color.Black)
		}
	}
}

func drawFrames(A, B, wa, wb, al, bl float64) ([]*image.Paletted, []int) {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	var ret_p []*image.Paletted
	var ret_t []int

	palette := color.Palette{}
	for i := 0; i < 16; i++ {
		palette = append(palette, color.Gray{uint8(i) * 0x11})
	}

	for i := 0; i < 64; i++ {
		clear(img)

		f := al - bl
		for a := 0.0; a < 20*pi; a += 0.01 {
			x := A * math.Cos(wa*a+f)
			y := B * math.Sin(wb*a)

			img.Set(int(x*100)+250, int(y*100)+250, color.White)
		}

		pm := image.NewPaletted(img.Bounds(), palette)
		draw.FloydSteinberg.Draw(pm, img.Bounds(), img, image.ZP)

		ret_p = append(ret_p, pm)
		ret_t = append(ret_t, 4)

		al += pi / 32

	}

	return ret_p, ret_t
}

func main() {
	// f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	// defer f.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println(r.URL.String())
		A, _ := strconv.ParseFloat(r.URL.Query().Get("A"), 64)
		B, _ := strconv.ParseFloat(r.URL.Query().Get("B"), 64)
		wa, _ := strconv.ParseFloat(r.URL.Query().Get("wa"), 64)
		wb, _ := strconv.ParseFloat(r.URL.Query().Get("wb"), 64)
		al, _ := strconv.ParseFloat(r.URL.Query().Get("al"), 64)
		bl, _ := strconv.ParseFloat(r.URL.Query().Get("bl"), 64)
		p, t := drawFrames(A, B, wa, wb, al, bl)

		w.Header().Set("Content-Type", "image/gif")
		gif.EncodeAll(w, &gif.GIF{
			Image: p,
			Delay: t,
		})
	})
	http.ListenAndServe(":8080", nil)
}
