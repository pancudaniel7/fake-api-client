FROM golang:1.15.6

# Create app dir
RUN mkdir /app

# Copy all needed packages to app dir
COPY . /app

# Set wokrdir
WORKDIR /app

# Build lib
CMD go install -ldflags "-X main.Version=v1.0.0" ./...
CMD go test -v --tags=integration ./...