package graph

type Graph interface {
	EdgeCount() int
	VertexCount() int
	//EdgeIter() chan *Edge
	VertexIter() chan Vertex
	NewVertex() Vertex
	AddVertex(Vertex)
	GetVertexById(id []byte) (Vertex, error)
	RemoveVertex(v Vertex) error
	RemoveEdge(e *Edge) error
	ConnectVertices(v1, v2 Vertex) (*Edge, error)
	DisconnectVertices(v1, v2 Vertex) error
	Dfs(v Vertex, f func(Vertex)) map[string]bool
}

type BaseGraph struct {
	VertexMap map[string]Vertex
	EdgeMap   map[string]*Edge
}

func NewBaseGraph() *BaseGraph {
	bg := new(BaseGraph)
	bg.VertexMap = make(map[string]Vertex)
	bg.EdgeMap = make(map[string]*Edge)
	return bg
}

func (bg *BaseGraph) VertexCount() int {
	return len(bg.VertexMap)
}

func (bg *BaseGraph) EdgeCount() int {
	return len(bg.EdgeMap)
}

func (bg *BaseGraph) AddVertex(v Vertex) {
	bg.VertexMap[string(v.Id())] = v
}

func (bg *BaseGraph) NewVertex() Vertex {
	v := NewBaseVertex()
	bg.AddVertex(v)
	return v
}

func (bg *BaseGraph) GetVertexById(id []byte) (Vertex, error) {
	if v, ok := bg.VertexMap[string(id)]; ok {
		return v, nil
	}
	return nil, VertexNotExistError
}

func (bg *BaseGraph) RemoveVertexById(id []byte) error {
	v, err := bg.GetVertexById(id)
	if err != nil {
		return VertexNotExistError
	}
	for e := range v.EdgeIter() {
		e.Destory()
		delete(bg.EdgeMap, string(e.Id()))
	}

	delete(bg.VertexMap, string(id))
	return nil
}

func (bg *BaseGraph) RemoveVertex(v Vertex) error {
	return bg.RemoveVertexById(v.Id())
}

func (bg *BaseGraph) ConnectVertices(v1, v2 Vertex) (*Edge, error) {
	e, _ := NewEdge(v1, v2)
	if e_old, err := bg.GetEdgeById(e.Id()); err == EdgeNotExistError {
		bg.EdgeMap[string(e.Id())] = e
		v1.RegisterEdge(e)
		v2.RegisterEdge(e)
		return e, nil
	} else if err == nil {
		return e_old, AlreadyConnectedError
	}
	return nil, nil
}

func (bg *BaseGraph) RemoveEdge(e *Edge) error {
	return bg.RemoveEdgeById(e.Id())
}

func (bg *BaseGraph) RemoveEdgeById(id []byte) error {
	e, err := bg.GetEdgeById(id)
	if err != nil {
		return err
	}
	delete(bg.EdgeMap, string(id))
	e.Destory()
	return nil
}

func (bg *BaseGraph) DisconnectVertices(v1, v2 Vertex) error {
	e, _ := NewEdge(v1, v2)
	if e, err := bg.GetEdgeById(e.Id()); err != nil {
		return err
	} else {
		bg.RemoveEdge(e)
	}
	return nil
}

func (bg *BaseGraph) GetEdgeById(id []byte) (e *Edge, err error) {
	if e, ok := bg.EdgeMap[string(id)]; ok {
		return e, nil
	} else {
		return nil, EdgeNotExistError
	}
}

func (bg *BaseGraph) VertexIter() chan Vertex {
	cv := make(chan Vertex)
	go func() {
		for _, v := range bg.VertexMap {
			cv <- v
		}
	}()
	return cv
}

func (bg *BaseGraph) Dfs(v Vertex, f func(Vertex)) map[string]bool {
	visited := make(map[string]bool)
	visited[string(v.Id())] = true
	dfs(visited, v.EdgeIter(), f)
	return visited
}

func (bg *BaseGraph) ConnectedGraph(v Vertex) Graph {
	g := NewBaseGraph()
	f := func(cv Vertex) {
		g.AddVertex(cv)
	}
	bg.Dfs(v, f)
	return g
}
