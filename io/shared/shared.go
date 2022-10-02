package shared

import (
	gostyle "github.com/leo-alvarenga/go-easy-style"
)

func GetNewStyle(txt, bg string, styles ...string) *gostyle.TextStyle {
	s := new(gostyle.TextStyle)
	s.New(txt, bg, styles...)

	return s
}
