package cmd

import (
	"fmt"
	"os"
)

func dieOnError(msg string, err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s: %v", msg, err)
		os.Exit(1)
	}
}
