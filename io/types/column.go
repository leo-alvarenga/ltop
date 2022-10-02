package types

import (
	"fmt"
	"strings"

	gostyle "github.com/leo-alvarenga/go-easy-style"
	"github.com/leo-alvarenga/ltop/io/shared"
)

type Column struct {
	labels, values []string
	styles         []*gostyle.TextStyle
	Length         int
	RowCount       int
}

func NewColumn() *Column {
	return new(Column)
}

func (c *Column) AddRow(label, value string, style *gostyle.TextStyle) {
	c.labels = append(c.labels, label)
	c.values = append(c.values, value)
	c.styles = append(c.styles, style)

	l := len(c.getDummyRow(len(c.labels) - 1))
	if l > c.Length {
		c.Length = l
	}

	c.RowCount++
}

func (c *Column) getDummyRow(targetRow int) string {
	if targetRow < 0 || targetRow > len(c.labels) {
		return " "
	}

	return fmt.Sprintf("%s: %s", c.labels[targetRow], c.values[targetRow])
}

func (c *Column) GetFormattedRow(targetRow, columnPosition int) string {
	label, value := "  ", " "

	if targetRow >= 0 && targetRow < c.RowCount {
		if c.styles[targetRow] == nil {
			label = c.labels[targetRow] + ":"
			value = c.values[targetRow]
		} else {
			label = c.styles[targetRow].Style(c.labels[targetRow]) + ":"
			value = c.styles[targetRow].Style(c.values[targetRow])
		}
	} else if c.RowCount == targetRow {
		if columnPosition == 0 {
			return getHorizontalBorder("endleft", c.Length)
		} else if columnPosition > 1 {
			return getHorizontalBorder("endright", c.Length)
		}

		return getHorizontalBorder("bottom", c.Length)
	}

	dif := c.Length - len(c.getDummyRow(targetRow))
	filler := ""

	if dif > 0 {
		filler = strings.Repeat(" ", dif)
	}

	return fmt.Sprintf("%s %s%s", label, filler, value)
}

func getHorizontalBorder(position string, innerLenght int) string {
	var l, fill, r string

	innerLenght += 2
	switch position {
	case "top":
		l = shared.UpLeft
		r = shared.UpRight
	case "bottom":
		l = shared.DownLeft
		r = shared.DownRight
	case "endleft":
		l = shared.DownLeft
		r = ""
	case "endright":
		l = ""
		r = shared.DownRight
	default:
		innerLenght--
		l = shared.SepLeft
		r = shared.SepRight
	}

	fill = strings.Repeat(shared.Horizontal, innerLenght)
	return fmt.Sprintf("%s%s%s", l, fill, r)
}
