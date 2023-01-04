package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type document struct {
	path    string
	content []string
}

func errHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func convertFile(path string) document {
	var wordsC []string

	f, err := os.Open(path)
	errHandling(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordsC = append(wordsC, strings.ToLower(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	words := document{path: path, content: wordsC}

	return words
}

func search(words []document, searchWord string) []string {

	var result []string

	for i := range words {
		for a := range words[i].content {
			if words[i].content[a] == searchWord {
				result = append(result, words[i].path)
			}
		}
	}

	return result
}

func main() {
	var words []document

	//example - ./main dir_name_with_files_inside_it search_query
	files, err := ioutil.ReadDir(os.Args[1])
	errHandling(err)

	for _, file := range files {
		words = append(words, convertFile("files/"+file.Name()))
	}

	result := search(words, os.Args[2])

	for i := range result {
		fmt.Println(result[i])
	}
}
