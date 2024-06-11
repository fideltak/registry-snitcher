FROM ubuntu:24.04

ARG USERNAME=snitcher
ARG GROUPNAME=snitchers
ARG UID=9999
ARG GID=9999

RUN groupadd -g $GID $GROUPNAME && \
    useradd -m -s /bin/bash -u $UID -g $GID $USERNAME
USER $USERNAME
WORKDIR /home/$USERNAME/

ADD ./bin/${BIN_FILE_NAME} /home/$USERNAME
CMD /home/snitcher/registry-snitcher