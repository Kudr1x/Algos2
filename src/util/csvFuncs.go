package util

import (
	"encoding/csv"
	"os"
	"strings"
)

func WriteCSV(fName string, data string) {
	file, _ := os.OpenFile("table/"+fName+".csv", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()

	writer := csv.NewWriter(file)
	reader := csv.NewReader(file)

	all, _ := reader.ReadAll()

	arr := strings.Split(data, ",")

	all = append(all, arr)
	file.Seek(0, 0)

	writer.WriteAll(all)

	defer writer.Flush()
}

func ReadCSV(fName string) []string {
	file, _ := os.Open("table/" + fName + ".csv")
	defer file.Close()

	reader := csv.NewReader(file)

	all, _ := reader.ReadAll()

	return all[0]
}
