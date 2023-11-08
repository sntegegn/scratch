package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/Sntegegn/scratch/learning_go/ch9/formatter"
	"github.com/Sntegegn/scratch/learning_go/ch9/math"
)

func buildGzipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		r.Close()
		gr.Close()
	}, nil
}

func countLetters(r io.Reader) (map[string]int, error) {
	result := map[string]int{}
	b := make([]byte, 2048)
	for {
		n, err := r.Read(b)
		for _, v := range b[:n] {
			if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
				result[string(v)]++
			}
		}
		if err == io.EOF {
			return result, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func main() {
	/* fileName := "test.txt.gz"
	r, closer, err := buildGzipReader(fileName)
	if err != nil {
		fmt.Printf("Error occured %v", err)
		return
	}
	defer closer()
	result, err := countLetters(r)
	if err != nil {
		fmt.Printf("Error occured while counting letters %v", err)
		return
	}
	fmt.Println(result) */

	fileName := "test.txt"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening files %v", err)
		return
	}
	defer f.Close()
	result, err := countLetters(f)
	if err != nil {
		fmt.Printf("Error occured while counting letters %v", err)
		return
	}
	fmt.Println(result)

	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
