package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var searchRegexp *regexp.Regexp
var fileExtRegexp *regexp.Regexp

func main() {
	rootFolder := ""
	flag.StringVar(&rootFolder, "F", "", "Folder to search")

	rootFile := ""
	flag.StringVar(&rootFile, "f", "", "File to search")

	regxExpr := ""
	flag.StringVar(&regxExpr, "e", "", "Regex expression for search")

	fileExt := ""
	flag.StringVar(&fileExt, "fe", "", "Regex expression for the file name for search")

	flag.Parse()

	fmt.Printf("Params folder(%s), file(%s), regex expression(%s) \n", rootFolder, rootFile, regxExpr)

	if len(regxExpr) == 0 {
		fmt.Println("Regex expression missing. Run ggrep -h for usage")
		return
	}

	//searchRegexp = regexp.MustCompile(regxExpr)
	var err error

	searchRegexp, err = compileRegex(regxExpr)

	if err != nil {
		return
	}

	fileExtRegexp, err = compileRegex(fileExt)

	if err != nil {
		return
	}

	if len(rootFile) > 0 {
		searchInFile(rootFile)
	} else if len(rootFolder) > 0 {
		searchInPath(rootFolder)
	} else {
		fmt.Println("Searching in STDIN")
		searchInStdIn()
	}

	os.Exit(0)
}

func compileRegex(regStr string) (*regexp.Regexp, error) {

	var res *regexp.Regexp
	var err error

	res = regexp.MustCompile(regStr)

	return res, err
}

func searchInFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	search(*scanner, fileName)

}

func searchInPath(folderPath string) {

	filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if !info.IsDir() {
			searchInFile(path)
		}

		return nil
	})
}

func searchInStdIn() {
	scanner := bufio.NewScanner(os.Stdin)
	search(*scanner, "STDIN")
}

func search(scanner bufio.Scanner, fileName string) {

	fileMatch := fileExtRegexp.FindAllString(fileName, -1)

	if !(len(fileMatch) > 0) {
		return
	}

	currentLineNumber := 0
	for scanner.Scan() {
		currentLine := scanner.Text()

		match := searchRegexp.FindAllString(currentLine, -1)

		if len(match) > 0 {
			fmt.Printf("File %s [line %d] [Match %s] %s\n", fileName, currentLineNumber, match, currentLine)
		}

		currentLineNumber = currentLineNumber + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
