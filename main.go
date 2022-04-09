package main

import (
	"fmt"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"io/ioutil"
)

func GenerateTerminalMap(i ...int) (ans map[int]struct{}) {
	ans = make(map[int]struct{})
	for _, value := range i {
		ans[value] = struct{}{}
	}
	return ans
}
func GenerateNodesSlice(graph *cgraph.Graph, name string, n int, terminal map[int]struct{}) (ans []*cgraph.Node) {
	graph.SetRankDir("LR")
	start, _ := graph.CreateNode("")
	start.SetShape("none")

	for i := 1; i <= n; i++ {
		buf, _ := graph.CreateNode(fmt.Sprintf("%s%d", name, i))
		if _, ok := terminal[i]; ok {
			buf.SetShape("doublecircle")
		} else {
			buf.SetShape("circle")
		}
		ans = append(ans, buf)
	}

	starta, _ := graph.CreateEdge("", start, ans[0])
	starta.SetLabel("start")

	return ans
}
func GenerateNodeMatrix(graph *cgraph.Graph, name1, name2 string, n, m int, terminal1 map[int]struct{},
	terminal2 map[int]struct{}) (ans [][]*cgraph.Node) {
	graph.SetRankDir("LR")
	start, _ := graph.CreateNode("")
	start.SetShape("none")

	for i := 1; i <= n; i++ {
		var bufSlice []*cgraph.Node
		for j := 1; j <= m; j++ {
			node, _ := graph.CreateNode(fmt.Sprintf("%s%d%s%d", name1, i, name2, j))
			node.SetShape("circle")
			if _, ok := terminal1[i]; ok {
				if _, ok := terminal2[j]; ok {
					node.SetShape("doublecircle")
				}
			}

			bufSlice = append(bufSlice, node)
		}
		ans = append(ans, bufSlice)
	}

	start_, _ := graph.CreateEdge("", start, ans[0][0])
	start_.SetLabel("start")
	return ans
}
func GenerateConnect(graph *cgraph.Graph, nodes []*cgraph.Node, a, b int, name string) {
	ab, _ := graph.CreateEdge(name, nodes[a-1], nodes[b-1])
	ab.SetLabel(name)
}
func GenerateDoubleConnect(graph *cgraph.Graph, matr [][]*cgraph.Node, a1, b1, a2, b2 int, name string) {
	ab, _ := graph.CreateEdge(name, matr[a1-1][b1-1], matr[a2-1][b2-1])
	ab.SetLabel(name)
}

func GenerateGraph_1_1() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "q", 2, GenerateTerminalMap(2))
	GenerateConnect(graph, nodes, 1, 2, "c")
	GenerateConnect(graph, nodes, 1, 1, "a, b")
	GenerateConnect(graph, nodes, 2, 2, "a, b")
	g.RenderFilename(graph, graphviz.SVG, "task 1/graph_1_1.svg")
}
func GenerateGraph_1_2() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "q", 13, GenerateTerminalMap(3, 5, 6, 11, 13))
	GenerateConnect(graph, nodes, 1, 2, "b")
	GenerateConnect(graph, nodes, 1, 7, "a")
	GenerateConnect(graph, nodes, 2, 3, "b")
	GenerateConnect(graph, nodes, 3, 3, "b")
	GenerateConnect(graph, nodes, 3, 5, "a")
	GenerateConnect(graph, nodes, 5, 5, "b")
	GenerateConnect(graph, nodes, 5, 6, "a")
	GenerateConnect(graph, nodes, 6, 6, "b")
	GenerateConnect(graph, nodes, 2, 4, "a")
	GenerateConnect(graph, nodes, 4, 5, "b")
	GenerateConnect(graph, nodes, 4, 12, "a")
	GenerateConnect(graph, nodes, 12, 11, "b")
	GenerateConnect(graph, nodes, 11, 11, "b")
	GenerateConnect(graph, nodes, 7, 8, "a")
	GenerateConnect(graph, nodes, 8, 9, "b")
	GenerateConnect(graph, nodes, 9, 11, "b")
	GenerateConnect(graph, nodes, 7, 10, "b")
	GenerateConnect(graph, nodes, 10, 9, "a")
	GenerateConnect(graph, nodes, 10, 13, "b")
	GenerateConnect(graph, nodes, 13, 13, "b")
	GenerateConnect(graph, nodes, 13, 11, "a")

	g.RenderFilename(graph, graphviz.SVG, "task 1/graph_1_2.svg")
}
func GenerateGraph_1_3() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "q", 3, GenerateTerminalMap(2, 3))
	GenerateConnect(graph, nodes, 1, 2, "a")
	GenerateConnect(graph, nodes, 2, 1, "b")

	GenerateConnect(graph, nodes, 1, 3, "b")
	GenerateConnect(graph, nodes, 3, 1, "a")

	g.RenderFilename(graph, graphviz.SVG, "task 1/graph_1_3.svg")
}
func GenerateGraph_1_4() {
	g := graphviz.New()
	graph, _ := g.Graph()

	_ = GenerateNodesSlice(graph, "q", 1, GenerateTerminalMap(1))
	g.RenderFilename(graph, graphviz.SVG, "task 1/graph_1_4.svg")
}

func GenerateGraph_2_1_1() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "q", 3, GenerateTerminalMap(3))
	GenerateConnect(graph, nodes, 1, 2, "a")
	GenerateConnect(graph, nodes, 2, 3, "a")
	GenerateConnect(graph, nodes, 1, 1, "b")
	GenerateConnect(graph, nodes, 2, 2, "b")
	GenerateConnect(graph, nodes, 3, 3, "a, b")

	g.RenderFilename(graph, graphviz.SVG, "task 2/1/graph_2_1_1.svg")
}
func GenerateGraph_2_1_2() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "z", 3, GenerateTerminalMap(3))
	GenerateConnect(graph, nodes, 1, 2, "b")
	GenerateConnect(graph, nodes, 2, 3, "b")
	GenerateConnect(graph, nodes, 1, 1, "a")
	GenerateConnect(graph, nodes, 2, 2, "a")
	GenerateConnect(graph, nodes, 3, 3, "a, b")

	g.RenderFilename(graph, graphviz.SVG, "task 2/1/graph_2_1_2.svg")
}
func GenerateGraph_2_1_3() {
	g := graphviz.New()
	graph, _ := g.Graph()

	matr := GenerateNodeMatrix(graph, "q", "z", 3, 3, GenerateTerminalMap(3), GenerateTerminalMap(3))

	GenerateDoubleConnect(graph, matr, 1, 1, 2, 1, "a")
	GenerateDoubleConnect(graph, matr, 2, 1, 3, 1, "a")
	GenerateDoubleConnect(graph, matr, 3, 1, 3, 1, "a")

	GenerateDoubleConnect(graph, matr, 1, 2, 2, 2, "a")
	GenerateDoubleConnect(graph, matr, 2, 2, 3, 2, "a")
	GenerateDoubleConnect(graph, matr, 3, 2, 3, 2, "a")

	GenerateDoubleConnect(graph, matr, 1, 3, 2, 3, "a")
	GenerateDoubleConnect(graph, matr, 2, 3, 3, 3, "a")
	GenerateDoubleConnect(graph, matr, 3, 3, 3, 3, "a, b")

	GenerateDoubleConnect(graph, matr, 1, 1, 1, 2, "b")
	GenerateDoubleConnect(graph, matr, 1, 2, 1, 3, "b")
	GenerateDoubleConnect(graph, matr, 1, 3, 1, 3, "b")

	GenerateDoubleConnect(graph, matr, 2, 1, 2, 2, "b")
	GenerateDoubleConnect(graph, matr, 2, 2, 2, 3, "b")
	GenerateDoubleConnect(graph, matr, 2, 3, 2, 3, "b")

	GenerateDoubleConnect(graph, matr, 3, 1, 3, 2, "b")
	GenerateDoubleConnect(graph, matr, 3, 2, 3, 3, "b")

	g.RenderFilename(graph, graphviz.SVG, "task 2/1/graph_2_1_3.svg")
}

func GenerateGraph_2_2_1() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "q", 4, GenerateTerminalMap(4))
	GenerateConnect(graph, nodes, 1, 2, "a")
	GenerateConnect(graph, nodes, 2, 3, "a")
	GenerateConnect(graph, nodes, 3, 4, "a")
	GenerateConnect(graph, nodes, 1, 1, "b")
	GenerateConnect(graph, nodes, 2, 2, "b")
	GenerateConnect(graph, nodes, 3, 3, "b")
	GenerateConnect(graph, nodes, 4, 4, "a, b")

	g.RenderFilename(graph, graphviz.SVG, "task 2/2/graph_2_2_1.svg")
}
func GenerateGraph_2_2_2() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "z", 2, GenerateTerminalMap(2))
	GenerateConnect(graph, nodes, 1, 2, "b")
	GenerateConnect(graph, nodes, 2, 1, "b")
	GenerateConnect(graph, nodes, 1, 1, "a")
	GenerateConnect(graph, nodes, 2, 2, "a")

	g.RenderFilename(graph, graphviz.SVG, "task 2/2/graph_2_2_2.svg")
}
func GenerateGraph_2_2_3() {
	g := graphviz.New()
	graph, _ := g.Graph()

	matr := GenerateNodeMatrix(graph, "q", "z", 4, 2, GenerateTerminalMap(4), GenerateTerminalMap(2))
	GenerateDoubleConnect(graph, matr, 1, 1, 1, 2, "b")
	GenerateDoubleConnect(graph, matr, 1, 2, 1, 1, "b")

	GenerateDoubleConnect(graph, matr, 1, 1, 2, 1, "a")
	GenerateDoubleConnect(graph, matr, 2, 1, 2, 2, "b")
	GenerateDoubleConnect(graph, matr, 2, 2, 2, 1, "b")
	GenerateDoubleConnect(graph, matr, 1, 2, 2, 2, "a")

	GenerateDoubleConnect(graph, matr, 2, 1, 3, 1, "a")
	GenerateDoubleConnect(graph, matr, 3, 1, 3, 2, "b")
	GenerateDoubleConnect(graph, matr, 3, 2, 3, 1, "b")
	GenerateDoubleConnect(graph, matr, 2, 2, 3, 2, "a")

	GenerateDoubleConnect(graph, matr, 3, 1, 4, 1, "a")
	GenerateDoubleConnect(graph, matr, 4, 1, 4, 2, "b")
	GenerateDoubleConnect(graph, matr, 4, 2, 4, 1, "b")
	GenerateDoubleConnect(graph, matr, 3, 2, 4, 2, "a")

	GenerateDoubleConnect(graph, matr, 4, 1, 4, 1, "a")
	GenerateDoubleConnect(graph, matr, 4, 2, 4, 2, "a")

	matr[0][0].SetGroup("q")
	matr[1][0].SetGroup("q")
	matr[2][0].SetGroup("q")
	matr[3][0].SetGroup("q")

	matr[0][1].SetGroup("z")
	matr[1][1].SetGroup("z")
	matr[2][1].SetGroup("z")
	matr[3][1].SetGroup("z")

	g.RenderFilename(graph, graphviz.SVG, "task 2/2/graph_2_2_3.svg")
}

func GenerateGraph_2_3_1() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "q", 2, GenerateTerminalMap(1))
	GenerateConnect(graph, nodes, 1, 2, "a")
	GenerateConnect(graph, nodes, 2, 1, "a")
	GenerateConnect(graph, nodes, 1, 1, "b")
	GenerateConnect(graph, nodes, 2, 2, "b")

	g.RenderFilename(graph, graphviz.SVG, "task 2/3/graph_2_3_1.svg")
}
func GenerateGraph_2_3_2() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "z", 3, GenerateTerminalMap(1))
	GenerateConnect(graph, nodes, 1, 2, "b")
	GenerateConnect(graph, nodes, 2, 3, "b")
	GenerateConnect(graph, nodes, 3, 1, "b")

	GenerateConnect(graph, nodes, 1, 1, "a")
	GenerateConnect(graph, nodes, 2, 2, "a")
	GenerateConnect(graph, nodes, 3, 3, "a")

	g.RenderFilename(graph, graphviz.SVG, "task 2/3/graph_2_3_2.svg")
}
func GenerateGraph_2_3_3() {
	g := graphviz.New()
	graph, _ := g.Graph()

	matr := GenerateNodeMatrix(graph, "q", "z", 2, 3, GenerateTerminalMap(1), GenerateTerminalMap(1))
	GenerateDoubleConnect(graph, matr, 1, 1, 2, 1, "a")
	GenerateDoubleConnect(graph, matr, 2, 1, 1, 1, "a")

	GenerateDoubleConnect(graph, matr, 1, 2, 2, 2, "a")
	GenerateDoubleConnect(graph, matr, 2, 2, 1, 2, "a")

	GenerateDoubleConnect(graph, matr, 1, 3, 2, 3, "a")
	GenerateDoubleConnect(graph, matr, 2, 3, 1, 3, "a")

	GenerateDoubleConnect(graph, matr, 1, 1, 1, 2, "b")
	GenerateDoubleConnect(graph, matr, 1, 2, 1, 3, "b")
	GenerateDoubleConnect(graph, matr, 1, 3, 1, 1, "b")

	GenerateDoubleConnect(graph, matr, 2, 1, 2, 2, "b")
	GenerateDoubleConnect(graph, matr, 2, 2, 2, 3, "b")
	GenerateDoubleConnect(graph, matr, 2, 3, 2, 1, "b")

	g.RenderFilename(graph, graphviz.SVG, "task 2/3/graph_2_3_3.svg")
}
func GenerateGraph_2_4() {
	g := graphviz.New()
	graph, _ := g.Graph()

	matr := GenerateNodeMatrix(graph, "q", "z", 2, 3, GenerateTerminalMap(1, 2), GenerateTerminalMap(1, 2, 3))
	GenerateDoubleConnect(graph, matr, 1, 1, 2, 1, "a")
	GenerateDoubleConnect(graph, matr, 2, 1, 1, 1, "a")

	GenerateDoubleConnect(graph, matr, 1, 2, 2, 2, "a")
	GenerateDoubleConnect(graph, matr, 2, 2, 1, 2, "a")

	GenerateDoubleConnect(graph, matr, 1, 3, 2, 3, "a")
	GenerateDoubleConnect(graph, matr, 2, 3, 1, 3, "a")

	GenerateDoubleConnect(graph, matr, 1, 1, 1, 2, "b")
	GenerateDoubleConnect(graph, matr, 1, 2, 1, 3, "b")
	GenerateDoubleConnect(graph, matr, 1, 3, 1, 1, "b")

	GenerateDoubleConnect(graph, matr, 2, 1, 2, 2, "b")
	GenerateDoubleConnect(graph, matr, 2, 2, 2, 3, "b")
	GenerateDoubleConnect(graph, matr, 2, 3, 2, 1, "b")

	matr[0][0].SetShape("circle")

	g.RenderFilename(graph, graphviz.SVG, "task 2/4/graph_2_4.svg")
}

func GenerateGraph_3_1() {
	g := graphviz.New()

	b, _ := ioutil.ReadFile("./task 3/dots/graph_3_1_1.dot")
	graph, _ := graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_1_1.svg")

	b, _ = ioutil.ReadFile("./task 3/dots/graph_3_1_2.dot")
	graph, _ = graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_1_2.svg")
}
func GenerateGraph_3_2() {
	g := graphviz.New()

	b, _ := ioutil.ReadFile("./task 3/dots/graph_3_2_1.dot")
	graph, _ := graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_2_1.svg")

	b, _ = ioutil.ReadFile("./task 3/dots/graph_3_2_2.dot")
	graph, _ = graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_2_2.svg")

	b, _ = ioutil.ReadFile("./task 3/dots/graph_3_2_3.dot")
	graph, _ = graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_2_3.svg")
}
func GenerateGraph_3_3() {
	g := graphviz.New()

	b, _ := ioutil.ReadFile("./task 3/dots/graph_3_3_1.dot")
	graph, _ := graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_3_1.svg")

	b, _ = ioutil.ReadFile("./task 3/dots/graph_3_3_2.dot")
	graph, _ = graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_3_2.svg")
}
func GenerateGraph_3_4() {
	g := graphviz.New()

	b, _ := ioutil.ReadFile("./task 3/dots/graph_3_4_2.dot")
	graph, _ := graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_4_2.svg")

	b, _ = ioutil.ReadFile("./task 3/dots/graph_3_4_3.dot")
	graph, _ = graphviz.ParseBytes(b)

	g.RenderFilename(graph, graphviz.SVG, "./task 3/rendered/graph_3_4_3.svg")
}

func GenerateGraph_4_1() {
	g := graphviz.New()
	graph, _ := g.Graph()

	nodes := GenerateNodesSlice(graph, "q", 8, GenerateTerminalMap(8))
	GenerateConnect(graph, nodes, 1, 2, "a")
	GenerateConnect(graph, nodes, 2, 3, "a")
	GenerateConnect(graph, nodes, 3, 4, "b")
	GenerateConnect(graph, nodes, 4, 1, "lmb")
	GenerateConnect(graph, nodes, 1, 4, "lmb")

	GenerateConnect(graph, nodes, 4, 5, "a")
	GenerateConnect(graph, nodes, 5, 8, "lmb")
	GenerateConnect(graph, nodes, 8, 5, "lmb")
	GenerateConnect(graph, nodes, 5, 6, "a")
	GenerateConnect(graph, nodes, 6, 7, "b")
	GenerateConnect(graph, nodes, 7, 8, "a")

	g.RenderFilename(graph, graphviz.SVG, "task 4/graph_4_1.svg")
}

func main() {
	GenerateGraph_1_1()
	GenerateGraph_1_2()
	GenerateGraph_1_3()
	GenerateGraph_1_4()

	GenerateGraph_2_1_1()
	GenerateGraph_2_1_2()
	GenerateGraph_2_1_3()

	GenerateGraph_2_2_1()
	GenerateGraph_2_2_2()
	GenerateGraph_2_2_3()

	GenerateGraph_2_3_1()
	GenerateGraph_2_3_2()
	GenerateGraph_2_3_3()

	GenerateGraph_2_4()

	GenerateGraph_3_1()
	GenerateGraph_3_2()
	GenerateGraph_3_3()
	GenerateGraph_3_4()

	GenerateGraph_4_1()
}
