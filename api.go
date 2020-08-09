package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func CallApiExecutions(e Execution, c Config) error {
	var s string

	if e.Status {
		s = "true"
	} else {
		s = "false"
	}

	requestBody, err := json.Marshal(map[string]string{

		"application": e.Application,
		"version":     e.Version,
		"typexec":     e.TypeExec,
		"date":        e.Date,
		"status":      s,
		"user":        e.User,
	})

	if err != nil {

		return err
	}

	resp, err := http.Post(c.Api.Url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {

		return errors.New("Status code is " + string(resp.StatusCode) + "")

	}

	return nil

}
