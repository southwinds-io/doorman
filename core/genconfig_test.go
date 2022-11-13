/*
Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
Contributors to this project, hereby assign copyright in this code to the project,
to be licensed under the same terms as the rest of the code.
*/
package core

import (
	"github.com/joho/godotenv"
	src "southwinds.dev/source_client"
	d "southwinds.dev/types/doorman"
	"testing"
)

func TestGenConfig(t *testing.T) {
	godotenv.Load("../doorman.env")
	s := src.New(GetSourceHost(), GetSourceUser(), GetSourcePwd(), &src.ClientOptions{
		InsecureSkipVerify: true,
		Timeout:            0,
	})
	s.Logger = new(RetryLogger)
	var err error
	err = s.Save("TEST_IN_ROUTE", d.InboundRouteType, d.InRoute{
		Name:             "TEST_IN_ROUTE",
		Description:      "a testing inbound route",
		ServiceHost:      "s3://127.0.0.1:9000",
		ServiceId:        "415edcb3-16db-4f6d-828a-48c5bc33e01b",
		BucketName:       "*",
		User:             "admin",
		Pwd:              "password",
		Verify:           false,
		WebhookToken:     "JFkxnsn++02UilVkYFFC9w==",
		WebhookWhitelist: []string{"127.0.0.1"},
		Filter:           "*.",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("TEST_OUT_ROUTE", d.OutboundRouteType, d.OutRoute{
		Name:        "TEST_OUT_ROUTE",
		Description: "a testing outbound route",
		PackageRegistry: &d.PackageRegistry{
			Domain: "localhost:8082",
			User:   "admin",
			Pwd:    "admin",
		},
		ImageRegistry: &d.ImageRegistry{
			Domain: "localhost:5000",
			User:   "admin",
			Pwd:    "admin",
		},
		S3Store: &d.S3Store{
			BucketURI: "s3://localhost:9000/ilink2",
			User:      "admin",
			Pwd:       "password",
			Partition: "minio",
			Service:   "sqs",
			Region:    "",
			AccountID: "_",
			Resource:  "webhook",
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("SUCCESSFUL_RELEASE_NOTIFICATION", d.NotificationType, d.Notification{
		Name:      "SUCCESSFUL_RELEASE_NOTIFICATION",
		Recipient: "test@email.com",
		Type:      "email",
		Template:  "NEW_RELEASE_TEMPLATE",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("FAIL_ERROR_NOTIFICATION", d.NotificationType, d.Notification{
		Name:      "FAIL_ERROR_NOTIFICATION",
		Recipient: "test@email.com",
		Type:      "email",
		Template:  "ISSUE_TEMPLATE",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("FAIL_SCAN_NOTIFICATION", d.NotificationType, d.Notification{
		Name:      "FAIL_SCAN_NOTIFICATION",
		Recipient: "test@email.com",
		Type:      "email",
		Template:  "QUARANTINE_TEMPLATE",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("QUARANTINE_TEMPLATE", d.NotificationTemplateType, d.NotificationTemplate{
		Name:    "QUARANTINE_TEMPLATE",
		Subject: "New Release <<release-name>> has been quarantined",
		Content: `A new release has been published and but has been quarantined.
   Error log:
   <<command-log>>`,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("NEW_RELEASE_TEMPLATE", d.NotificationTemplateType, d.NotificationTemplate{
		Name:    "NEW_RELEASE_TEMPLATE",
		Subject: "New Release <<release-name>> is available",
		Content: `   A new release has been published and is now available to deploy.
   Available artefacts:
   <<release-artefacts>>`,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("ISSUE_TEMPLATE", d.NotificationTemplateType, d.NotificationTemplate{
		Name:    "ISSUE_TEMPLATE",
		Subject: "Issue on new release <<release-name>>",
		Content: `   A new release has been published but failed ingestion:
   <<issue-log>>
`,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("CLAM_SCAN_CMD", d.CommandType, d.Command{
		Name:        "CLAM_SCAN_CMD",
		Description: "scan files in specified path using clamav",
		Value:       "echo Infected files: 0",
		ErrorRegex:  ".*Infected files: [^0].*",
		StopOnError: true,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = s.Save("TEST_PIPELINE", d.PipelineType, d.PipelineConf{
		Name:           "TEST_PIPELINE",
		InboundRoutes:  []string{"TEST_IN_ROUTE"},
		OutboundRoutes: []string{"TEST_OUT_ROUTE"},
		Commands:       []string{"CLAM_SCAN_CMD"},
		CMDB: &d.CMDB{
			Catalogue: true,
			Events:    []string{"SETUP"},
			Tag:       []string{"TEST"},
		},
		SuccessNotification:   "SUCCESSFUL_RELEASE_NOTIFICATION",
		ErrorNotification:     "FAIL_ERROR_NOTIFICATION",
		CmdFailedNotification: "FAIL_SCAN_NOTIFICATION",
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
}
