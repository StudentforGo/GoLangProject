package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(inputFile string, outputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file")
		return
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file")
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	words := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Fields(line)...)
	}

	for i := len(words) - 1; i >= 0; i-- {
		output.WriteString(words[i] + " ")
	}
}

func rearrangeEvenOdd(inputFile string, outputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file")
		return
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file")
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	even := true

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			if even {
				output.WriteString(word + " ")
			}
			even = !even
		}
	}
}
func cut(inputFile string, outputFile string, maxLength int) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file")
		return
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file")
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			if len(word) > maxLength {
				word = word[:maxLength]
			}
			output.WriteString(word + " ")
		}
	}
}

func main() {
	inputFile := "inputfile.txt"
	reverseOutputFile := "reversefile"
	rearrangeOutputFile := "rearrangefile"
	cutOutputFile := "cutfile"
	reverse(inputFile, reverseOutputFile)
	rearrangeEvenOdd(inputFile, rearrangeOutputFile)
	cut(inputFile, cutOutputFile, 8)
}
