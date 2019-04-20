FROM golang:1.11.4-alpine3.8 AS build

# Install tools required to build the project.
# We need to run `docker build --no-cache .` to update those dependencies.
RUN apk add --no-cache git
ENV DEP_RELEASE_TAG=v0.5.0
RUN wget -O - https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Gopkg.toml and Gopkg.lock lists project dependencies.
# These layers are only re-built when Gopkg files are updated.
COPY Gopkg.lock Gopkg.toml /go/src/github.com/pinaraws/mxnet-operator/
WORKDIR /go/src/github.com/pinaraws/mxnet-operator/

# Install library dependencies.
RUN dep ensure -vendor-only

RUN mkdir -p /opt/kubeflow

# Copy all project and build it.
# This layer is rebuilt when ever a file has changed in the project directory.
COPY . /go/src/github.com/pinaraws/mxnet-operator/
RUN go build -o /opt/kubeflow/mxnet-operator.v1beta1 github.com/pinaraws/mxnet-operator/cmd/mxnet-operator.v1beta1

FROM alpine:3.8
COPY --from=build /opt/kubeflow/mxnet-operator.v1beta1 /opt/kubeflow/mxnet-operator.v1beta1
RUN chmod a+x /opt/kubeflow/mxnet-operator.v1beta1

CMD ["/opt/kubeflow/mxnet-operator.v1beta1", "--alsologtostderr", "-v=1"]
