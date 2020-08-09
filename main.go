package main

import (
	"fmt"
	"os"
)

func main() {

	err := InitdB()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		ErrorLog.Printf(err.Error())
		os.Exit(1)
	}

	err = checkConfigExist()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		ErrorLog.Printf(err.Error())
		os.Exit(1)
	}
	_, _, err = checkBinary()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		ErrorLog.Printf(err.Error())
		os.Exit(1)
	}

	err = createLogFile()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		ErrorLog.Printf(err.Error())
		os.Exit(1)
	}

	fla, err := getFlag()
	if err != nil {
		fmt.Println("Error ", err.Error())
		ErrorLog.Printf(err.Error())
		os.Exit(1)
	}

	if fla.History {
		err := GenerateHistory()
		if err != nil {
			ErrorLog.Printf(err.Error())

		}
		os.Exit(0)
	}

	err = runCommand(fla)
	if err != nil {
		fmt.Println("Error ", err.Error())
		ErrorLog.Printf(err.Error())
		os.Exit(1)

	}

}
