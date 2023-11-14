package main

import (
	"main/geom"
	"main/io"
	"math"
	"os"
	"path/filepath"

	"github.com/celer/csg/csg"
	"github.com/celer/csg/qhull"
	"github.com/golang/geo/r3"
	"github.com/markus-wa/quickhull-go"
)

// Finds the convex hull of a hole using the quickhull algorithm
// if it is the same as the original vertices => convex, else => not convex
func Convex3d(v []geom.Vertex) []geom.Vertex {
	rV := []r3.Vector{}
	for _, vert := range v {
		rV = append(rV, r3.Vector{X: float64(vert.X), Y: float64(vert.Y), Z: float64(vert.Z)})
	}

	ch := quickhull.ConvexHull(rV)

	// turn back into vertex slice
	chVertexSlice := []geom.Vertex{}

	for _, vert := range ch {
		chVertexSlice = append(chVertexSlice, geom.Vertex{X: float32(vert.X), Y: float32(vert.Y), Z: float32(vert.Z)})
	}

	return chVertexSlice

}

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

	//verts := spiral(1000, 1.0, 0.1, 0.1)

	//io.WriteVerticesToOBJ(verts, "spiral.obj")

	//v := Convex3d(m.V)

	//io.WriteVerticesToOBJ(v, "ch.obj")

	m := io.ReadBinarySTL("ring2.stl")

	// for _, vert := range m.V {

	// }

	polygons := []*csg.Polygon{}

	for i := 0; i < len(m.T)/3; i++ {
		idx1 := m.T[i*3]
		idx2 := m.T[i*3+1]
		idx3 := m.T[i*3+2]

		v1 := m.V[idx1]
		v2 := m.V[idx2]
		v3 := m.V[idx3]

		normal := geom.CrossProduct(geom.Subtract(v2, v1), geom.Subtract(v3, v1))

		normal = geom.Normalize(normal)

		vec1 := &csg.Vector{float64(v1.X), float64(v1.Y), float64(v1.Z)}
		vec2 := &csg.Vector{float64(v2.X), float64(v2.Y), float64(v2.Z)}
		vec3 := &csg.Vector{float64(v3.X), float64(v3.Y), float64(v3.Z)}
		normalvec := &csg.Vector{float64(normal.X), float64(normal.Y), float64(normal.Z)}

		csgVert1 := csg.NewVertexFromVectors(vec1, normalvec)
		csgVert2 := csg.NewVertexFromVectors(vec2, normalvec)
		csgVert3 := csg.NewVertexFromVectors(vec3, normalvec)

		polygon := csg.NewPolygonFromVertices([]*csg.Vertex{csgVert1, csgVert2, csgVert3})

		polygons = append(polygons, polygon)

	}

	csgShape := csg.NewCSGFromPolygons(polygons)

	h := &qhull.Hull{}
	h.BuildFromCSG([]*csg.CSG{csgShape})

	//polygonsAfter := csgShape.ToPolygons()

	file := "output.stl"
	os.Mkdir("output", 0755)
	out, err := os.Create(filepath.Join("output", file))
	if err != nil {
		panic(err)
	}

	// convex hull
	c := h.ToCSG()

	//c.MarshalToASCIISTL(out)

	c = c.Subtract(csgShape)

	c.MarshalToASCIISTL(out)

	out.Close()
}
