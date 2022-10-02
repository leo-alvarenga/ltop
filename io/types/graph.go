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

	l := int((value * float64(g.GraphLength-len(label))) / total)

	if l == 0 {
		v = ""
	}

	var bar string
	if l > len(v) {
		bar = v
		l -= len(v)
	}

	bar = strings.Repeat(" ", l) + bar

	g.Values[row] = append(
		g.Values[row],
		shared.GetNewStyle("", color).Style(bar),
	)
	g.originals[row] = append(g.originals[row], bar)
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
