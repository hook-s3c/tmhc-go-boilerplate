# Use phusion/baseimage as base image. To make your builds reproducible, make
# sure you lock down to a specific version, not to `latest`!
# See https://github.com/phusion/baseimage-docker/blob/master/Changelog.md for
# a list of version numbers.
FROM phusion/baseimage:0.11

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]

# ...put your own build instructions here...
RUN apt-get update
RUN apt-get install tmux git golang neovim make -y

# set the GOPATH environment variable
ENV GOPATH="${HOME}/tmhc-norvegicus"
ENV PATH="${GOPATH}/bin:${PATH}"

# Go dependencies / tooling
RUN go get github.com/smartystreets/goconvey
