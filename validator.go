package main

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func checkConfigExist() error {

	c, err := LoadConfig()
	if err != nil {

		return err
	}

	host := c.Ansible.Ansiblerootdir + c.Ansible.Hosts
	Install := c.Ansible.Ansiblerootdir + c.Ansible.Install
	uninstall := c.Ansible.Ansiblerootdir + c.Ansible.Uninstall
	update := c.Ansible.Ansiblerootdir + c.Ansible.Update

	if _, err := os.Stat(host); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(Install); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(uninstall); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(update); os.IsNotExist(err) {
		return err
	}

	return nil
}

func checkSudoRun() error {

	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	if err != nil {
		return err
	}
	i, err := strconv.Atoi(string(output[:len(output)-1]))

	if err != nil {
		return err
	}
	if i != 0 {
		return errors.New("This program must be run with sudo")
	}

	return nil
}

func checkOS() error {

	if runtime.GOOS == "windows" {
		return errors.New("Windows not supported")
	}
	return nil
}
