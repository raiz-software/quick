package main

import (
	"os"
	"fmt"
	"flag"
	"os/exec"
	"io/ioutil"
	"encoding/json"
)

func main() {

	runcmd := flag.NewFlagSet("run", flag.ExitOnError)

	N := len(os.Args)

	if N > 1 {

		switch os.Args[1] {
		case "run":
			if N > 2 {
				command := os.Args[N-1]
				runcmd.Parse(os.Args[2:N-1])
				run(command)
			} else {
				fmt.Println("Are you sure? I can't read any command")
				fmt.Println("Try 'quick run something'")
			}
		default:
			fmt.Printf("Well, we can't do '%s'\n", os.Args[1])
			fmt.Println("Try 'quick run something'")
		}

	} else {

		fmt.Println("What are you doing?")
		fmt.Println("Try 'quick run something'")
	}
}

func run(command string) {

	quickFile, err := os.Open("quick.json")
	if err != nil {
		fmt.Println(err)
	}

	defer quickFile.Close()

	quickBytes, err := ioutil.ReadAll(quickFile)
	if err != nil {
		fmt.Println(err)
	}

	var quick map[string]interface{}
	json.Unmarshal([]byte(quickBytes), &quick)

	if value, ok := quick[command]; ok {

		cmd := exec.Command("bash", "-c", value.(string))

		cmd.Stdin = os.Stdin;
		cmd.Stdout = os.Stdout;
		cmd.Stderr = os.Stderr;

		fmt.Printf("> %s: %s\n", command, value)

		err := cmd.Run()

		if err != nil {
    		return
		}

	} else {

		fmt.Printf("We couldn't run '%s'\n", command)
		fmt.Println("Are you sure this is defined in quick.json?")
	}
}