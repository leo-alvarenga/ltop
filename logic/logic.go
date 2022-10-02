package logic

import (
	"bufio"
	"os"
	"strings"

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

func GetSysInfo() (data map[string]string, err error) {
	file, err := os.Open(sysInfoFile)

	if err != nil {
		return
	}

	defer file.Close()
	data = make(map[string]string)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		key, value := shared.GetKeyAndValueFromString(strings.ReplaceAll(scan.Text(), " ", ""))

		switch key {
		case "NAME":
			data["DistroName"] = value
		case "VERSION":
			data["Version"] = value
		case "BUILD_ID":
			data["Version"] = value
		}
	}

	uptime, err := getUptime()
	if err != nil {
		return
	}

	data["Uptime"] = uptime

	return
}

func getUptime() (string, error) {
	file, err := os.Open(uptimeFile)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Scan()

	uptime := int(shared.StringToFloat64(strings.Split(scan.Text(), " ")[0]))

	return shared.TimeToString(uptime), nil
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

	case "SwapTotal":
		m["swaptotal"] = value
	case "SwapFree":
		m["swapfree"] = value
	case "swapUsed":
		m["swapused"] = value
	}
}
