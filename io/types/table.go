package types

import (
	"fmt"

	"github.com/leo-alvarenga/ltop/io/shared"
)

type MainInfoTable struct {
	SysColumn  *Column
	MemColumn  *Column
	SwapColumn *Column
	Length     int
}

func NewMainInfoTable() *MainInfoTable {
	t := new(MainInfoTable)
	t.SysColumn, t.MemColumn, t.SwapColumn = NewColumn(), NewColumn(), NewColumn()

	return t
}

func (t *MainInfoTable) GetRows() (table []string) {
	t.CalculateLength()

	// teorically, since there are more info on memory than on swap mem, this just not happen
	// but, then again, just to be safe...
	if t.SwapColumn.RowCount > t.MemColumn.RowCount {
		t.MemColumn, t.SwapColumn = t.SwapColumn, t.MemColumn
	}

	for i := 0; i < t.MemColumn.RowCount; i++ {
		sys, mem, swp := -1, i, -1

		if i < t.SysColumn.RowCount {
			sys = i
		}

		if i < t.SwapColumn.RowCount {
			swp = i
		}

		println(t.GetCombinedRows(sys, mem, swp))
	}

	return
}

func (t *MainInfoTable) GetCombinedRows(sys, mem, swp int) string {
	sep := shared.Vertical

	memSide := t.MemColumn.GetFormattedRow(mem)
	sysSide := t.SysColumn.GetFormattedRow(sys)
	swpSide := ""

	memSide = fmt.Sprintf("%s %s ", memSide, sep)
	sysSide = fmt.Sprintf(" %s %s %s ", sep, sysSide, sep)

	if swp >= 0 {
		swpSide = t.SwapColumn.GetFormattedRow(swp)
		swpSide = fmt.Sprintf("%s %s ", swpSide, sep)
	}

	return sysSide + memSide + swpSide
}

func (t *MainInfoTable) CalculateLength() int {
	i, l := getBiggest(t.SysColumn.Length, t.MemColumn.Length, t.SwapColumn.Length)

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

	return t.Length
}

func getBiggest(numbers ...int) (index, biggest int) {
	if len(numbers) == 0 {
		return -1, -1
	}

	biggest = numbers[0]
	index = 0

	for i, n := range numbers {
		if n > biggest {
			biggest = n
			index = i
		}
	}

	return
}
