package dot

import (
	"os"
	"text/template"
)

type Graph struct {
	Name      string
	IsDigraph bool
	Nodes     []Node
	Edges     [][]string
}

type Node struct {
	Title string
	Url   string
}

func NewGraph(name string, isDigraph bool, nodes []Node, edges [][]string) Graph {
	return Graph{
		Name:      name,
		IsDigraph: isDigraph,
		Nodes:     nodes,
		Edges:     edges,
	}
}

func Gen(g Graph) error {
	tmpl, err := template.ParseFiles("../../dot.template")
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, g)
	if err != nil {
		return err
	}

	return nil
}
