/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"math"
	"net/http"
	"path/filepath"
	c "southwinds.dev/artisan/core"
	"southwinds.dev/artisan/release"
	"southwinds.dev/d-proxy/types"
	"southwinds.dev/doorman/core"
	h "southwinds.dev/http"
	src "southwinds.dev/source_client"
	"southwinds.dev/types/doorman"
	"southwinds.dev/types/dproxy"
	"strings"
	"time"
)

const DoormanLogging = "DOORMAN_LOGGING"

var defaultAuth func(r http.Request) *h.UserPrincipal

type HandlerInfo struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
	Methods []string
}

type Doorman struct {
	Process ProcFactory
	proxy   *src.Client
}

func NewDoorman(pf ProcFactory) (*Doorman, error) {
	// https://textkool.com/en/ascii-art-generator?hl=default&vl=default&font=Broadway%20KB&text=dproxy%0A
	fmt.Printf(`
++++++++++++++++++++++++++++++++++++++++++++++++++++
|    ___   ___   ___   ___   _   _   __    _   _   |
|   | | \ / / \ / / \ | |_) | |\/|  / /\  | |\ |   |
|   |_|_/ \_\_/ \_\_/ |_| \ |_|  | /_/--\ |_| \|   |
|                                                  |
+++++++++| automated package distribution |+++++++++
%s

`, core.Version)
	p, err := getProxyClient()
	if err != nil {
		return nil, err
	}
	return &Doorman{
		Process: pf,
		proxy:   p,
	}, nil
}

func (d *Doorman) Start() {
	interval := core.GetPollInterval()
	if err := d.initConfig(); err != nil {
		log.Fatalf(err.Error())
	}
	var (
		release    *types.Release
		anyRelease interface{}
		err        error
	)
	for {
		anyRelease, err = d.proxy.PopOldest(dproxy.ReleaseType, new(types.Release))
		if err != nil {
			log.Fatalf(err.Error())
		}
		// if no release was found
		if anyRelease == nil {
			c.Debug("no release found, retrying later in %v...", interval)
			// wait for a while
			time.Sleep(core.GetPollInterval())
			// then try again
			continue
		} else {
			release = anyRelease.(*types.Release)
		}
		// creates a dedicated, randomly named artisan local registry home
		artHome, artHomeErr := newArtHome()
		if artHomeErr != nil {
			log.Fatalf("cannot create artisan home: %s, cannot continue", err)
		}
		// creates a new process running using the dedicated registry
		proc, e := D.Process.New(release.DeploymentId, release.BucketName, release.FolderName, artHome)
		if e != nil {
			c.ErrorLogger.Printf("cannot create pipeline processor: %s", e)
		}
		// starts the process asynchronously
		proc.Start()
	}
}

func (d *Doorman) initConfig() error {
	// pipelines
	if err := d.proxy.SetType(doorman.PipelineType, doorman.PipelineConf{
		Name:           "sample-pipeline",
		InboundRoutes:  []string{"sample-inbound-route"},
		OutboundRoutes: []string{"sample-outbound-route"},
		Commands:       []string{"sample-command"},
		CMDB: &doorman.CMDB{
			Catalogue: false,
			Events:    []string{"sample-event"},
			Tag:       []string{"tag-1", "tag-2"},
		},
		SuccessNotification:   "success-notification",
		ErrorNotification:     "error-notification",
		CmdFailedNotification: "cmd-failed-notification",
	}); err != nil {
		return fmt.Errorf("cannot set pipeline type in source: %s", err)
	}
	if err := d.proxy.SetType(doorman.InboundRouteType, doorman.InRoute{
		Name:             "",
		Description:      "",
		ServiceHost:      "",
		ServiceId:        "",
		BucketName:       "",
		User:             "",
		Pwd:              "",
		Verify:           false,
		WebhookToken:     "",
		WebhookWhitelist: nil,
		Filter:           "",
	}); err != nil {
		return fmt.Errorf("cannot set inbound route type in source: %s", err)
	}
	if err := d.proxy.SetType(doorman.OutboundRouteType, doorman.OutRoute{
		Name:        "SAMPLE-OUT-ROUTE",
		Description: "this is a sample outbound route",
		PackageRegistry: &doorman.PackageRegistry{
			Domain: "packages.sample.com",
			Group:  "sample-group",
			User:   "sample-user",
			Pwd:    "sample-pwd",
		},
		ImageRegistry: &doorman.ImageRegistry{
			Domain: "images.sample.com",
			Group:  "sample-group",
			User:   "sample-user",
			Pwd:    "sample-pwd",
		},
		S3Store: &doorman.S3Store{
			BucketURI: "s3.sample.com/sample-bucket",
			User:      "sample-user",
			Pwd:       "sample-pwd",
			Partition: "",
			Service:   "",
			Region:    "",
			AccountID: "",
			Resource:  "",
		},
	}); err != nil {
		return fmt.Errorf("cannot set outbound route type in source: %s", err)
	}
	if err := d.proxy.SetType(doorman.CommandType, doorman.Command{
		Name:        "SAMPLE-COMMAND",
		Description: "This is a command run by Doorman",
		Value:       "scan files",
		ErrorRegex:  "",
		StopOnError: false,
	}); err != nil {
		return fmt.Errorf("cannot set command type in source: %s", err)
	}
	if err := d.proxy.SetType(doorman.NotificationTemplateType, doorman.NotificationTemplate{
		Name:    "SAMPLE-SUCCESS-TEMPLATE",
		Subject: "Hello there",
		Content: "everything has gone ok",
	}); err != nil {
		return fmt.Errorf("cannot set notification template type in source: %s", err)
	}
	if err := d.proxy.SetType(doorman.NotificationType, doorman.Notification{
		Name:      "SUCCESSFUL_RELEASE_NOTIFICATION",
		Recipient: "test@email.com",
		Type:      "email",
		Template:  "NEW_RELEASE_TEMPLATE",
	}); err != nil {
		return fmt.Errorf("cannot set notification type in source: %s", err)
	}
	// catalogue items
	if err := d.proxy.SetType(core.CatalogueItemType, core.CatalogueItem{
		Name: "",
		Spec: &release.Spec{
			Name:        "",
			Description: "",
			Author:      "",
			License:     "",
			Version:     "",
			Info:        "",
			Images:      nil,
			Packages:    nil,
			OsPackages:  nil,
			Run:         nil,
		},
		Tags:       nil,
		Attributes: nil,
	}); err != nil {
		return fmt.Errorf("cannot set catalogue item type in source: %s", err)
	}
	// commands
	if err := d.proxy.SetType(core.ArtisanCommandType, core.Command{
		Name:        "",
		Description: "",
		Package:     "",
		Function:    "",
		RegUser:     "",
		RegPwd:      "",
		Input:       nil,
		Tag:         nil,
	}); err != nil {
		return fmt.Errorf("cannot set command type in source: %s", err)
	}
	return nil
}

// backoffTime exponentially increase backoff time until reaching 1 hour
func backoffTime(attempts int) time.Duration {
	var exponentialBackoffCeilingSecs int64 = 3600 // 1 hour
	delaySecs := int64(math.Floor((math.Pow(2, float64(attempts)) - 1) * 0.5))
	if delaySecs > exponentialBackoffCeilingSecs {
		delaySecs = exponentialBackoffCeilingSecs
	}
	return time.Duration(delaySecs) * time.Second
}

type ProcFactory interface {
	New(serviceId, bucketPath, folderName, artHome string) (core.Processor, error)
}

func NewDefaultProcFactory() ProcFactory {
	return new(DefaultProcFactory)
}

type DefaultProcFactory struct {
}

func (df *DefaultProcFactory) New(serviceId, bucketPath, folderName, artHome string) (core.Processor, error) {
	return core.NewProcess(serviceId, bucketPath, folderName, artHome)
}

// newArtHome generates a new random path for the artisan home
func newArtHome() (string, error) {
	path := filepath.Join(c.HomeDir(), ".doorman", strings.Replace(uuid.NewString(), "-", "", -1)[:12])
	err := c.EnsureRegistryPath(path)
	if err != nil {
		return "", err
	}
	c.Debug("the local registry home is: '%s'\n", path)
	return path, nil
}

func getProxyClient() (*src.Client, error) {
	uri, err := core.GetProxyURI()
	if err != nil {
		return nil, err
	}
	user, err := core.GetProxyUser()
	if err != nil {
		return nil, err
	}
	pwd, err := core.GetProxyPwd()
	if err != nil {
		return nil, err
	}
	insecureSkip, err := core.GetProxyInsecureSkip()
	if err != nil {
		return nil, err
	}
	s := src.New(uri, user, pwd, &src.ClientOptions{
		InsecureSkipVerify: insecureSkip,
		Timeout:            60 * time.Second,
	})
	s.Logger = new(core.RetryLogger)
	return s, nil
}
