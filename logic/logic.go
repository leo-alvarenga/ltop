package logic

import (
	"bufio"
	"os"
	"strings"

	"github.com/leo-alvarenga/ltop/io/shared"
)

func getMemInfo() (keys [9]string, values [9]float64, err error) {
	file, err := os.Open(memInfoFile)
	if err != nil {
		return
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		key, value := shared.GetKeyAndValueFromString(scan.Text())
		keys, values = setValue(key, shared.StringToFloat64(value), keys, values)
	}

	keys, values = setValue("used", values[0]-values[5]-values[2]-values[3], keys, values)
	keys, values = setValue("swapUsed", values[6]-values[8], keys, values)

	return
}

func getSysInfo() (keys [3]string, values [3]string, err error) {
	file, err := os.Open(sysInfoFile)

	if err != nil {
		return
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		key, value := shared.GetKeyAndValueFromString(strings.ReplaceAll(scan.Text(), " ", ""))

		switch key {
		case "NAME":
			keys[0] = "Distro"
			values[0] = value
		case "VERSION":
			keys[1] = "Version"
			values[1] = value
		case "BUILD_ID":
			keys[1] = "Version"
			values[1] = value
		}
	}

	uptime, err := getUptime()
	if err != nil {
		return
	}

	keys[2] = "Uptime"
	values[2] = uptime

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

func setValue(key string, value float64, keys [9]string, values [9]float64) ([9]string, [9]float64) {
	switch key {
	case "MemTotal":
		keys[0] = "total"
		values[0] = value
	case "used":
		keys[1] = "used"
		values[1] = value
	case "Cached":
		keys[2] = "cach"
		values[2] += value
	case "SReclaimable":
		keys[2] = "cach"
		values[2] += value
	case "Shmem":
		keys[3] = "shrd"
		values[3] = value
	case "Buffers":
		keys[4] = "buff"
		values[4] = value
	case "MemFree":
		keys[5] = "free"
		values[5] = value

	case "SwapTotal":
		keys[6] = "swaptotal"
		values[6] = value
	case "swapUsed":
		keys[7] = "swapused"
		values[7] = value
	case "SwapFree":
		keys[8] = "swapfree"
		values[8] = value
	}

	return keys, values
}
