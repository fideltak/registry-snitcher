package app

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/types"
	log "github.com/sirupsen/logrus"
)

type ContainerImage struct {
	ImageName string `env:"RS_IMAGE_NAME,required"`
	OsType    string `env:"RS_OS_TYPE" envDefault:"linux"`
	CpuArch   string `env:"RS_CPU_ARCH" envDefault:"amd64"`
}

func (c *ContainerImage) Inspect() error {
	log.Debugf("%s", c)
	if c.ImageName == "" {
		log.Errorf("Image name: %s", os.Getenv("RS_IMAGE_NAME"))
		panic(fmt.Errorf("No target image name"))
	}

	log.Info(fmt.Sprintf("The target image is %s", c.ImageName))
	res_str := fmt.Sprintf("//%s", c.ImageName)
	log.Debug(res_str)
	ref, err := docker.ParseReference(res_str)
	if err != nil {
		return err
	}

	sc := types.SystemContext{
		OSChoice:                    c.OsType,
		ArchitectureChoice:          c.CpuArch,
		OCIInsecureSkipTLSVerify:    true,
		DockerInsecureSkipTLSVerify: types.OptionalBoolTrue,
	}
	ctx := context.Background()
	img, err := ref.NewImage(ctx, &sc)
	if err != nil {
		return err
	}
	defer img.Close()
	b, _, err := img.Manifest(ctx)
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("%s", string(b)))
	return nil
}
