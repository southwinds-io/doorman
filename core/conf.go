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
	"os"
	"strconv"
	"time"
)

const (
	// notification service settings
	NotificationURI  = "DOORMAN_NOTIFICATION_URI"
	NotificationUser = "DOORMAN_NOTIFICATION_USER"
	NotificationPwd  = "DOORMAN_NOTIFICATION_PWD"

	// artisan registry settings
	ArtRegUser = "ART_REG_USER"
	ArtRegPwd  = "ART_REG_PWD"

	// proxy settings
	ProxySourceUri          = "DOORMAN_PROXY_SRC_URI"
	ProxySourceUser         = "DOORMAN_PROXY_SRC_USER"
	ProxySourcePwd          = "DOORMAN_PROXY_SRC_PWD"
	ProxySourceInsecureSkip = "DOORMAN_PROXY_SRC_INSECURE_SKIP"
	ProxyPollIntervalSecs   = "DOORMAN_PROXY_POLL_INTERVAL_SECS"

	// pipeline configuration settings
	ConfigSourceUri          = "DOORMAN_CFG_SRC_URI"
	ConfigSourceUser         = "DOORMAN_CFG_SRC_USER"
	ConfigSourcePwd          = "DOORMAN_CFG_SRC_PWD"
	ConfigSourceInsecureSkip = "DOORMAN_CFG_SRC_INSECURE_SKIP"

	// S3 service settings for
	LogsS3URI  = "DOORMAN_LOGS_S3_URI"
	LogsS3User = "DOORMAN_LOGS_S3_USER"
	LogsS3Pwd  = "DOORMAN_LOGS_S3_PWD"
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
	value, _ := strconv.Atoi(os.Getenv(ProxyPollIntervalSecs))
	if value == 0 {
		value = 60
	}
	return time.Duration(int64(value)) * time.Second
}

func GetS3URI() string {
	return os.Getenv(LogsS3URI)
}

func GetS3User() (string, error) {
	return getString(LogsS3User)
}

func GetS3Pwd() (string, error) {
	return getString(LogsS3Pwd)
}

func GetProxyURI() (string, error) {
	return getString(ProxySourceUri)
}

func GetProxyUser() (string, error) {
	return getString(ProxySourceUser)
}

func GetProxyPwd() (string, error) {
	return getString(ProxySourcePwd)
}

func GetProxyInsecureSkip() (bool, error) {
	return getBoolean(ProxySourceInsecureSkip)
}

func GetCfgURI() (string, error) {
	return getString(ConfigSourceUri)
}

func GetCfgUser() (string, error) {
	return getString(ConfigSourceUser)
}

func GetCfgPwd() (string, error) {
	return getString(ConfigSourcePwd)
}

func GetCfgInsecureSkip() (bool, error) {
	return getBoolean(ConfigSourceInsecureSkip)
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

func GetArRegUser() (string, error) {
	return getString(ArtRegUser)
}

func GetArRegPwd() (string, error) {
	return getString(ArtRegPwd)
}
