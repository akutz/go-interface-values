FROM golang:1.18beta2


## --------------------------------------
## Authorship
## --------------------------------------

LABEL org.opencontainers.image.authors="sakutz@gmail.com"


## --------------------------------------
## Apt and standard packages
## --------------------------------------

# Update the local apt cache.
RUN apt-get update -y

# Install vim.
RUN apt-get install -y vim python3


## --------------------------------------
## Golang
## --------------------------------------

# Install graphviz so "go tool pprof" can export images.
RUN apt-get install -y graphviz

# Install the Go debugger.
RUN go install github.com/go-delve/delve/cmd/dlv@latest


## --------------------------------------
## Main
## --------------------------------------

# Create the working directory.
RUN mkdir -p /go-generics-the-hard-way
WORKDIR /go-interface-values/

# Copy the current repo into the working directory.
COPY . /go-interface-values/
