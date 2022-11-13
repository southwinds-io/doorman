/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	"fmt"
	"southwinds.dev/artisan/core"
	"southwinds.dev/doorman/types"
	"southwinds.dev/os"
	"southwinds.dev/types/doorman"
	"strings"
	"time"
)

func LogJob(pipeline *doorman.Pipeline, process *Process, jobNo string, started *time.Time, status string) error {
	completed := time.Now().UTC()
	job := &types.Job{
		Number:    jobNo,
		ServiceId: process.serviceId,
		Bucket:    process.bucketName,
		Folder:    process.folderName,
		Pipeline:  pipeline,
		Status:    status,
		Started:   started,
		Completed: &completed,
	}
	uri := GetS3URI()
	// if an S3 URI has been defined
	if len(uri) > 0 && strings.Contains(uri, "s3") {
		user, err := GetS3User()
		if err != nil {
			return err
		}
		pwd, err := GetS3Pwd()
		if err != nil {
			return err
		}
		// write the job data to S3
		return os.WriteFile(
			job.Bytes(),
			fmt.Sprintf("%s/%s", uri, time.Now().Format(time.RFC3339)),
			fmt.Sprintf("%s:%s", user, pwd))
	} else {
		// otherwise write to stdout
		core.InfoLogger.Printf("%s\n", job)
	}
	return nil
}
