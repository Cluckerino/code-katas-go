# Start with Ubuntu
FROM ubuntu:latest as go-buntu

# Install apt-utils
RUN DEBIAN_FRONTEND=noninteractive apt-get update -qq \
        && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        apt-utils

# Install golang
RUN DEBIAN_FRONTEND=noninteractive apt-get update -qq \
        && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        golang-go

# Install misc. tools
RUN DEBIAN_FRONTEND=noninteractive apt-get update -qq \
        && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        ca-certificates \
        git

# Setup go environment
ENV GOPATH /go
ENV GOSRC ${GOPATH}/src
ENV GOBIN ${GOPATH}/bin
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "${GOSRC}" "${GOBIN}" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

# Start from image with with Go
FROM go-buntu as go-tester

# Get the testify extension
RUN go get github.com/stretchr/testify

# Start from image with testify
FROM go-tester

# Set up the correct path for the katas
ENV CODE_KATAS_PATH ${GOSRC}/github.com/cluckerino/code-katas-go

# Actually make the folder
RUN mkdir -p $CODE_KATAS_PATH

WORKDIR $CODE_KATAS_PATH

# Copy repo to image (use '.' instead of '*' to preserve subdirs)
COPY . ./

# Run the go test runner
CMD ["go", "test", "./..."]