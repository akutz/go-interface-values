# Docker

Docker is the preferred approach to running the examples from this repository.

* [**Build the image**](#build-the-image): build the Docker image locally
* [**Use the image**](#use-the-image): use the Docker image to run the examples in this repository

## Build the image

Build the Docker image by executing the following commands:

1. Clone this repository locally:

    ```bash
    git clone https://github.com/akutz/go-interface-values
    ```

1. Change directories to the root of the cloned repository:

    ```
    cd go-interface-values
    ```

1. Build the image:

    ```bash
    docker build -t go-interface-values .
    ```

And voilá, you have the image!

---

:star: **Please note** the image is also available remotely and can obtained by executing the following commands:

```bash
docker pull akutz/go-interface-values && \
docker tag akutz/go-interface-values go-interface-values
```

---

## Use the image

The image may be used one of three ways:

* [**Self-contained**](#self-contained): run the examples from the image's filesystem
* [**Mount repository**](#mount-repository): run the examples from your local filesystem using the image
* [**VS Code dev container**](#dev-container): run the examples in VS Code using the `Remote - Containers` extension

### Self-contained

This method does not require cloning this repository locally since the Docker image includes the contents of this repository. For example, the following command will run the boxing benchmark using the examples sources from inside the image:

```bash
docker run -it --rm go-interface-values go test -v ./examples/
```

### Mount repository

Alternatively, it is also possible to run the examples cloned on your local filesystem using the Docker image. This has two benefits:

* The examples will be up-to-date
* It is easier to capture file output from the examples such as profiles

For example, the following command will:

* Run the benchmarks
* Produce a memory profile
* Produce an SVG image from the memory profile

```bash
docker run -it --rm -v "$(pwd):/go-interface-values" \
  go-interface-values bash -c ' \
  go test -bench . -run Bench -benchmem -benchtime 100x -memprofile mem.profile && \
  go tool pprof -svg mem.profile'
```

At the end there will be a file in the current, local directory named `profile001.svg` that visualizes the memory profile produced from the benchmark.

### Dev container

Alternatively, it is also possible to run the examples through opening the cloned reposity in a VS Code devcontainer:

1. To get started, please install the extension `Remote - Containers` by Microsoft (ms-vscode-remote.remote-containers). 
1. Open up your Command Palette and type or select the command `Remote-Containers: Reopen in Container`.

    The extension should be able to pick up the settings described in the `.devcontainer/devcontainer.json` and build the Dockerfile
1. Open up your terminal and you should be inside a container and at the project root. Verify this with:

    ```bash 
    go version
    ```

    If you see `go version go1.18beta2 linux/arm64` then you should be good to go.

---

Next: [Interface values](../02-interface-values/)
