package main

import (
	"flag"
	"log"

	"github.com/deeean/go-vector/vector3"
	"github.com/fogleman/gg"
)

func main() {

	const pctX = 0.5
	const pctY = 0.9
	timeStart := flag.String("i", "22:30", "time of start (HH:mm)")
	timeEnd := flag.String("o", "", "time of end. Ignore if not wanted (HH:mmm)")
	minPlayersRequired := flag.Int("p", 4, "Num of minimum Players Required. 0 for not drawing")
	flag.Parse()

	numFlags := flag.NFlag()
	if numFlags == 0 {
		flag.PrintDefaults()
	}
	text := ""
	start := *timeStart
	end := *timeEnd
	minPlayers := *minPlayersRequired

	if len(end) == 0 {
		text = start
	} else {
		text = start + " - " + end
	}

	im, err := gg.LoadImage("assets/potato.png")
	if err != nil {
		log.Fatal(err)
	}

	var width = im.Bounds().Max.X
	var height = im.Bounds().Max.Y

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("assets/COMIC.ttf", 20); err != nil {
		panic(err)
	}

	dc.DrawImage(im, 0, 0)

	drawMegaText(text, float64(width)*pctX, float64(height)*pctY, dc)

	if minPlayers != 0 {
		drawPlayers(minPlayers, 2, 2, dc)
	}
	dc.Clip()

	dc.SavePNG("out/out.png")

}

func drawMegaText(text string, x float64, y float64, context *gg.Context) {

	const shadowOffset = 2
	colorText := vector3.New(132, 247, 41)
	colorShadow := vector3.New(48, 0, 224)

	context.SetRGB255(int(colorShadow.X), int(colorShadow.Y), int(colorShadow.Z))
	context.DrawStringAnchored(text, x+shadowOffset, y+shadowOffset, 0.5, 0.5)

	context.SetRGB255(int(colorText.X), int(colorText.Y), int(colorText.Z))
	context.DrawStringAnchored(text, x, y, 0.5, 0.5)
}

func drawPlayers(numPlayers int, x int, y int, context *gg.Context) {
	im, err := gg.LoadImage("assets/player.png")

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < numPlayers; i++ {
		context.DrawImage(im, x+i*(im.Bounds().Max.X+4), y)
	}

}
