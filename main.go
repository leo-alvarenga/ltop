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

	data, err := logic.GetMemInfo()
	if err != nil {
		return
	}

	for key, value := range data {
		if strings.Contains(key, "swap") {
			label := strings.ReplaceAll(key, "swap", "")

			table.SwapColumn.AddRow(
				shared.PrettierMemLabels[key],
				shared.Float64ToString(value),
				shared.GetNewStyle(shared.ColorsByLabel[label], ""),
			)
		} else {
			table.MemColumn.AddRow(
				shared.PrettierMemLabels[key],
				shared.Float64ToString(value),
				shared.GetNewStyle(shared.ColorsByLabel[key], ""),
			)
		}
	}

	for _, line := range table.GetRows(w) {
		println(line)
	}
}
