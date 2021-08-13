# Gateway Service

This is the Gateway service

Generated with

```
micro new --namespace=go.micro --type=service gateway
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.gateway
- Type: service
- Alias: gateway

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./gateway-service
```

Build a docker image
```
make docker
```