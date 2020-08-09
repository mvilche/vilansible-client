package main

import (
	"errors"
	"flag"
	"strings"
)

type Flag struct {
	Type    string
	History bool
}

func getFlag() (Flag, error) {
	var f Flag
	passed := false
	var err error
	runtype := flag.String("type", "install, update, uninstall", "A string")
	h := flag.Bool("history", false, "bool")
	flag.Parse()

	f.Type = strings.TrimSpace(*runtype)
	f.History = *h
	if f.Type == "install" && !*h {
		passed = true
	}
	if f.Type == "update" && !*h {
		passed = true
	}
	if f.Type == "uninstall" && !*h {
		passed = true
	}

	if f.History {
		passed = true
	}

	if !passed {
		err = errors.New("Value for -type is install, update or uninstall")

	}

	return f, err
}
