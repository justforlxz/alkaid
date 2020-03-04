/*
 * Copyright 2020. The Alkaid Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 * Alkaid is a BaaS service based on Hyperledger Fabric.
 */

package docker

import (
	"context"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Controller struct {
	cli *client.Client
}

func NewController() (*Controller, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Controller{cli: cli}, err
}

type CreateRequest struct {
	ImageName string
	ImageTag  string
}

func (c *Controller) Create(cr *CreateRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	config := &container.Config{}
	c.cli.ContainerCreate(ctx, config, nil, nil, "")

	return nil
}

func (c *Controller) Restart() error {
	return nil
}

func (c *Controller) Stop() error {
	return nil
}

func (c *Controller) Delete() error {
	return nil
}

func (c *Controller) CreateNetworkWithDockerMode(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := c.cli.NetworkCreate(ctx, name, dockertypes.NetworkCreate{
		CheckDuplicate: true,
	})
	if err != nil {
		return "", err
	}

	return result.ID, err
}
