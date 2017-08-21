# fractal

Fractal image generator written in Go.

Requires draw2d, see https://godoc.org/github.com/llgcode/draw2d. 
Radius.go creates a curve made of perpendicular line segments, randomly oriented (essentialy a random walk),
then grabs that as a seed and creates a fractal curve. 
