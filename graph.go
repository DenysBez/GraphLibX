package graphlibx

type graph struct {
	name   string
	source edge
}

type edge struct {
	name       string
	linkedEdge []edge
}
