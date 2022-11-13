/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package types

import (
	"encoding/json"
	"fmt"
	"southwinds.dev/types/doorman"
	"time"
)

type Job struct {
	Number    string            `bson:"_id" json:"number" yaml:"number"`
	ServiceId string            `bson:"service_id" json:"service_id" yaml:"service_id"`
	Bucket    string            `bson:"bucket" json:"bucket" yaml:"bucket"`
	Folder    string            `bson:"folder" json:"folder" yaml:"folder"`
	Pipeline  *doorman.Pipeline `bson:"pipeline" json:"pipeline" yaml:"pipeline"`
	Status    string            `bson:"status" json:"status" yaml:"status"`
	Log       []string          `bson:"log" json:"log" yaml:"log"`
	Started   *time.Time        `bson:"started" json:"started" yaml:"started"`
	Completed *time.Time        `bson:"completed" json:"completed" yaml:"completed"`
}

func (j *Job) GetName() string {
	return j.Number
}

func (j *Job) String() string {
	return string(j.Bytes()[:])
}

func (j *Job) Bytes() []byte {
	jBytes, err := json.MarshalIndent(j, "", " ")
	if err != nil {
		return []byte(fmt.Sprintf("cannot marshal job data: %s", err))
	}
	return jBytes
}
