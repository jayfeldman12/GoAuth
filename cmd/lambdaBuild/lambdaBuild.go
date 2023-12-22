package main

import (
	"GoAuth/cmd/common"
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
		common.Execute("mkdir -p dist")
		_, err = os.Stat(fmt.Sprintf("dist/%s.zip", cleanFilename))

		if err == nil {
			// File exists, delete it
			common.Execute(fmt.Sprintf("rm dist/%s.zip", cleanFilename))
		}
		common.Execute(fmt.Sprintf("GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ./dist/bootstrap %s", filename))
		common.Execute(fmt.Sprintf("cd dist && zip %s.zip bootstrap && rm bootstrap", cleanFilename))
		common.Execute(fmt.Sprintf("aws lambda update-function-code --function-name %s --zip-file fileb://dist/%s.zip", cleanFilename, cleanFilename))
	}
}
