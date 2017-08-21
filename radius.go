package main

import (
	"fmt"	
	"time"
	"math/rand"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
)

type Vector struct {
	dir int
	a uint8
	b uint8
	c uint8
}

//code to draw ortho shape, then copy and append to itseld in random direction, again and again
func main() {
	rand.Seed(time.Now().UnixNano())
	height, width := 20000, 20000
	dest	:= image.NewRGBA(image.Rect(0, 0, height, width))
	canvas	:= draw2dimg.NewGraphicContext(dest)
	canvas.SetFillColor(color.RGBA{0x00, 0xff, 0xff, 0xff})
	canvas.SetStrokeColor(color.RGBA{0x00, 0xff, 0xff, 0xff})
	canvas.SetLineWidth(2)
	canvas.SetFillColor(image.White)
	// fill the background
	canvas.Clear()

	n := 4
	n_steps := 20
	fmt.Println("N:", n)
	var allvect []Vector = make([]Vector, 0)
	allvect, lastx, lasty := initial(allvect, n, width, height, canvas)
	for k :=0;  k < n_steps; k++ {
		rotate := int(rand.Int31n(1))
		//fset the following to > -1 to be always true; for one style
		if rotate > -1 {
			rotate = 1
			fmt.Println("neg rot?")
		}
		//so rotate becomes 1 or -1, and we can add this to rotate the direction of each vector
		fmt.Println("step: ", k)
		a := uint8(rand.Int31n(255))
		b := uint8(rand.Int31n(255))
		c := uint8(rand.Int31n(255))
		//pick the color now as draw just draws the given colors, here so one coor per step
		var newvect []Vector = make([]Vector, 0)
		for _, currvect := range allvect {
			var addvect Vector
			addvect.a = a
			addvect.b = b
			addvect.c = c
			addvect.dir = compare(currvect.dir, rotate)
			newvect = append(newvect, addvect)
		}
		lastx, lasty = draw(canvas, lastx, lasty, allvect)
		for _, vect := range newvect {
			allvect = append(allvect, vect)
		}
	}
	SAVE := "doubleposrad.png"
	draw2dimg.SaveToPngFile(SAVE, dest)
	fmt.Println("DONE")
}

func initial(allvect []Vector, n int, w int, h int, canvas *draw2dimg.GraphicContext) (totalvect []Vector, lastx,lasty float64) {
//	red == MakeColor(255,0,0)
//	green == MakeColor(0,255,0)	
//	yellow == MakeColor(255,255,0)
//	blue == MakeColor(0,0,255)
	var a,b,c uint8
	for i:= 0; i < n; i++ {
		predir := rand.Int31n(4)
		dir := int(predir)
		dir, a,b,c = assign(allvect, dir)
		var new Vector
		new.dir = dir
		new.a = a
		new.b = b
		new.c = c
		allvect = append(allvect, new)
	}
	x,y := float64(w/2),float64(h/2)
	lastx, lasty = draw(canvas, x, y, allvect)
	return allvect, lastx, lasty
}



func assign(allvect []Vector, dir int) (direc int, a,b,c uint8){
		if dir == 0 {
			a = 0
			b = 255
			c = 0		
			if len(allvect) != 0 {
				if allvect[len(allvect) -1].dir == -1*dir {
					dir = dir+1
				}
			} 	
		}
		if dir == 1 {
			a = 255
			b = 0
			c = 0
			if len(allvect) != 0 {
				if allvect[len(allvect) -1].dir == -1*dir {
					dir = dir+1
				}
			} 	
		}
		if dir == 2 {
			a = 0
			b = 0
			c = 255
			if len(allvect) != 0 {
				if allvect[len(allvect) -1].dir == -1*dir {
					dir = dir+1
				}
			} 	
		}
		if dir == 3 {
			a = 0
			b = 100
			c = 100
			if len(allvect) != 0 {
				if allvect[len(allvect) -1].dir == -1*dir {
					dir = 0
				}
			} 	
		}
		direc = dir
		return direc, a, b, c
}

//looks at vect detaisl and draws it accordingly
func draw(canvas *draw2dimg.GraphicContext, x float64, y float64, allvect []Vector) (x2, y2 float64){
	canvas.MoveTo(x,y)
	var a,b,c uint8
	for _, currvect := range allvect {
		canvas.SetStrokeColor(MakeColor(a, b, c))
		a,b,c = currvect.a, currvect.b, currvect.c
		if currvect.dir == 2 {
			y = y + 5
		}
		if currvect.dir == 0 {
			y = y - 5
		}
		if currvect.dir == 1 {
			x = x + 5
		}
		if currvect.dir == 3 {
			x = x - 5
		}
		canvas.LineTo(x,y)
	}
	canvas.Stroke()
	canvas.FillStroke()
	return x,y
}

func compare(currdir, rotate int) int{
	if currdir + rotate > 3{
		dir := 0
		if currdir + rotate > 4 {
			dir := 1
			return dir
		}
		return dir
	}
	if currdir + rotate < 0 {
		dir := 3
		if currdir + rotate < -1 {
			dir := 2
			return dir
		}
		return dir
	}
	dir := currdir + 2*rotate
	return dir
}


func MakeColor(r, g, b uint8) color.Color {
	return &color.RGBA{r, g, b, 255}
}




