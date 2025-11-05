package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type TestCase struct {
	Name   string
	Input  string
	Output string
	Flags  Flags
}

func parseFlags(filename string) Flags {
	commandFlags := Flags{}

	base := filepath.Base(filename)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	var end int
	for i := 0; i < len(name); i++ {
		if !unicode.IsDigit(rune(name[i])) {
			break
		}
		end = i

	}
	name = name[end+1:]
	for i := 0; i < len(name); i++ {
		switch name[i] {
		case 'c':
			commandFlags.C = true
		case 'd':
			commandFlags.D = true
		case 'u':
			commandFlags.U = true
		case 'i':
			commandFlags.I = true
		case 'f':
			start := i + 1
			end := start
			for end < len(name) && name[i+1] > '0' && name[i+1] < '9' {
				end++
			}
			if end > start {
				num, _ := strconv.Atoi(name[start:end])
				commandFlags.F = num
				i = end - 1
			}
		case 's':
			start := i + 1
			end := start
			for end < len(name) && name[i+1] >= '0' && name[i+1] <= '9' {
				end++
			}
			if end > start {
				num, _ := strconv.Atoi(name[start:end])
				commandFlags.S = num
				i = end - 1
			}

		}
	}
	fmt.Println(commandFlags, "флаги!!")
	return commandFlags

}

func LoadTestCases() ([]TestCase, error) {
	var testcases []TestCase

	inputFiles, err := filepath.Glob("./acceptance-tests/input/*.txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	outputFiles, err := filepath.Glob("./acceptance-tests/output/*.txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for i, inputFile := range inputFiles {
		input, err := os.ReadFile(inputFile)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		output, _ := os.ReadFile(outputFiles[i])

		flags := parseFlags(inputFile)
		fmt.Println(flags)
		testcases = append(testcases, TestCase{Name: filepath.Base(inputFile),
			Input:  string(input),
			Output: string(output),
			Flags:  flags})

	}
	return testcases, nil
}
