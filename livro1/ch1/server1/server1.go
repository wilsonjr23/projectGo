// Server 1 é um servidor de "eco" mínimo.
package main

import (
	"fmt"
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)

var palette = []color.Color{color.Black,color.RGBA{0xAD, 0xFF, 0x2F, 0xff},color.RGBA{0xFF, 0x33, 0x00, 0xff},color.RGBA{0xFF, 0xFF, 0x00, 0xff},color.RGBA{0x00, 0x33, 0xFF, 0xff}}

const (
	blackIndex = 0 // próxima cor da paleta
	greenIndex = 1 // próxima cor da paleta
	redIndex = 2 // próxima cor da paleta
	yellowIndex = 3 // próxima cor da paleta
	blueIndex = 4 // próxima cor da paleta
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { lissajous(w, r.URL.Query().Get("cycles"))}) // cada requisição chama handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, parm1 string) {
	const (
		res 	= 0.001		// resolução angular
		size	= 100 		// canvas da imagem cobre de [-size..+size]
		nframes = 64 		// número de quadros da animação
		delay 	= 8 		// tempo entre quadros em unidades de 10ms
	)
	var opccor uint8
	cycles, err := strconv.ParseFloat(parm1, 64)
	if err != nil {
		fmt.Println(err)
	}
	freq := rand.Float64() * 3.0 	// frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 	// diferença de fase
	for i := 0; i < nframes; i++ {
		
		opccor++
		if opccor > 4 {
			opccor = 0
		}
		
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),opccor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}