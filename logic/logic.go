package logic

import (
	"bufio"
	"os"

	"github.com/leo-alvarenga/ltop/io/shared"
)

func GetMemInfo() (data map[string]float64, err error) {
	file, err := os.Open(memInfoFile)
	if err != nil {
		return nil, err
	}

	m := make(map[string]float64)
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		key, value := shared.GetKeyAndValueFromString(scan.Text())
		setValue(key, shared.StringToFloat64(value), m)
	}

	setValue("used", m["total"]-m["free"]-m["buff"]-m["cach"], m)
	setValue("swapUsed", m["swaptotal"]-m["swapfree"], m)

	return m, nil
}

func setValue(key string, value float64, m map[string]float64) {
	switch key {
	case "MemTotal":
		m["total"] = value
	case "Buffers":
		m["buff"] = value
	case "Shmem":
		m["shrd"] = value
	case "Cached":
		m["cach"] += value
	case "SReclaimable":
		m["cach"] += value
	case "used":
		m["used"] = value
	case "MemFree":
		m["free"] = value
	case "MemAvailable":
		m["available"] = value

	case "SwapTotal":
		m["swaptotal"] = value
	case "SwapFree":
		m["swapfree"] = value
	case "swapUsed":
		m["swapused"] = value
	}
}
