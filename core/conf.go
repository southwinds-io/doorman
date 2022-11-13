/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	NotificationURI          = "DOORMAN_NOTIFICATION_URI"
	NotificationUser         = "DOORMAN_NOTIFICATION_USER"
	NotificationPwd          = "DOORMAN_NOTIFICATION_PWD"
	PollIntervalSecs         = "DOORMAN_POLL_INTERVAL_SECS"
	OxWapiUri                = "OX_WAPI_URI"
	OxWapiUser               = "OX_WAPI_USER"
	OxWapiPwd                = "OX_WAPI_PWD"
	OxWapiInsecureSkipVerify = "OX_WAPI_INSECURE_SKIP_VERIFY"
	ArtRegUser               = "ART_REG_USER"
	ArtRegPwd                = "ART_REG_PWD"
	SourceHost               = "DOORMAN_SRC_HOST"
	SourceUser               = "DOORMAN_SRC_USER"
	SourcePwd                = "DOORMAN_SRC_PWD"
	S3URI                    = "DOORMAN_S3_URI"
	S3User                   = "DOORMAN_S3_USER"
	S3Pwd                    = "DOORMAN_S3_PWD"
)

func init() {
	// load env vars from file if present
	godotenv.Load("doorman.env")
}

func getString(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return "", fmt.Errorf("variable %s is required and not defined", key)
	}
	return value, nil
}

func getBoolean(key string) (bool, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return false, fmt.Errorf("variable %s is required and not defined", key)
	}
	return strconv.ParseBool(value)
}

func GetPollInterval() time.Duration {
	value, _ := strconv.Atoi(os.Getenv(PollIntervalSecs))
	if value == 0 {
		value = 60
	}
	return time.Duration(int64(value)) * time.Second
}

func GetS3URI() string {
	return os.Getenv(S3URI)
}

func GetS3User() (string, error) {
	return getString(S3User)
}

func GetS3Pwd() (string, error) {
	return getString(S3Pwd)
}

func GetSourceHost() string {
	value, err := getString(SourceHost)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func GetSourceUser() string {
	value, err := getString(SourceUser)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func GetSourcePwd() string {
	value, err := getString(SourcePwd)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func GetNotificationURI() (string, error) {
	return getString(NotificationURI)
}

func GetNotificationUser() (string, error) {
	return getString(NotificationUser)
}

func GetNotificationPwd() (string, error) {
	return getString(NotificationPwd)
}

func GetOxWapiUri() (string, error) {
	return getString(OxWapiUri)
}

func GetOxWapiUser() (string, error) {
	return getString(OxWapiUser)
}

func GetOxWapiPwd() (string, error) {
	return getString(OxWapiPwd)
}

func GetOxWapiInsecureSkipVerify() (bool, error) {
	return getBoolean(OxWapiInsecureSkipVerify)
}

func GetArRegUser() (string, error) {
	return getString(ArtRegUser)
}

func GetArRegPwd() (string, error) {
	return getString(ArtRegPwd)
}
