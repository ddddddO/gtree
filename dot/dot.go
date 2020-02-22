package dot

import (
	"fmt"
)

/*
parseFileで読み込むファイル内は、以下形式
node: A, B, C...
edge: [A,B], [B,C], [C,D]...
*/
type DotGenerator interface {
	parseFile() ([]string, []string, error)
	generateDotFile([]string, []string) error
}

// 有向グラフ
type Digraph struct{ Path string }

// 無向グラフ
type Graph struct{ Path string }

func NewDotGenerator(isDigraph bool, path string) DotGenerator {
	if isDigraph {
		return Digraph{Path: path}
	}
	return Graph{Path: path}
}

func Run(generator DotGenerator) error {
	node, edge, err := generator.parseFile()
	if err != nil {
		return err
	}

	err = generator.generateDotFile(node, edge)
	if err != nil {
		return err
	}

	return nil
}

func (g Graph) parseFile() ([]string, []string, error) {
	fmt.Println("Graph")
	fmt.Println("FilePath:", g.Path)
	return nil, nil, nil
}
func (g Graph) generateDotFile(node, edge []string) error {
	return nil
}

func (dg Digraph) parseFile() ([]string, []string, error) {
	fmt.Println("Digraph")
	fmt.Println("FilePath:", dg.Path)
	return nil, nil, nil
}
func (dg Digraph) generateDotFile(node, edge []string) error {
	return nil
}
