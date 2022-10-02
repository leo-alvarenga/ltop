package main

import (
	"strings"

	"github.com/leo-alvarenga/ltop/io/shared"
	"github.com/leo-alvarenga/ltop/io/types"
	"github.com/leo-alvarenga/ltop/logic"
	"golang.org/x/term"
)

func main() {
	table := types.NewMainInfoTable()

	w, _, err := term.GetSize(0)
	if err != nil {
		return
	}

	graph := getData(table)
	for _, line := range table.GetRows(w, graph) {
		println(line)
	}
}

func getData(table *types.MainInfoTable) *types.Graph {
	data, err := logic.GetMemInfo()
	if err != nil {
		return nil
	}

	for key, value := range data {
		if strings.Contains(key, "swap") {
			label := strings.ReplaceAll(key, "swap", "")
			table.SwapColumn.AddRow(
				shared.PrettierMemLabels[key],
				shared.Float64ToString(value),
				shared.GetNewStyle(shared.ColorsByLabel[label], ""),
			)

			continue
		}

		table.MemColumn.AddRow(
			shared.PrettierMemLabels[key],
			shared.Float64ToString(value),
			shared.GetNewStyle(shared.ColorsByLabel[key], ""),
		)
	}

	sysdata, err := logic.GetSysInfo()
	if err != nil {
		return nil
	}

	for key, value := range sysdata {
		table.SysColumn.AddRow(
			key,
			value,
			shared.GetNewStyle("", ""),
		)
	}

	table.CalculateLength()
	return setupGraph(data, table.TotalInnerLength)
}

func setupGraph(data map[string]float64, length int) *types.Graph {
	graph := types.NewGraph(length + 2)

	for key, value := range data {
		if strings.Contains(key, "total") || strings.Contains(key, "available") {
			continue
		}

		if strings.Contains(key, "swap") {
			label := strings.ReplaceAll(key, "swap", "")
			graph.AddValue(value, data["swaptotal"], shared.ColorsByLabel[label], "Swap :  ", 1)

			continue
		}

		graph.AddValue(value, data["total"], shared.ColorsByLabel[key], "Memory: ", 0)
	}

	return graph
}
