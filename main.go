package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

const currentVersion string = "1.0"
const packageName string = "extract-hostname"

var (
	help    = flag.Bool("help", false, "Show help screen.")
	version = flag.Bool("version", false, "Display version info.")
)

// const (
// 	resultsError            = "ERROR: %s '%s'\n"
// )

func main() {
	flag.Parse()

	if *help {
		printUsage()
		return
	}

	if *version {
		fmt.Println(packageName + " version " + currentVersion)
		return
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: cat file | " + packageName)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		hostName := extractHostName(input)

		if hostName != "" {
			fmt.Printf("%s\n", hostName)
		}
	}
}

func printUsage() {
	fmt.Print(packageName + " [-flag value]\n\n")
	flag.Usage()
	fmt.Println()
}

func extractHostName(str string) string {
	var hostName = ""

	re := regexp.MustCompile(`(([a-zA-Z\d-]+\.){0,}([a-zA-Z\d-]+){0,}\.([a-zA-Z\d-]+){2,})`)
	if re.MatchString(str) {
		hostName = re.FindString(str)
	}

	return hostName
}
