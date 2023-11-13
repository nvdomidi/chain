package main

import (
	"main/geom"
	"main/io"
	"math"
)

func spiral(numPoints int, radius, heightIncrement, thetaStep float64) []geom.Vertex {
	verts := make([]geom.Vertex, numPoints)

	for i := 0; i < numPoints; i++ {

		theta := thetaStep * float64(i)

		x := radius * math.Cos(theta)
		y := radius * math.Sin(theta)
		z := heightIncrement * theta / (2 * math.Pi)

		verts = append(verts, geom.Vertex{float32(x), float32(y), float32(z)})

	}

	return verts
}

func main() {

	verts := spiral(1000, 1.0, 0.1, 0.1)

	io.WriteVerticesToOBJ(verts, "spiral.obj")

}
