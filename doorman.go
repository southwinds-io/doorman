/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"southwinds.dev/doorman/core"
	h "southwinds.dev/http"
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
	Server   *h.Server
	Handlers []HandlerInfo
	Process  ProcFactory
}

func NewDoorman(pf ProcFactory) *Doorman {
	// https://textkool.com/en/ascii-art-generator?hl=default&vl=default&font=Broadway%20KB&text=dproxy%0A
	fmt.Printf(`
++++++++++++++++++++++++++++++++++++++++++++++++++++
|    ___   ___   ___   ___   _   _   __    _   _   |
|   | | \ / / \ / / \ | |_) | |\/|  / /\  | |\ |   |
|   |_|_/ \_\_/ \_\_/ |_| \ |_|  | /_/--\ |_| \|   |
|                                                  |
+++++++++|  software distribution agent  |++++++++++
%s

`, core.Version)
	return &Doorman{
		Server:   h.New("DOORMAN", core.Version),
		Handlers: Handlers(),
		Process:  pf,
	}
}

// Start the registry http server
func (d *Doorman) Start() {
	d.Server.Serve()
}

// RegisterHandlers register the http handlers for the registry
func (d *Doorman) RegisterHandlers() {
	d.Server.Http = func(router *mux.Router) {
		// enable encoded path  vars
		router.UseEncodedPath()
		// conditionally enable middleware
		if len(os.Getenv(DoormanLogging)) > 0 {
			router.Use(d.Server.LoggingMiddleware)
		}
		// apply authentication
		router.Use(d.Server.AuthenticationMiddleware)
		// register handlers
		for _, handler := range d.Handlers {
			router.HandleFunc(handler.Path, handler.Handler).Methods(handler.Methods...)
		}
	}
	// grab a reference to default auth to use it in the proxy override below
	defaultAuth = d.Server.DefaultAuth
	// set up specific authentication for doorman proxy
	d.Server.Auth = map[string]func(http.Request) *h.UserPrincipal{
		"^/token.*":  dProxyAuth,
		"^/event/.*": dProxyAuth,
	}
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

func Handlers() []HandlerInfo {
	return []HandlerInfo{
		// admin facing endpoints
		{"/key", upsertKeyHandler, []string{http.MethodPut}},
		{"/command", upsertCommandHandler, []string{http.MethodPut}},
		{"/route/in", upsertInboundRouteHandler, []string{http.MethodPut}},
		{"/route/out", upsertOutboundRouteHandler, []string{http.MethodPut}},
		{"/notification", upsertNotificationHandler, []string{http.MethodPut}},
		{"/notification", getAllNotificationsHandler, []string{http.MethodGet}},
		{"/notification-template", upsertNotificationTemplateHandler, []string{http.MethodPut}},
		{"/notification-template", getAllNotificationTemplatesHandler, []string{http.MethodPut}},
		{"/pipe", upsertPipelineHandler, []string{http.MethodPut}},
		{"/pipe/{name}", getPipelineHandler, []string{http.MethodGet}},
		{"/pipe", getAllPipelinesHandler, []string{http.MethodGet}},
		{"/job", getTopJobsHandler, []string{http.MethodGet}},
		// doorman proxy facing endpoints
		{"/event/{service-id}/{bucket-name}/{folder-name}", eventHandler, []string{http.MethodPost}},
		{"/token/{token-value}", getWebhookAuthInfoHandler, []string{http.MethodGet}},
		{"/token", getWebhookAllAuthInfoHandler, []string{http.MethodGet}},
	}
}

// dProxyAuth authenticates doorman's proxy requests using either proxy specific or admin credentials
func dProxyAuth(r http.Request) *h.UserPrincipal {
	user, userErr := core.GetProxyUser()
	if userErr != nil {
		fmt.Printf("cannot authenticate proxy: %s", userErr)
		return nil
	}
	pwd, pwdErr := core.GetProxyPwd()
	if pwdErr != nil {
		fmt.Printf("cannot authenticate proxy: %s", pwdErr)
		return nil
	}
	// try proxy specific credentials first
	if r.Header.Get("Authorization") == h.BasicToken(user, pwd) {
		return &h.UserPrincipal{
			Username: user,
			Created:  time.Now(),
		}
	} else if defaultAuth != nil {
		// try admin credentials
		if principal := defaultAuth(r); principal != nil {
			return principal
		}
	}
	// otherwise, fail authentication
	return nil
}
