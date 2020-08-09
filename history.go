package main

import (
	"os"

	"github.com/landoop/tableprinter"
)

func GenerateHistory() error {

	e := GetExecutions()

	tableprinter.Print(os.Stdout, e)

	return nil
}
