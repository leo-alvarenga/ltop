package types

import (
	"fmt"
	"strings"

	gostyle "github.com/leo-alvarenga/go-easy-style"
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

	c.RowCount++

	l := len(c.getDummyRow(len(c.labels) - 1))
	if l > c.Length {
		c.Length = l
	}
}

func (c *Column) getDummyRow(targetRow int) string {
	if targetRow < 0 || targetRow > len(c.labels) {
		return " "
	}

	return fmt.Sprintf("%s: %s", c.labels[targetRow], c.values[targetRow])
}

func (c *Column) GetFormattedRow(targetRow int) string {
	var label, value string

	if targetRow >= 0 && targetRow < len(c.labels) {
		label = c.styles[targetRow].Style(c.labels[targetRow]) + ":"
		value = c.styles[targetRow].Style(c.values[targetRow])
	}

	dif := c.Length - len(c.getDummyRow(targetRow))
	filler := ""

	if dif > 0 {
		filler = strings.Repeat(" ", dif)
	}

	return fmt.Sprintf("%s %s%s", label, filler, value)
}
