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
		containers.containers = append(containers.containers, container)
	}

	return containers
}

func (cli *DockerClient) FindByID(ID string) (*types.ContainerJSON, error) {
	container, err := cli.C.ContainerInspect(context.Background(), ID)
	if err != nil {
		log.Logger.Error("Get container logs with error" + ID, zap.Error(err))
	}

	return &container, err
}

type ContainerInfo struct {
	*types.ContainerJSON
}

func (c *ContainerInfo) GetStatus() string {
	return c.ContainerJSON.State.Status
}

func (c *ContainerInfo) GetImage() string {
	return c.ContainerJSON.Image
}

func (c *ContainerInfo) GetPorts() nat.PortMap {
	return c.ContainerJSON.NetworkSettings.Ports
}

func (c *ContainerInfo) GetVolumes() map[string]struct{} {
	return c.ContainerJSON.Config.Volumes
}

func (c *ContainerInfo) GetRunningTime() time.Duration {
	start, err := time.Parse("2006-01-02T15:04:05.999999999Z", c.State.StartedAt)
	if err != nil {
		log.Logger.Error("Parse time in GetContainerRunningTime function with error", zap.Error(err))
	}

	return time.Now().UTC().Sub(start)
}

func (c *ContainerInfo) GetNetworkSettings() *types.NetworkSettingsBase {
	return &c.NetworkSettings.NetworkSettingsBase
}

type Containers struct {
	containers []*types.ContainerJSON
}
