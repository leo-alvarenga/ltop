package types

import (
	"fmt"
	"strings"

	"github.com/leo-alvarenga/ltop/io/shared"
)

type Graph struct {
	Labels      []string
	Values      [][]string
	originals   [][]string
	GraphLength int
}

func NewGraph(length int) *Graph {
	g := new(Graph)
	g.GraphLength = length - 4

	g.Labels = append(g.Labels, "Memory: ", "Swap:   ")
	g.Values = append(g.Values, []string{}, []string{})
	g.originals = append(g.originals, []string{}, []string{})

	return g
}

func (g *Graph) AddValue(value, total float64, color, label string, row int) {
	v := shared.Float64ToString(value)
	abs := int(value)
	length := (abs * (g.GraphLength - len(label))) / int(total)

	if length < len(v) {
		v = ""
	}

	fill := strings.Repeat(" ", length-len(v))

	if length == 0 {
		fill = " "
	}

	g.Values[row] = append(
		g.Values[row],
		shared.GetNewStyle("", color).Style(fmt.Sprintf("%s%s", fill, v)),
	)
	g.originals[row] = append(
		g.originals[row],
		fmt.Sprintf("%s%s", fill, v),
	)
}

func (g *Graph) GetValues() (bars []string) {
	for i, row := range g.Values {
		r := g.Labels[i]
		orig := g.Labels[i]

		for j, v := range row {
			r += v
			orig += g.originals[i][j]
		}

		if len(orig) <= g.GraphLength {
			count := 1

			if len(orig) > g.GraphLength {
				count += g.GraphLength - len(orig)
			}

			r += strings.Repeat(" ", count)
		}
		bars = append(bars, fmt.Sprintf("%s %s %s", shared.Vertical, r, shared.Vertical))
	}

	return
}
