/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package main

import (
	"log"
	"path/filepath"
	artCore "southwinds.dev/artisan/core"
	"southwinds.dev/doorman/core"
)

var D *Doorman

func main() {
	var err error
	if err := checkDoormanHome(); err != nil {
		log.Fatalf("cannot launch  doorman, cannot write to file system: %s", err)
	}
	D, err = NewDoorman(core.NewDefaultProcFactory())
	if err != nil {
		log.Fatalf(err.Error())
	}
	D.Start()
}

func checkDoormanHome() error {
	path := filepath.Join(artCore.HomeDir(), ".doorman")
	return artCore.EnsureRegistryPath(path)
}
