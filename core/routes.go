/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	"fmt"
	"southwinds.dev/types/doorman"
	"strings"
)

func (p *Process) MatchInboundRoutes(serviceId, bucketName string) (routes []*doorman.InRoute, err error) {
	items, err := p.getInRoutes()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		route := item.(*doorman.InRoute)
		if strings.EqualFold(route.BucketName, "*") && route.ServiceId == serviceId {
			routes = append(routes, route)
		}
	}
	if routes != nil {
		if len(routes) > 1 {
			return []*doorman.InRoute{routes[0]}, nil
		} else {
			return routes, nil
		}
	} else {
		// if no match is found, then run a search using bucket name
		for _, item := range items {
			route := item.(*doorman.InRoute)
			if strings.EqualFold(route.BucketName, bucketName) && route.ServiceId == serviceId {
				routes = append(routes, route)
			}
		}
	}
	return routes, nil
}

func (p *Process) FindInboundRoutesByWebHookToken(token string) (routes []*doorman.InRoute, err error) {
	items, err := p.getInRoutes()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		route := item.(*doorman.InRoute)
		if strings.EqualFold(route.WebhookToken, token) {
			routes = append(routes, route)
		}
	}
	return routes, nil
}

func (p *Process) FindAllInRoutes() (routes []*doorman.InRoute, err error) {
	items, err := p.getInRoutes()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		routes = append(routes, item.(*doorman.InRoute))
	}
	return routes, nil
}

func (p *Process) getInRoutes() (routes []any, err error) {
	routes, err = p.src.LoadItemsByType(func() any {
		return new(doorman.InRoute)
	}, doorman.InboundRouteType)
	if err != nil {
		return nil, fmt.Errorf("cannot load inbound routes: %s", err)
	}
	return routes, err
}
