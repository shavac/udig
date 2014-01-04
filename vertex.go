package graph

import (
	"bytes"
)

type Vertex interface {
	RegisterEdge(*Edge)
	RemoveEdge(*Edge)
	Id() []byte
	Equals(Vertex) bool
	EdgeIter() chan *Edge
	EdgeCount() int
	Destory()
}

type BaseVertex struct {
	edgeMap map[string]*Edge
	id      []byte
}

func NewBaseVertex() *BaseVertex {
	id, _ := NewID()
	return &BaseVertex{
		id:      id,
		edgeMap: make(map[string]*Edge),
	}
}

func (v *BaseVertex) Id() []byte {
	return v.id
}

func (v *BaseVertex) RegisterEdge(e *Edge) {
	v.edgeMap[string(e.Id())] = e
}

func (v *BaseVertex) RemoveEdge(e *Edge) {
	delete(v.edgeMap, string(e.Id()))
}

func (v *BaseVertex) EdgeIter() chan *Edge {
	ec := make(chan *Edge)
	go func() {
		for _, v := range v.edgeMap {
			ec <- v
		}
		close(ec)
	}()
	return ec
}

func (v *BaseVertex) Equals(v2 Vertex) bool {
	return bytes.Equal(v.Id(), v2.Id())
}

func (v *BaseVertex) EdgeCount() int {
	return len(v.edgeMap)
}

func (v *BaseVertex) Destory() {
	for e := range v.EdgeIter() {
		e.endPoint1.RemoveEdge(e)
		e.endPoint2.RemoveEdge(e)
	}
	v = nil
}

