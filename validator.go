package main

import "os"

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
