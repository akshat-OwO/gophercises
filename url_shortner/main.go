package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "path.csv", ".csv file containing paths in format 'path,url'")
	// jsonFile := flag.String("json", "path.json", ".json file containing paths in format [{'path': 'url'}]")
	// yamlFile := flag.String("yaml", "path.yaml", ".yaml file containing paths")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("couldn't find csv with path %s", *csvFile))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal("failed to parse the provided csv")
	}

	paths := parseCsv(lines)

	http.HandleFunc("/", healthCheck())
	http.HandleFunc("/csv/", csvMapping(paths))

	fmt.Println("Starting server on port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("couldn't start the server")
	}
}

func healthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		io.WriteString(w, "Hello World")
	}
}

func csvMapping(paths []path) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortPath := strings.TrimPrefix(r.URL.Path, "/csv")
		for _, pa := range paths {
			if pa.p == shortPath {
				http.Redirect(w, r, pa.u, http.StatusMovedPermanently)
			}
		}
		http.NotFound(w, r)
	}
}

type path struct {
	p string
	u string
}

func parseCsv(lines [][]string) []path {
	ret := make([]path, len(lines))
	for i, line := range lines {
		ret[i] = path{
			p: line[0],
			u: line[1],
		}
	}
	return ret
}
