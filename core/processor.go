/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	"southwinds.dev/types/doorman"
)

type Processor interface {
	// Start the processing of an event
	Start()

	// Pipeline executes a pipeline
	Pipeline(pipe *doorman.Pipeline) error

	// InboundRoute executes an inbound route in a pipeline
	InboundRoute(pipe *doorman.Pipeline, route *doorman.InRoute) error

	// PreImport runs pre import checks
	PreImport(route *doorman.InRoute, err error) error

	// Command executes the specified command
	Command(command *doorman.Command) error

	// OutboundRoute execute the outbound route
	OutboundRoute(outRoute *doorman.OutRoute) error

	// ExportFiles from a specification to an S3 bucket
	ExportFiles(s3Store *doorman.S3Store) error

	// PushImages in a specification to a container registry
	PushImages(imageRegistry *doorman.ImageRegistry) error

	// PushPackages in a specification to an artisan registry
	PushPackages(pkgRegistry *doorman.PackageRegistry) error

	// ImportFiles from a specification
	ImportFiles() error

	// SendNotification to the notification service
	SendNotification(nType NotificationType) error

	// BeforeComplete executes tasks at the end of the pipeline process
	BeforeComplete(pipe *doorman.Pipeline) error

	// Info logger
	Info(format string, a ...interface{})

	// Error logger
	Error(format string, a ...interface{}) error

	// Warn logger
	Warn(format string, a ...interface{})
}
