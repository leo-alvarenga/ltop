package main

import (
	"github.com/leo-alvarenga/ltop/io/shared"
	"github.com/leo-alvarenga/ltop/io/types"
)

func main() {
	table := types.NewMainInfoTable()
	table.MemColumn.AddRow("o", "adsasdasadsasdasadsasdas", shared.GetNewStyle("blue", ""))
	table.MemColumn.AddRow("a", "adsasdas", shared.GetNewStyle("blue", ""))
	table.SwapColumn.AddRow("b", "adsasdas", shared.GetNewStyle("blue", ""))
	table.SysColumn.AddRow("bruhhh", "adsasdas", shared.GetNewStyle("blue", ""))

	table.GetRows()
}
