# syntax=docker/dockerfile:experimental

FROM ubuntu:22.04
LABEL maintainer="Quang V. Nguyen <quang@2meo.com>"
LABEL org.opencontainers.image.description "Base Image"

RUN \
  apt-get update && \
  RUNLEVEL=1 DEBIAN_FRONTEND=noninteractive \
  apt-get install -y --no-install-recommends \
  aptitude \
  autoconf \
  automake \
  bash-completion \
  build-essential \
  ca-certificates \
  curl \
  gcc \
  git \
  git-core \
  iputils-ping \
  jq \
  net-tools \
  pkg-config \
  python3-pip \
  python3-setuptools \
  rsync \
  telnet \
  tree \
  tzdata \
  vim \
  wget \
  zip \
  zlib1g-dev && \
  find /usr/local -depth \
  \( \
  \( -type d -a -name test -o -name tests \) \
  -o \
  \( -type f -a -name '*.pyc' -o -name '*.pyo' \) \
  \) -exec rm -rf '{}' +;

CMD ["/bin/bash"]
