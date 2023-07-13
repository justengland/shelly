package main

import (
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func ToYAML(filename string, obj interface{}) error {
	yamlBytes, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}

	// Get the current user's home directory
	usr, err := user.Current()
	if err != nil {
		return err
	}

	// Create the "shelly" directory if it doesn't exist
	shellyDir := filepath.Join(usr.HomeDir, "shelly")
	err = os.MkdirAll(shellyDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Set the file path for YAML serialization
	filePath := filepath.Join(shellyDir, filename)

	err = ioutil.WriteFile(filePath, yamlBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func FromYAML(filename string, obj interface{}) error {
	// Get the current user's home directory
	usr, err := user.Current()
	if err != nil {
		return err
	}

	// Set the file path for YAML deserialization
	filePath := filepath.Join(usr.HomeDir, "shelly", filename)

	yamlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlBytes, obj)
	if err != nil {
		return err
	}

	return nil
}
