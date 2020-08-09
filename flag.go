package main

import (
	"errors"
	"flag"
	"strings"
)

type Flag struct {
	Type string
}

func getFlag() (Flag, error) {
	var f Flag
	passed := false
	var err error
	runtype := flag.String("type", "install, update, uninstall", "A string")
	flag.Parse()
	f.Type = strings.TrimSpace(*runtype)

	if f.Type == "install" {
		passed = true
	}
	if f.Type == "update" {
		passed = true
	}
	if f.Type == "uninstall" {
		passed = true
	}

	if !passed {
		err = errors.New("Value for -type is install, update or uninstall")

	}

	return f, err
}
