package shared

const (
	Horizontal string = "─"
	Vertical   string = "│"

	SepUp    string = "├"
	SepDown  string = "┤"
	SepLeft  string = "├"
	SepRight string = "┤"

	UpLeft    string = "╭"
	UpRight   string = "╮"
	DownLeft  string = "╰"
	DownRight string = "╯"
)

var PrettierMemLabels = map[string]string{
	"total":     "Total",
	"buff":      "Buffers",
	"shrd":      "Shared memory",
	"cach":      "Cached memory",
	"used":      "Used",
	"free":      "Free",
	"available": "Available",
	"swaptotal": "Total",
	"swapfree":  "Free",
	"swapused":  "Used",
}

var ColorsByLabel = map[string]string{
	"used": "red",
	"buff": "yellow",
	"shrd": "purple",
	"cach": "blue",
	"free": "green",
}
