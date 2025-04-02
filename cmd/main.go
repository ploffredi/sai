package main

import (
	"fmt"
	"os"
	"regexp"
	"sai/pkg/actions"
//	"sai/pkg/providers"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: sai <software> <action> [provider]")
		os.Exit(1)
	}

	software := os.Args[1]
	re := regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)
	if !re.MatchString(software) {
		fmt.Println("Invalid software name provided.")
		os.Exit(1)
	}
	action := os.Args[2]
	provider := ""
	if len(os.Args) > 3 {
		provider = os.Args[3]
	}

		if actions.IsActionSupported(provider, action) {
			err := actions.PerformAction(software, provider, action)
			if err != nil {
				fmt.Printf("TODO: sai %s %s %s\n%v\n",software, action, provider, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Warning: Action '%s' is currently unsupported. Run sai <software> support to request support\n", action, provider)
			os.Exit(1)
		}

}
