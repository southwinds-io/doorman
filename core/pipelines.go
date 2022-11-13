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
)

func (p *Process) FindPipeline(pipeName string) (*doorman.Pipeline, error) {
	result, err := p.src.Load(pipeName, new(doorman.PipelineConf))
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve pipeline %s: %s", pipeName, err)
	}
	pipeConf := result.(doorman.PipelineConf)
	var inRoutes []doorman.InRoute
	for _, route := range pipeConf.InboundRoutes {
		result, err = p.src.Load(route, new(doorman.InRoute))
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve inbound route %s: %s", route, err)
		}
		inRoute := result.(doorman.InRoute)
		inRoutes = append(inRoutes, inRoute)
	}
	var outRoutes []doorman.OutRoute
	for _, route := range pipeConf.OutboundRoutes {
		result, err = p.src.Load(route, new(doorman.OutRoute))
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve outbound route %s: %s", route, err)
		}
		outRoute := result.(doorman.OutRoute)
		outRoutes = append(outRoutes, outRoute)
	}
	var cmds []doorman.Command
	for _, cmd := range pipeConf.Commands {
		result, err = p.src.Load(cmd, new(doorman.Command))
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve command %s: %s", cmd, err)
		}
		cmdObj := result.(doorman.Command)
		cmds = append(cmds, cmdObj)
	}
	successN, err := p.FindNotification(pipeConf.SuccessNotification)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve success notification %s: %s", pipeConf.SuccessNotification, err)
	}
	cmdFailedN, err := p.FindNotification(pipeConf.CmdFailedNotification)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve command failed notification %s: %s", pipeConf.CmdFailedNotification, err)
	}
	errorN, err := p.FindNotification(pipeConf.ErrorNotification)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve error notification %s: %s", pipeConf.ErrorNotification, err)
	}
	pipe := &doorman.Pipeline{
		Name:                  pipeConf.Name,
		InboundRoutes:         inRoutes,
		OutboundRoutes:        outRoutes,
		Commands:              cmds,
		SuccessNotification:   successN,
		CmdFailedNotification: cmdFailedN,
		ErrorNotification:     errorN,
		CMDB:                  pipeConf.CMDB,
	}
	return pipe, nil
}

func (p *Process) MatchPipelines(serviceId, bucketName string) ([]doorman.Pipeline, error) {
	var (
		pipes    []doorman.Pipeline
		routes   []doorman.InRoute
		pipeline *doorman.Pipeline
		err      error
	)
	routes, err = p.MatchInboundRoutes(serviceId, bucketName)
	if err != nil {
		return nil, err
	}
	if routes == nil {
		return nil, fmt.Errorf("no inbound routes for service id '%s' found in doorman configuration", serviceId)
	}
	pipeConfs, err := p.FindPipelinesConfigurationWithInboundRoutes()
	if err != nil {
		return nil, err
	}
	for _, conf := range pipeConfs {
		pipeline, err = p.FindPipeline(conf.Name)
		if err != nil {
			return nil, err
		}
		if err = pipeline.Valid(); err != nil {
			return nil, err
		}
		pipes = append(pipes, *pipeline)
	}
	return pipes, nil
}

func (p *Process) FindPipelinesConfigurationWithInboundRoutes() (pipelines []doorman.PipelineConf, err error) {
	var items []any
	items, err = p.getPipelinesConf()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		pipeline := item.(doorman.PipelineConf)
		if len(pipeline.InboundRoutes) > 0 {
			pipelines = append(pipelines, pipeline)
		}
	}
	return pipelines, nil
}

func (p *Process) getPipelinesConf() (routes []any, err error) {
	routes, err = p.src.LoadItemsByType(func() any {
		return new(doorman.PipelineConf)
	}, doorman.PipelineType)
	if err != nil {
		return nil, fmt.Errorf("cannot load pipelines: %s", err)
	}
	return routes, err
}
