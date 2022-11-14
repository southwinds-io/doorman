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

type NotificationType string

const (
	SuccessNotification   NotificationType = "SUCCESS"
	CmdFailedNotification NotificationType = "CMD_FAILED"
	ErrorNotification     NotificationType = "ERROR"
)

func (p *Process) FindNotification(name string) (*doorman.PipeNotification, error) {
	result, err := p.src.Load(name, new(doorman.Notification))
	if err != nil {
		return nil, err
	}
	notif := result.(*doorman.Notification)
	// load the template
	result, err = p.src.Load(notif.Template, new(doorman.NotificationTemplate))
	if err != nil {
		return nil, err
	}
	templ := result.(*doorman.NotificationTemplate)
	return &doorman.PipeNotification{
		Name:      notif.Name,
		Recipient: notif.Recipient,
		Type:      notif.Type,
		Subject:   templ.Subject,
		Content:   templ.Content,
	}, nil
}
