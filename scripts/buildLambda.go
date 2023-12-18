package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func execute(command string) {
	fmt.Println("Executing command:", command)
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error: %q", err)
	}
	fmt.Println(string(output))
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("Filename is empty, please provide a filename as the first arg")
	} else {
		_, err := os.Stat(filename)
		if err != nil {
			log.Fatalf("Invalid file: %q", filename)
		}

		cleanFilename := strings.TrimSuffix(path.Base(filename), ".go")
		fmt.Printf("Using file %s\n", cleanFilename)
		execute("mkdir -p dist")
		_, err = os.Stat(fmt.Sprintf("dist/%q.zip", cleanFilename))
		if err != nil {
			// File exists, delete it
			execute(fmt.Sprintf("rm dist/%q.zip", cleanFilename))
		}
		execute(fmt.Sprintf("GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ./dist/bootstrap %q", filename))
		execute(fmt.Sprintf("cd dist && zip %q.zip bootstrap && rm bootstrap", cleanFilename))
		execute(fmt.Sprintf("aws lambda update-function-code --function-name %q --zip-file fileb://dist/%q.zip", cleanFilename, cleanFilename))
	}
}
