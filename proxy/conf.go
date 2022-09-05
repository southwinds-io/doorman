/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func init() {
	// load env vars from file if present
	godotenv.Load("proxy.env")
}

func isLoggingEnabled() bool {
	value := os.Getenv("DPROXY_LOGGING")
	return len(value) > 0
}

func getEmailFrom() (string, error) {
	value := os.Getenv("DPROXY_EMAIL_FROM")
	if len(value) == 0 {
		return "", fmt.Errorf("variable DPROXY_EMAIL_FROM is required and not defined")
	}
	return value, nil
}

func getSmtpUser() (string, error) {
	value := os.Getenv("DPROXY_SMTP_USER")
	if len(value) == 0 {
		fmt.Printf("WARNING: DPROXY_SMTP_USER not defined, using DPROXY_EMAIL_FROM instead\n")
		return getEmailFrom()
	}
	return value, nil
}

func getSmtpPwd() (string, error) {
	value := os.Getenv("DPROXY_SMTP_PWD")
	if len(value) == 0 {
		return "", fmt.Errorf("variable DPROXY_SMTP_PWD is required and not defined")
	}
	return value, nil
}

func getSmtpHost() (string, error) {
	value := os.Getenv("DPROXY_SMTP_HOST")
	if len(value) == 0 {
		return "", fmt.Errorf("variable DPROXY_SMTP_HOST is required and not defined")
	}
	return value, nil
}

func getSmtpPort() (int, error) {
	value := os.Getenv("DPROXY_SMTP_PORT")
	if len(value) == 0 {
		return -1, fmt.Errorf("variable DPROXY_SMTP_PORT is required and not defined")
	}
	port, err := strconv.Atoi(value)
	if err != nil {
		return -1, fmt.Errorf("DPROXY_SMTP_PORT value should be numeric but its value was found to be '%s'; %s", value, err)
	}
	return port, nil
}

func getDoormanBaseURI() (string, error) {
	value := os.Getenv("DPROXY_DOORMAN_URI")
	if len(value) == 0 {
		return "", fmt.Errorf("variable DPROXY_DOORMAN_URI is required and not defined")
	}
	return value, nil
}

func getDoormanUser() (string, error) {
	value := os.Getenv("DPROXY_DOORMAN_USER")
	if len(value) == 0 {
		return "", fmt.Errorf("variable DPROXY_DOORMAN_USER is required and not defined")
	}
	return value, nil
}

func getDoormanPwd() (string, error) {
	value := os.Getenv("DPROXY_DOORMAN_PWD")
	if len(value) == 0 {
		return "", fmt.Errorf("variable DPROXY_DOORMAN_PWD is required and not defined")
	}
	return value, nil
}