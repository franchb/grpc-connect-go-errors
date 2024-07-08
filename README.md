# grpc-connect-go-errors

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/franchb/grpc-connect-go-errors.svg)](https://pkg.go.dev/github.com/franchb/grpc-connect-go-errors)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/franchb/grpc-connect-go-errors)
![GitHub License](https://img.shields.io/github/license/franchb/grpc-connect-go-errors)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/franchb/grpc-connect-go-errors/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/franchb/grpc-connect-go-errors)](https://goreportcard.com/report/github.com/franchb/grpc-connect-go-errors)
[![Codecov](https://codecov.io/gh/franchb/grpc-connect-go-errors/branch/main/graph/badge.svg)](https://codecov.io/gh/franchb/grpc-connect-go-errors)


## Description

This small Go package converts `google.golang.org/grpc` error codes to `connectrpc.com/connect-go` error codes.

## Motivation

During the migration process from gRPC to connect-go, developers often face the challenge of 
translating `google.golang.org/grpc` errors to their `connect-go` equivalents. 
This package aims to address these issues by offering a streamlined conversion process.

### Specific Use Case: Migrating from grpc-gateway to connect-go

A common scenario where this package proves particularly useful is when migrating from grpc-gateway to connect-go:

1. In existing [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) implementations, 
2. the `status.Error` with appropriate `codes.Code` [is automatically translated](https://github.com/grpc-ecosystem/grpc-gateway/blob/cd8ebb289a410418160f42c68a5fcbe67067a102/runtime/errors.go#L36) to the corresponding HTTP status.
3. When wrapping a gRPC service implementation with a connect-go handler, errors returned by the gRPC service are passed to the connect-go handler. However, connect-go expects explicit `connect.Code` and `connect.Error`.
4. Without proper error conversion, all errors from the gRPC service will result in HTTP 500 status codes when handled by connect-go.

This package solves this issue by allowing developers to wrap errors returned by the gRPC service 
before passing them to the connect-go handler, ensuring correct HTTP status codes are maintained during the migration process.

I've created an [issue](https://github.com/connectrpc/connect-go/issues/763) in the connect-go repo, and @jhump responded:

 > One big objective of Connect is providing support for HTTP 1.1 for web and mobile RPC clients, another is to 
 > provide libraries that are lightweight & simple and that use standard (or widely used) 
 > libraries and idioms for the target language. So pulling the behemoth that is grpc-go into connect-go's 
 > dependency graph is a non-starter.
 >
https://github.com/connectrpc/connect-go/issues/763


@jhump's insightful feedback was instrumental in shaping the direction of this project. 
After considering his response, I decided to address this specific need by implementing this small Go package. 
It provides a focused solution for converting gRPC error codes to connect-go error codes without introducing heavy dependencies, 
aligning with connect-go's philosophy of lightweight and simple libraries.

## Usage

```go
package controller

import (
	"context"
	"github.com/franchb/grpc-connect-go-errors"
)

// ...

func (ps *PingServer) Ping(
        ctx context.Context,
        req *connect.Request[pingv1.PingRequest],
) (*connect.Response[pingv1.PingResponse], error) {
	response, err := ps.grpccontroller.Ping(ctx, ...)
	if err != nil {
		return nil, connectgrpcerr.FromGRPCError(err)
	}

	// ...other code here
}
```

## Status: Alpha

This package is in its alpha stage and should be considered unstable.


## License

Offered under the [Apache 2 license](LICENSE).
