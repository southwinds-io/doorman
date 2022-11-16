/*
Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
Contributors to this project, hereby assign copyright in this code to the project,
to be licensed under the same terms as the rest of the code.
*/
package core

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
	src "southwinds.dev/source_client"
	d "southwinds.dev/types/doorman"
	"testing"
)

func TestGenConfig(t *testing.T) {
	godotenv.Load("../doorman.env")
	uri, _ := GetCfgURI()
	user, _ := GetCfgUser()
	pwd, _ := GetCfgPwd()
	s := src.New(uri, user, pwd, &src.ClientOptions{
		InsecureSkipVerify: true,
		Timeout:            0,
	})
	s.Logger = new(RetryLogger)
	// var err error
	testInRoute := d.InRoute{
		Name:           "TEST_IN_ROUTE",
		Description:    "a testing inbound route",
		ServiceHost:    "s3://127.0.0.1:9000",
		ServiceId:      "55913ea4-454a-4d15-a8d4-2b669154ab77",
		BucketName:     "*",
		User:           "admin",
		Pwd:            "password",
		AllowedAuthors: []string{},
	}
	b, _ := json.MarshalIndent(testInRoute, "", "  ")
	os.WriteFile("../data/in_route.json", b, os.ModePerm)
	// err = s.Save("TEST_IN_ROUTE", d.InboundRouteType, testInRoute)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	testOutRoute := d.OutRoute{
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
			BucketURI:    "s3://localhost:9000/ilink2",
			User:         "admin",
			Pwd:          "password",
			WebhookEvent: false,
			// Partition: "minio",
			// Service:   "sqs",
			// Region:    "",
			// AccountID: "_",
			// Resource:  "webhook",
		},
	}
	b, _ = json.MarshalIndent(testOutRoute, "", "  ")
	os.WriteFile("../data/out_route.json", b, os.ModePerm)
	// err = s.Save("TEST_OUT_ROUTE", d.OutboundRouteType, testOutRoute)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	successNotif := d.Notification{
		Name:      "SUCCESSFUL_RELEASE_NOTIFICATION",
		Recipient: "test@email.com",
		Type:      "email",
		Template:  "NEW_RELEASE_TEMPLATE",
	}
	b, _ = json.MarshalIndent(successNotif, "", "  ")
	os.WriteFile("../data/success_notification.json", b, os.ModePerm)
	// err = s.Save("SUCCESSFUL_RELEASE_NOTIFICATION", d.NotificationType, successNotif)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	failErrorNotif := d.Notification{
		Name:      "FAIL_ERROR_NOTIFICATION",
		Recipient: "test@email.com",
		Type:      "email",
		Template:  "ISSUE_TEMPLATE",
	}
	b, _ = json.MarshalIndent(failErrorNotif, "", "  ")
	os.WriteFile("../data/fail_error_notification.json", b, os.ModePerm)
	// err = s.Save("FAIL_ERROR_NOTIFICATION", d.NotificationType, failErrorNotif)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	failScanNotif := d.Notification{
		Name:      "FAIL_SCAN_NOTIFICATION",
		Recipient: "test@email.com",
		Type:      "email",
		Template:  "QUARANTINE_TEMPLATE",
	}
	b, _ = json.MarshalIndent(failScanNotif, "", "  ")
	os.WriteFile("../data/fail_scan_notification.json", b, os.ModePerm)
	// err = s.Save("FAIL_SCAN_NOTIFICATION", d.NotificationType, failScanNotif)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	quarantineTempl := d.NotificationTemplate{
		Name:    "QUARANTINE_TEMPLATE",
		Subject: "New Release <<release-name>> has been quarantined",
		Content: `A new release has been published and but has been quarantined.
   Error log:
   <<command-log>>`,
	}
	b, _ = json.MarshalIndent(quarantineTempl, "", "  ")
	os.WriteFile("../data/quarantine_template.json", b, os.ModePerm)
	// err = s.Save("QUARANTINE_TEMPLATE", d.NotificationTemplateType, quarantineTempl)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	releaseTempl := d.NotificationTemplate{
		Name:    "NEW_RELEASE_TEMPLATE",
		Subject: "New Release <<release-name>> is available",
		Content: `   A new release has been published and is now available to deploy.
   Available artefacts:
   <<release-artefacts>>`,
	}
	// err = s.Save("NEW_RELEASE_TEMPLATE", d.NotificationTemplateType, releaseTempl)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	b, _ = json.MarshalIndent(releaseTempl, "", "  ")
	os.WriteFile("../data/new_release_template.json", b, os.ModePerm)
	issueTempl := d.NotificationTemplate{
		Name:    "ISSUE_TEMPLATE",
		Subject: "Issue on new release <<release-name>>",
		Content: `   A new release has been published but failed ingestion:
   <<issue-log>>
`,
	}
	b, _ = json.MarshalIndent(issueTempl, "", "  ")
	os.WriteFile("../data/issue_template.json", b, os.ModePerm)
	// err = s.Save("ISSUE_TEMPLATE", d.NotificationTemplateType, issueTempl)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	cmd := d.Command{
		Name:        "CLAM_SCAN_CMD",
		Description: "scan files in specified path using clamav",
		Value:       "echo Infected files: 0",
		ErrorRegex:  ".*Infected files: [^0].*",
		StopOnError: true,
	}
	b, _ = json.MarshalIndent(cmd, "", "  ")
	os.WriteFile("../data/clam_cmd.json", b, os.ModePerm)
	// err = s.Save("CLAM_SCAN_CMD", d.CommandType, cmd)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
	pipeConf := d.PipelineConf{
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
	}
	b, _ = json.MarshalIndent(pipeConf, "", "  ")
	os.WriteFile("../data/pipe.json", b, os.ModePerm)
	// err = s.Save("TEST_PIPELINE", d.PipelineType, pipeConf)
	// if err != nil {
	//     t.Fatalf(err.Error())
	// }
}
