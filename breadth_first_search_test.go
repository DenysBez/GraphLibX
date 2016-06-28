package graphlibx

import (
	"fmt"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {

	mainEdge := edge{"test_edge", make([]edge, 1, 5)}
	mainEdge.linkedEdge = LinkEdge(mainEdge, edge{"test_edge_two", make([]edge, 1, 5)})

	testGraph := graph{"test_graph", mainEdge}

	fmt.Print(testGraph)
}

func LinkEdge(mainEdge edge, edgeToLink edge) []edge {
	return append(mainEdge.linkedEdge, edgeToLink)
}
