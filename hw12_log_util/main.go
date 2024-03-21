package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	file   string
	level  string
	output string
)

func init() {
	if flag.Lookup("file") == nil {
		flag.StringVar(&file, "file",
			defaultOrEnv("LOG_ANALYZER_FILE", ""), "The path to the log file.")
	}
	if flag.Lookup("level") == nil {
		flag.StringVar(&level, "level",
			defaultOrEnv("LOG_ANALYZER_LEVEL", "info"), "Log level. Supported: info, warning, error.")
	}
	if flag.Lookup("output") == nil {
		flag.StringVar(&output, "output",
			defaultOrEnv("LOG_ANALYZER_OUTPUT", ""), "The path to the output file.")
	}
}

func main() {
	var counter map[string]int
	var err error

	flag.Parse()
	if file == "" {
		log.Fatal("the '-file' argument is not specified.")
	}
	fmt.Println("Log file analysis started...")

	counter, err = analysis(file, level)
	if err != nil {
		log.Fatalf("An error occurred while analyzing the log file: %v", err)
	}

	statistics := prepareOutput(counter, file)

	if output == "" {
		fmt.Print(statistics)
		fmt.Println("\nDone")
	} else {
		var file *os.File

		file, err = os.Create(output)
		if err != nil {
			log.Fatal("Unable to create file:", err)
		}
		defer file.Close()
		file.WriteString(statistics)
		fmt.Printf("Statistics have been written to a file %s\n", output)
	}
}

func defaultOrEnv(env string, def string) string {
	value, ok := os.LookupEnv(env)
	if ok {
		return value
	}
	return def
}

func analysis(path string, level string) (map[string]int, error) {
	var file *os.File
	var err error

	counter := make(map[string]int)
	levels := coverage(level)
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, lvl := range levels {
			if strings.HasPrefix(line, lvl) {
				counter[lvl]++
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return counter, nil
}

func coverage(level string) []string {
	var list []string
	list = append(list, level)
	if level == "info" {
		list = append(list, "warning", "error")
	}
	if level == "warning" {
		list = append(list, "error")
	}
	return list
}

func prepareOutput(counter map[string]int, filename string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Statistics of the %s\n\n", filename))
	builder.WriteString("|  Level   |  Count   |\n")
	builder.WriteString("| -------- | -------- |\n")
	for k, v := range counter {
		builder.WriteString(fmt.Sprintf("| %-8v | %-8v |\n", k, v))
	}
	return builder.String()
}
