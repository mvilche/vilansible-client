package main

import (
	"fmt"
	"os"
)

func main() {

	err := checkConfigExist()
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

	err = runCommand(fla)
	if err != nil {
		fmt.Println("Error ", err.Error())
		ErrorLog.Printf(err.Error())
		os.Exit(1)

	}

}
