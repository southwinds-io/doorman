/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	"errors"
	"southwinds.dev/artisan/data"
	"southwinds.dev/artisan/release"
)

const (
	Name               = "doorman"
	ArtisanCommandType = "ARTISAN-COMMAND"
	CatalogueItemType  = "CATALOGUE-ITEM"
)

var (
	ErrDocumentAlreadyExists = errors.New("mongo: the document already exists")
	ErrDocumentNotFound      = errors.New("mongo: the document was not found")
)

type CatalogueItem struct {
	Name       string
	Spec       *release.Spec
	Tags       []string
	Attributes map[string]interface{}
}

func (i *CatalogueItem) Validate() error {
	return nil
}

type Command struct {
	Name        string
	Description string
	Package     string
	Function    string
	RegUser     string
	RegPwd      string
	Input       *data.Input
	Tag         []string
}

func (i *Command) Validate() error {
	return nil
}
