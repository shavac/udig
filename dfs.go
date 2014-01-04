package graph

func dfs(visited map[string]bool, iter <-chan *Edge, f func(Vertex)) {
	for e := range iter {
		if e == nil {
			continue
		}
		v1, v2 := e.EndPoints()

		if !visited[string(v1.Id())] {
			f(v1)
			visited[string(v1.Id())] = true
			dfs(visited, v1.EdgeIter(), f)
		}

		if !visited[string(v2.Id())] {
			f(v2)
			visited[string(v2.Id())] = true
			dfs(visited, v2.EdgeIter(), f)
		}
	}
}
