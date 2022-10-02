package shared

import (
	"fmt"
	"strconv"
	"strings"

	gostyle "github.com/leo-alvarenga/go-easy-style"
)

func GetNewStyle(txt, bg string, styles ...string) *gostyle.TextStyle {
	s := new(gostyle.TextStyle)
	s.New(txt, bg, styles...)

	return s
}

func GetBiggest(numbers ...int) (index, biggest int) {
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

func Float64ToString(value float64) string {
	const kilo float64 = 1024

	if value <= 0 {
		return "0k"
	}

	abrvs := [9]string{"k", "M", "G", "T", "P", "E", "Z"}
	for _, u := range abrvs {
		if value < kilo {
			return fmt.Sprintf("%.2f%s", value, u)
		}

		value /= kilo
	}

	return "A lot!"
}

func GetKeyAndValueFromString(raw string) (key, value string) {

	// if B is on the last position, remove the unit (kB, mB, etc...)
	if raw[len(raw)-1] == 'B' {
		raw = raw[:len(raw)-2]
	}

	raw = strings.ReplaceAll(raw, "\"", "")
	raw = strings.ReplaceAll(raw, " ", "")

	pair := strings.Split(raw, ":")
	if len(pair) == 1 {
		pair = strings.Split(raw, "=")
	}

	return pair[0], pair[1]
}

func StringToFloat64(data string) float64 {
	if data == "" {
		return 0
	}

	value, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return 0
	}

	return value
}

func TimeToString(time int) string {
	hrs := time / 3600
	time = time % 3600

	min := time / 60
	time = time % 60

	if hrs > 24 {
		days := hrs / 24
		hrs = hrs % 24

		return fmt.Sprintf("%d days %dh %dmin %ds", days, hrs, min, time)
	}

	return fmt.Sprintf("%dh %dmin %ds", hrs, min, time)
}
