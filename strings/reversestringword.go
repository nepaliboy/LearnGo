package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	usage()
	revstringword(os.Args[1], os.Args[2], os.Args[3])
}

func revstringword(filenamei string, char string, filenameo string) {
	fp, err := os.Open(filenamei)
	if err != nil {
		fmt.Printf("error operning file %v", err)
		os.Exit(1)
	}
	fp1, err1 := os.Create(filenameo)
	if err1 != nil {
		fmt.Printf("can't create a file %v", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if char == "c" {
			c := []rune(line)
			for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
				c[i], c[j] = c[j], c[i]
			}
			output := string(c)
			fmt.Printf("%s\n", output)
			fmt.Fprintf(fp1, "%s \n", output)
		} else if char == "w" {
			w := strings.Fields(line)
			for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
				w[i], w[j] = w[j], w[i]
			}
			words := strings.Join(w, " ")
			fmt.Printf("%s \n", words)
			fmt.Fprintf(fp1, "%s \n", words)
		}
	}
	defer fp1.Close()
}

func usage() {
	if len(os.Args) < 4 {
		fmt.Printf(" \n Usage %v <inputfile> <c/w> <outputfile> \n", os.Args[0])
		os.Exit(1)

	}
}
