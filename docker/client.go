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
 *     Initial: 2017/07/15        Yang Chenglong
 *     Modify : 2017/07/16        Li Zebang
 */

package docker

import (
	"go-docker-pratice/libs/log"

	"github.com/docker/docker/client"
	"go.uber.org/zap"
)

type DockerClient struct {
	C *client.Client
}

func Test() {

}

func NewDockerClient() (*DockerClient, error) {
	c, err := client.NewEnvClient()

	if err != nil {
		log.Logger.Error("Create docker client with error", zap.Error(err))
		return nil, err
	}

	return &DockerClient{C: c}, nil
}

func (cli *DockerClient) ClientVersion() string {
	return cli.C.ClientVersion()
}
