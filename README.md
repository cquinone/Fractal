# fractal

Fractal image generator written in Go.

Requires draw2d, see https://godoc.org/github.com/llgcode/draw2d.

Radius.go creates an inital N-length curve made of perpendicular line segments, randomly oriented (essentially a random walk),
then uses that as a seed and creates a fractal curve. 
