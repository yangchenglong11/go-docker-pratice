/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/07/16        Li Zebang
 */

package docker

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"go.uber.org/zap"

	"go-docker-pratice/libs/log"
)

func (cli *DockerClient) ContainerStart(containerID string) error {
	err := cli.C.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		log.Logger.Error("Start a container with error", zap.Error(err))
	}

	return err
}

func (cli *DockerClient) ContainerStop(containerID string) error {
	timeout := new(time.Duration)
	err := cli.C.ContainerStop(context.Background(), containerID, timeout)
	if err != nil {
		log.Logger.Error("Start a container with error", zap.Error(err))
	}

	return err
}

func (cli *DockerClient) ContainerKill(containerID string) error {
	return cli.C.ContainerKill(context.Background(), containerID, "KILL")
}

func (cli *DockerClient) ContainerRemove(containerID string) error {
	cli.ContainerStop(containerID)

	err := cli.C.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{RemoveVolumes: false})
	if err != nil {
		log.Logger.Error("Remove a container with error", zap.Error(err))
	}

	return err
}

func (cli *DockerClient) ContainerLogs(containerID string) io.ReadCloser {
	readerLogs, err := cli.C.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: true, Details: true})
	if err != nil {
		log.Logger.Error("Run ContainerLogs with error", zap.Error(err))
	}

	return readerLogs
}

func (cli *DockerClient) ContainerList() []types.Container {
	containers, err := cli.C.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Logger.Error("Run ContainerList with error", zap.Error(err))
	}

	return containers
}

func (cli *DockerClient) Containers() *Containers {
	containers := new(Containers)
	for _, container := range cli.ContainerList() {
		container, _ := cli.FindByID(container.ID)
		containers.Containers = append(containers.Containers, container)
	}

	return containers
}

func (cli *DockerClient) FindByID(ID string) (*Container, error) {
	container, err := cli.C.ContainerInspect(context.Background(), ID)
	if err != nil {
		log.Logger.Error("Get container logs with error" + ID, zap.Error(err))
	}

	return &Container{Container: &container}, err
}

type Container struct {
	Container *types.ContainerJSON
}

func (c *Container) GetStatus() string {
	return c.Container.State.Status
}

func (c *Container) GetImage() string {
	return c.Container.Image
}

func (c *Container) GetPorts() nat.PortMap {
	return c.Container.NetworkSettings.Ports
}

func (c *Container) GetVolumes() map[string]struct{} {
	return c.Container.Config.Volumes
}

func (c *Container) GetRunningTime() time.Duration {
	start, err := time.Parse("2006-01-02T15:04:05.999999999Z", c.Container.State.StartedAt)
	if err != nil {
		log.Logger.Error("Parse time in GetContainerRunningTime function with error", zap.Error(err))
	}

	return time.Now().UTC().Sub(start)
}

func (c *Container) GetNetworkSettings() *types.NetworkSettingsBase {
	return &c.Container.NetworkSettings.NetworkSettingsBase
}

type Containers struct {
	Containers []*Container
}
