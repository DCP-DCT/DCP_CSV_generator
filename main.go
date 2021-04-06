package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	fn := os.Args[1]
	file, err := os.Open(fn)
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	run := RunConfig{}
	err = json.Unmarshal([]byte(os.Args[2]), &run)
	if err != nil {
		panic(err.Error())
	}

	data, _ := ioutil.ReadAll(file)

	var nodes []CtNode
	err = json.Unmarshal(data, &nodes)
	if err != nil {
		panic(err.Error())
	}

	parts := strings.Split(fn, ".")

	output, err := os.Create(parts[0] + ".csv")
	if err != nil {
		panic(err.Error())
	}

	defer output.Close()

	w := csv.NewWriter(output)
	defer w.Flush()

	headers := []string{"N", "L", "DT", "TTL", "T", "B", "EU", "R", "D", "PK", "IU", "PD"}
	values := getDiagnosisData(nodes, run)

	if err := w.Write(headers); err != nil {
		panic(err.Error())
	}

	for _, value := range values {
		err := w.Write(value)
		if err != nil {
			panic(err.Error())
		}
	}

}

func getDiagnosisData(nodes []CtNode, run RunConfig) [][]string {
	var data [][]string
	for _, node := range nodes {
		data = append(data, []string{
			strconv.Itoa(run.N),
			strconv.Itoa(run.L),
			strconv.Itoa(run.DT),
			strconv.Itoa(run.TTL),
			strconv.Itoa(run.T),
			strconv.Itoa(node.Diagnosis.NumberOfBroadcasts),
			strconv.Itoa(node.Diagnosis.NumberOfUpdates),
			strconv.Itoa(node.Diagnosis.NumberOfRejectedDueToThreshold),
			strconv.Itoa(node.Diagnosis.NumberOfDuplicates),
			strconv.Itoa(node.Diagnosis.NumberOfPkMatches),
			strconv.Itoa(node.Diagnosis.NumberOfInternalUpdates),
			strconv.Itoa(node.Diagnosis.NumberOfPacketsDropped)},
		)
	}

	return data
}
