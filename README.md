# registry-snitcher
Simple monitoring tool for container registry.

## Architecture
This tool retrieves metadata(manifest) of a container in the registry which you want to monitor.
Which means if failed to retrieve metadata, your registry may be down.
You can monitor the succeeded/failed access count from your prometheus via an exporter in this tool.

## Configurations
You can set env values for the configurations.

| Key | Default Value | Required | Example | Description |
| ---- | ---- | ---- | ---- | ---- |
| RS_IMAGE_NAME | - | Y | docker.io/busybox:latest | The container image name in your registry which you want to monitor. |
| RS_OS_TYPE | linux | N | linux | OS type of the container image. |
| RS_CPU_ARCH | amd64 | N | arm64 | CPU architecture of the conatainer image. |
| RS_DURATION_SEC | 60 | N | 20 | Pull image metadata duration seconds. |
| RS_DEBUG | - | N | YES | Debug mode. No value indicates turning debug off. |