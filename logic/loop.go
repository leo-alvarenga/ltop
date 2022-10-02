package logic

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/leo-alvarenga/ltop/io/shared"
	"github.com/leo-alvarenga/ltop/io/types"
	"golang.org/x/term"
)

func Loop() {
	for {
		clean := make(chan bool)
		table := types.NewMainInfoTable()

		clearStdout()
		w, _, err := term.GetSize(0)
		if err != nil {
			return
		}

		graph := getData(table)
		for _, line := range table.GetRows(w, graph) {
			println(line)
		}

		go cleanupJob(clean)
		time.Sleep(time.Second)

		<-clean
	}
}

func cleanupJob(done chan bool) {
	runtime.GC()
	done <- true
}

func clearStdout() {
	c := exec.Command("clear")

	c.Stdout = os.Stdout
	c.Run()
}

func getData(table *types.MainInfoTable) *types.Graph {
	keys, values, err := getMemInfo()
	if err != nil {
		return nil
	}

	for i, value := range values {
		key := keys[i]

		if strings.Contains(key, "swap") {
			label := strings.ReplaceAll(key, "swap", "")
			table.SwapColumn.AddRow(
				shared.PrettierMemLabels[key],
				shared.Float64ToString(value),
				shared.GetNewStyle(shared.ColorsByLabel[label], ""),
			)

			continue
		}

		table.MemColumn.AddRow(
			shared.PrettierMemLabels[key],
			shared.Float64ToString(value),
			shared.GetNewStyle(shared.ColorsByLabel[key], ""),
		)
	}

	sysKeys, sysValues, err := getSysInfo()
	if err != nil {
		return nil
	}

	for i, value := range sysValues {
		key := sysKeys[i]

		table.SysColumn.AddRow(
			key,
			value,
			shared.GetNewStyle("", ""),
		)
	}

	table.CalculateLength()
	return setupGraph(keys, values, table.TotalInnerLength)
}

func setupGraph(keys [9]string, values [9]float64, length int) *types.Graph {
	graph := types.NewGraph(length + 2)

	for i, value := range values {
		key := keys[i]

		if strings.Contains(key, "total") || strings.Contains(key, "available") {
			continue
		}

		if strings.Contains(key, "swap") {
			label := strings.ReplaceAll(key, "swap", "")
			graph.AddValue(value, values[6], shared.ColorsByLabel[label], "Swap :  ", 1)

			continue
		}

		graph.AddValue(value, values[0], shared.ColorsByLabel[key], "Memory: ", 0)
	}

	return graph
}
