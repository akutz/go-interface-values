## --------------------------------------
## Build diagrams from ASCII art (ditaa)
## --------------------------------------
FROM openjdk:11 as ditaa-builder

RUN apt-get update -y && \
    apt-get install -y leiningen
RUN git clone https://github.com/akutz/ditaa.git /ditaa && \
    git -C /ditaa checkout feature/set-font
RUN cd /ditaa && lein uberjar


## --------------------------------------
## Go interface values
## --------------------------------------
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
## Java
## --------------------------------------

# Install the OpenJRE.
RUN apt-get install -y openjdk-11-jre-headless

# Configure the Java environment variables.
ARG TARGETARCH
ENV JAVA_HOME="/usr/lib/jvm/java-11-openjdk-${TARGETARCH}"


## --------------------------------------
## Diagrams from ASCII art (ditaa)
## --------------------------------------
COPY --from=ditaa-builder \
     /ditaa/target/ditaa-0.11.0-standalone.jar \
     /ditaa.jar
ENV DITAA="java -jar /ditaa.jar"
ENV LANG="C.UTF-8"


## --------------------------------------
## Main
## --------------------------------------

# Create the working directory.
RUN mkdir -p /go-interface-values
WORKDIR /go-interface-values/

# Copy the current repo into the working directory.
COPY . /go-interface-values/
