/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

type ProcFactory interface {
	New(serviceId, bucketPath, folderName, artHome string) (Processor, error)
}

func NewDefaultProcFactory() ProcFactory {
	return new(DefaultProcFactory)
}

type DefaultProcFactory struct {
}

func (df *DefaultProcFactory) New(serviceId, bucketPath, folderName, artHome string) (Processor, error) {
	return NewProcess(serviceId, bucketPath, folderName, artHome, nil, nil)
}
