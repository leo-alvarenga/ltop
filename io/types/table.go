package types

import (
	"fmt"
	"strings"

	"github.com/leo-alvarenga/ltop/io/shared"
)

type MainInfoTable struct {
	SysColumn        *Column
	MemColumn        *Column
	SwapColumn       *Column
	Length           int
	TotalLength      int
	TotalInnerLength int
}

func NewMainInfoTable() *MainInfoTable {
	t := new(MainInfoTable)
	t.SysColumn, t.MemColumn, t.SwapColumn = NewColumn(), NewColumn(), NewColumn()

	t.MemColumn.AddRow("Memory", "", shared.GetNewStyle("cyan", "", "bold", "underline"))
	t.SwapColumn.AddRow("Swap", "", shared.GetNewStyle("cyan", "", "bold", "underline"))

	return t
}

func (t *MainInfoTable) GetRows(terminalWidth int, graph *Graph) (table []string) {
	t.CalculateLength()

	if terminalWidth < t.TotalLength {
		println("Too small...", terminalWidth, t.TotalLength)
		return
	}

	// teorically, since there are more info on memory than on swap mem, this just not happen
	// but, then again, just to be safe...
	if t.SwapColumn.RowCount > t.MemColumn.RowCount {
		t.MemColumn, t.SwapColumn = t.SwapColumn, t.MemColumn
	}

	table = append(table, t.getHeader()...)
	table = append(table, getHorizontalBorder("sep", t.TotalInnerLength))
	table = append(table, graph.GetValues()...)
	table = append(table, getHorizontalBorder("sep", t.TotalInnerLength))
	for i := 0; i <= t.MemColumn.RowCount; i++ {
		sys, mem, swp := -1, i, -1

		if i <= t.SysColumn.RowCount {
			sys = i
		}

		if i <= t.SwapColumn.RowCount {
			swp = i
		}

		table = append(table, t.GetCombinedRows(sys, mem, swp))
	}

	return
}

func (t *MainInfoTable) GetCombinedRows(sys, mem, swp int) string {
	sep := shared.Vertical

	sysSide := t.SysColumn.GetFormattedRow(sys, 0)
	memSide := t.MemColumn.GetFormattedRow(mem, 1)
	swpSide := t.SwapColumn.GetFormattedRow(swp, 2)

	if sys < t.SysColumn.RowCount && sys != -1 {
		sysSide = fmt.Sprintf("%s %s ", sep, sysSide)
	}

	if swp < t.SwapColumn.RowCount && swp != -1 {
		swpSide = fmt.Sprintf(" %s %s ", swpSide, sep)
	}

	if mem >= 0 && mem != t.MemColumn.RowCount {
		memSide = fmt.Sprintf("%s %s %s", sep, memSide, sep)
	}

	return sysSide + memSide + swpSide
}

func (t *MainInfoTable) CalculateLength() int {
	i, l := shared.GetBiggest(t.SysColumn.Length, t.MemColumn.Length, t.SwapColumn.Length)

	t.Length = l
	switch i {
	case 0:
		t.MemColumn.Length = l
		t.SwapColumn.Length = l
	case 1:
		t.SysColumn.Length = l
		t.SwapColumn.Length = l
	case 2:
		t.SysColumn.Length = l
		t.MemColumn.Length = l
	}

	t.TotalLength = (t.Length * 3) + 13
	t.TotalInnerLength = t.TotalLength - 6
	return t.Length
}

func (t *MainInfoTable) getHeader() []string {
	sep := shared.Vertical
	msg := "ltop - System and Memory info"

	length := t.TotalInnerLength - len(msg) - 2
	return []string{
		getHorizontalBorder("top", t.TotalInnerLength-t.TotalInnerLength%2),
		fmt.Sprintf(
			"%s%s %s %s%s ",
			sep,
			strings.Repeat(" ", length/2),
			shared.GetNewStyle("purple", "", "bold").Style(msg),
			strings.Repeat(" ", length/2+t.TotalInnerLength%2),
			sep,
		),
	}
}
