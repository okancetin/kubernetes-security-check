FROM golang:1.10
WORKDIR $GOPATH/src/kubeconfiginit

COPY . .
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["kube-security-check"]