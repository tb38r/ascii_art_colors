package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// error handling
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Println("EX: go run . something --color=<color>")
		os.Exit(0)
	}

	// opening the file in read-only mode. The file must exist (in the current working directory)
	file, _ := os.Open("chars.txt")

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanned := bufio.NewScanner(file)

	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}

	file.Close()
	// empty
	ascMap := make(map[int][]string)
	id := 31
	// for i, _ := range lines {
	for _, line := range lines {
		if string(line) == "" {
			id++
		} else {
			ascMap[id] = append(ascMap[id], line)
		}
	}

	args := os.Args[1]

	for i := 0; i < len(args); i++ {
		if args[i] == 92 && args[i+1] == 110 {
			Newline(string(args[:i]), ascMap)
			Newline(string(args[i+2:]), ascMap)

		}
	}

	if !strings.Contains(args, "\\n") {
		Newline(args, ascMap)
	}
}

// Newline() prints the ascii art on the terminal horizontally
func Newline(p string, y map[int][]string) {
	//prints horizontally
	//"1 + 1 = 2" --color=green
	n := os.Args[1]
	args2 := os.Args[2]
	slice := args2[8:]

	cash := map[string]string{

		"black":   "\033[1;30m",
		"red":     "\033[1;31m",
		"green":   "\033[1;32m",
		"yellow":  "\033[1;33m",
		"blue":    "\033[1;94m",
		"purple":  "\033[38;5;56m",
		"magenta": "\033[1;35m",
		"teal":    "\033[1;36m",
		"white":   "\033[1;37m",
		"orange":  "\033[38;5;208m",
		"clear":   "\033[0m",
	}

	if !strings.Contains(slice, ":") {
		for j := 0; j < len(y[32]); j++ {
			for i := 0; i < len(n); i++ {
				if i == findN(slice) && n[i] == 32 {
					fmt.Print(cash[findC(slice)], y[int(n[i])][j])
					fmt.Print(cash[findC(slice)], y[int(n[i+1])][j])
					i++
					fmt.Print(cash["clear"])

				} else if i == findN(slice) && n[i] != 32 {
					fmt.Print(cash[findC(slice)], y[int(n[i])][j])
					fmt.Print(cash["clear"])

				} else {
					fmt.Print(cash[slice], y[int(n[i])][j])
					fmt.Print(cash["clear"])

				}
			}
			fmt.Println()

		}
	}
}