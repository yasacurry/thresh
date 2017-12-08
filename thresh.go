package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const defaultFile = "data.csv"
const layout = "15:04 2006-01-02"

func main() {
	fileName := flag.String("f", defaultFile, "tail file name")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Seek(0, os.SEEK_END)

	c := csv.NewReader(f)

	for {
		s, err := c.Read()
		if len(s) > 0 {
			formatPrint(s)
		}
		if err == io.EOF {
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func formatPrint(record []string) {
	t, err := time.Parse("2006-01-02 15:04:05 +0900 MST", record[0])
	if err != nil {
		log.Fatal(err)
	}

	if record[7] == "" {
		fmt.Printf("@%v / %v %v\n%v\n\n", record[4], record[5], t.Format(layout), record[6])
	} else {
		fmt.Printf("@%v / %v RT from @%v %v\n%v\n\n", record[4], record[5], record[7], t.Format(layout), record[6])
	}
}
