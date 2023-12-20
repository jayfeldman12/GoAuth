package scripts

import (
	"fmt"
	"log"
	"os/exec"
)

func Execute(command string) {
	fmt.Println("Executing command:", command)
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %s\n%s", err, output)
	}
	fmt.Println(string(output))
}
