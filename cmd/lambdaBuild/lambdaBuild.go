package main

import (
	"GoAuth/scripts"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("Filename is empty, please provide a filename as the first arg")
	} else {
		_, err := os.Stat(filename)
		if err != nil {
			log.Fatalf("Invalid file: %s", filename)
		}

		cleanFilename := strings.TrimSuffix(path.Base(filename), "Cmd.go")
		fmt.Printf("Using file %s\n", cleanFilename)
		scripts.Execute("mkdir -p dist")
		_, err = os.Stat(fmt.Sprintf("dist/%s.zip", cleanFilename))

		if err == nil {
			// File exists, delete it
			scripts.Execute(fmt.Sprintf("rm dist/%s.zip", cleanFilename))
		}
		scripts.Execute(fmt.Sprintf("GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ./dist/bootstrap %s", filename))
		scripts.Execute(fmt.Sprintf("cd dist && zip %s.zip bootstrap && rm bootstrap", cleanFilename))
		scripts.Execute(fmt.Sprintf("aws lambda update-function-code --function-name %s --zip-file fileb://dist/%s.zip", cleanFilename, cleanFilename))
	}
}
