package graph

import (
	"bytes"
)

type Edge struct {
	id        []byte
	weight    int
	endPoint1 Vertex
	endPoint2 Vertex
}

func NewEdge(ep1, ep2 Vertex) (*Edge, error) {
	if ep1 == nil || ep2 == nil {
		return nil, VertexNullError
	}
	var id []byte
	switch bytes.Compare(ep1.Id(), ep2.Id()) {
	case 0:
		return nil, ConnectSelfError
	case -1:
		id = append(ep1.Id(), ep2.Id()...)
	case 1:
		id = append(ep2.Id(), ep1.Id()...)
	}
	e := &Edge{
		id:        id,
		weight:    0,
		endPoint1: ep1,
		endPoint2: ep2,
	}
	ep1.RegisterEdge(e)
	ep2.RegisterEdge(e)
	return e, nil
}

func (e *Edge) Id() []byte {
	return e.id
}

func (e *Edge) SetWeight(weight int) {
	e.weight = weight
}

func (e *Edge) Weight() int {
	return e.weight
}

func (e *Edge) EndPoints() (Vertex, Vertex) {
	return e.endPoint1, e.endPoint2
}

func (e *Edge) Destory() {
	e.endPoint1.RemoveEdge(e)
	e.endPoint2.RemoveEdge(e)
	e = nil
}

func (e *Edge) Equals(e1 *Edge) bool {
	return bytes.Equal(e.Id(), e1.Id())
}
