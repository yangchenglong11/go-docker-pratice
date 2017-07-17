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
	"os"

	"github.com/docker/docker/api/types"
	"go.uber.org/zap"

	"go-docker-pratice/libs/log"
)

func (cli *DockerClient) ImageList() *[]types.ImageSummary {
	images, err := cli.C.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		log.Logger.Error("List image(s) with error", zap.Error(err))
	}

	return &images
}

func (cli *DockerClient) ImagePull(image, tag string) error {
	out, err := cli.C.ImagePull(context.Background(), image + ":" + tag, types.ImagePullOptions{})
	if err != nil {
		log.Logger.Error("Pull image with error", zap.Error(err))
	}

	defer out.Close()

	io.Copy(os.Stdout, out)

	return err
}
