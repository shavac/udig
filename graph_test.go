package graph

import (
	"testing"
)

func TestGraphCreation(t *testing.T) {
	g := NewBaseGraph()

	if g.EdgeCount() != 0 {
		t.Errorf("Started with edges")
	}

	if g.VertexCount() != 0 {
		t.Errorf("Started with vertices")
	}
}

func TestNewVertex(t *testing.T) {
	g := NewBaseGraph()

	g.NewVertex()

	if g.VertexCount() != 1 {
		t.Errorf("Didn't add vertex")
	}
}

func TestRemoveVertex(t *testing.T) {
	g := NewBaseGraph()
	v := g.NewVertex()
	g.RemoveVertex(v)

	if g.VertexCount() != 0 {
		t.Errorf("Didn't remove vertex from count")
	}
}

func TestRemoveEdgeViaRemoveVertex(t *testing.T) {
	g := NewBaseGraph()

	v1 := g.NewVertex()
	v2 := g.NewVertex()
	g.ConnectVertices(v1, v2)
	g.RemoveVertex(v1)

	if g.EdgeCount() != 0 {
		t.Errorf("Didn't remove edge from count")
	}

	if v2.EdgeCount() != 0 {
		t.Errorf("Didn't remove edge from second vertex after edge removal")
	}
}

func TestDisconnectVertices(t *testing.T) {
	g := NewBaseGraph()

	v1 := g.NewVertex()
	v2 := g.NewVertex()
	g.ConnectVertices(v1, v2)

	g.DisconnectVertices(v1, v2)

	if g.EdgeCount() != 0 {
		t.Errorf("Didn't remove edge from count")
	}

	if v1.EdgeCount() != 0 || v2.EdgeCount() != 0 {
		t.Errorf("Didn't correctly remove edge from vertices")
	}
}

func TestLookupVertex(t *testing.T) {
	g := NewBaseGraph()
	v := g.NewVertex()

	if v, err := g.GetVertexById([]byte{1, 2, 3}); v != nil || err == nil {
		t.Errorf("Found a nonsense vertex")
	}

	if v, err := g.GetVertexById(v.Id()); v == nil || err != nil {
		t.Errorf("Didn't find a proper vertex")
	}
}

func TestJoinVerticesWithEdge(t *testing.T) {
	g := NewBaseGraph()

	v1 := g.NewVertex()
	v2 := g.NewVertex()
	edge, _ := g.ConnectVertices(v1, v2)

	if g.EdgeCount() != 1 {
		t.Errorf("Didn't correctly add edge to edge count")
	}

	v3, v4 := edge.EndPoints()
	if v3 != v1 || v4 != v2 {
		t.Errorf("Didn't correctly set endpoints")
	}

	if v1.EdgeCount() != 1 || v2.EdgeCount() != 1 {
		t.Errorf("Didn't correctly add edge reference to vertex")
	}
}

func TestEdgeExistence(t *testing.T) {
	g := NewBaseGraph()

	v1 := g.NewVertex()
	v2 := g.NewVertex()
	edge, _ := g.ConnectVertices(v1, v2)

	if e, err := g.GetEdgeById(edge.Id()); !(e.Equals(edge) && err == nil) {
		println(string(e.Id()), "    ", string(edge.Id()))
		t.Errorf("Failed to correctly determine if an edge exists in the graph")
	}
}

func TestSingleEdgeBetweenVertices(t *testing.T) {
	g := NewBaseGraph()

	v1 := g.NewVertex()
	v2 := g.NewVertex()
	edge, _ := g.ConnectVertices(v1, v2)
	edge.SetWeight(2)

	//Sneaking a second test that makes sure params are unordered
	edge1, _ := g.ConnectVertices(v2, v1)

	if edge1.Weight() != 2 {
		t.Errorf("Didn't correctly return the same edge: %d", edge1.Weight())
	}

}

func TestDfsTraversal(t *testing.T) {
	g := NewBaseGraph()

	vertices := new([100]Vertex)

	for i := 0; i < 100; i++ {
		vertices[i] = g.NewVertex()
	}

	for i := 0; i < 99; i++ {
		v1 := vertices[i]
		v2 := vertices[i+1]
		g.ConnectVertices(v1, v2)
	}

	var i int32 = 0
	visiter := func(v Vertex) {
		i++
	}

	g.Dfs(vertices[0], visiter)

	if i != 99 {
		t.Errorf("Did not visit all vertices: %i", i)
	}
}

func TestFindConnectedComponentsDifferentComponents(t *testing.T) {
	g := NewBaseGraph()

	vertices := new([1000]Vertex)

	for i := 0; i < 1000; i++ {
		vertices[i] = g.NewVertex()
	}

	for i := 0; i < 499; i++ {
		v1 := vertices[i]
		v2 := vertices[i+1]
		g.ConnectVertices(v1, v2)
	}

	for i := 500; i < 999; i++ {
		v1 := vertices[i]
		v2 := vertices[i+1]
		g.ConnectVertices(v1, v2)
	}

	g1 := g.ConnectedGraph(vertices[80])
	g2 := g.ConnectedGraph(vertices[999])
	if g1.VertexCount() != 499 || g2.VertexCount() != 499 {
		t.Errorf("g1 and g2 should have 499 and 499 vertex but got %d and %d", g1.VertexCount(), g2.VertexCount())
	}
}
