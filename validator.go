package main

import "os"

func checkConfigExist() error {

	c, err := LoadConfig()
	if err != nil {

		return err
	}

	host := c.General.Ansiblerootdir + c.General.Hosts
	Install := c.General.Ansiblerootdir + c.General.Install
	uninstall := c.General.Ansiblerootdir + c.General.Uninstall
	update := c.General.Ansiblerootdir + c.General.Update

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
