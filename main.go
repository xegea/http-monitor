package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := readCsvFile("urls.csv")

	for _, v := range urls {
		get(v[0])
	}
}

func get(url string) {
	req, _ := http.NewRequest("GET", url, nil)

	var start time.Time
	var duration time.Duration

	req = req.WithContext(req.Context())
	cc := &http.Client{Timeout: time.Second * 30}

	start = time.Now()
	res, err := cc.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	duration = time.Since(start)

	// b, _ := io.ReadAll(res.Body)
	// res.Body.Close()

	fmt.Printf("%s\t%d\t%v\n", url, res.StatusCode, duration)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
