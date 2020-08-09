package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Application application
	General     general
	Api         api
}

type application struct {
	Name    string
	Version string
}

type api struct {
	Url       string
	Enableapi bool
}

type general struct {
	Ansiblerootdir string
	Hosts          string
	Install        string
	Update         string
	Uninstall      string
}

func LoadConfig() (Config, error) {

	c := "conf/app.ini"
	var conf Config

	if _, err := os.Stat(c); os.IsNotExist(err) {
		fmt.Printf("Config " + c + " does not exist")
		return conf, err
	}

	if _, err := toml.DecodeFile(c, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}

func runCommand(fla Flag) error {

	u, err := user.Current()

	if err != nil {
		return err
	}

	c, err := LoadConfig()
	if err != nil {
		return err
	}

	ansibleplaybook, interprete, err := checkBinary()
	if err != nil {
		return err
	}

	var ansiblecommand string

	InfoLog.Println("Start " + fla.Type + "")
	InfoLog.Println("Started by: " + u.Name + "")
	InfoLog.Println("Application: " + c.Application.Name + "")
	InfoLog.Println("Version: " + c.Application.Version + "")
	fmt.Println("Start " + fla.Type + "")
	fmt.Println("Started by: " + u.Name + "")
	fmt.Println("Application: " + c.Application.Name + "")
	fmt.Println("Version: " + c.Application.Version + "")

	if fla.Type == "install" {
		ansiblecommand = c.General.Ansiblerootdir + c.General.Hosts + " " + c.General.Ansiblerootdir + c.General.Install
	}
	if fla.Type == "update" {
		ansiblecommand = c.General.Ansiblerootdir + c.General.Hosts + " " + c.General.Ansiblerootdir + c.General.Update
	}
	if fla.Type == "uninstall" {
		ansiblecommand = c.General.Ansiblerootdir + c.General.Hosts + " " + c.General.Ansiblerootdir + c.General.Uninstall
	}

	cmd := exec.Command(interprete, "-c", ""+ansibleplaybook+" -i "+ansiblecommand+"")

	// Crear pipe para capturar commando en vivo

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\t> %s\n", scanner.Text())
			ExecutionLog.Printf("\t%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		ErrorLog.Printf(err.Error())
		var e Execution
		e.Application = c.Application.Name
		e.Version = c.Application.Version
		e.Status = false
		e.TypeExec = fla.Type
		e.User = u.Name
		e.Date = time.Now().Format("2006-01-02 15:06:23")

		if c.Api.Enableapi {

			err = CallApiExecutions(e, c)
			if err != nil {
				return err
			}

		}

		err = SaveExecution(e)
		if err != nil {
			return err
		}

		return err
	}

	InfoLog.Println(fla.Type + "finish successful!")
	fmt.Println(fla.Type + " finish successful!")
	var e Execution
	e.Application = c.Application.Name
	e.Version = c.Application.Version
	e.Status = true
	e.TypeExec = fla.Type
	e.User = u.Name
	e.Date = time.Now().Format("2006-01-02 15:06:23")

	if c.Api.Enableapi {
		fmt.Printf("ENTREEEEE API")
		err = CallApiExecutions(e, c)
		if err != nil {
			return err
		}

	}

	err = SaveExecution(e)
	if err != nil {
		fmt.Printf("ENTREEEEE ERROR SAVE")
		return err
	}

	return nil

}

func checkBinary() (string, string, error) {

	var ansibleplaybook string
	var interprete string

	if _, err := os.Stat("/bin/bash"); os.IsNotExist(err) {
		if _, err := os.Stat("/bin/sh"); os.IsNotExist(err) {
			return ansibleplaybook, interprete, err
		} else {
			interprete = "/bin/sh"
		}
	} else {
		interprete = "/bin/bash"
	}

	if _, err := os.Stat("/bin/ansible-playbook"); os.IsNotExist(err) {
		if _, err := os.Stat("/usr/bin/ansible-playbook"); os.IsNotExist(err) {
			return ansibleplaybook, interprete, err
		} else {
			ansibleplaybook = "/usr/bin/ansible-playbook"
		}

	} else {
		ansibleplaybook = "/bin/ansible-playbook"
	}

	return ansibleplaybook, interprete, nil
}
