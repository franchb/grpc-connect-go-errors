package connectgrpcerr

import (
	connect "connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FromGRPCError maps the gRPC error to the corresponding connect-go error code.
// If the error is nil, it returns nil. If the error cannot be converted to a
// gRPC error, it returns a connect-go internal error. Otherwise, it returns the
// connect-go error code based on the gRPC status code.
func FromGRPCError(err error) error {
	if err == nil {
		return nil
	}
	errorStatus, ok := status.FromError(err)
	if !ok {
		return connect.NewError(connect.CodeInternal, err)
	}

	//nolint:exhaustive // codes.OK returns always a nil error handled at the first line of function.
	switch errorStatus.Code() {
	// case codes.OK:
	// The zero code in gRPC is OK, which indicates that the operation was a
	// success. We don't define a constant for it because it overlaps awkwardly
	// with Go's error semantics: what does it mean to have a non-nil error with
	// an OK status? (Also, the Connect protocol doesn't use a code for
	// successes.)
	case codes.Canceled:
		return connect.NewError(connect.CodeCanceled, err)
	case codes.Unknown:
		return connect.NewError(connect.CodeUnknown, err)
	case codes.InvalidArgument:
		return connect.NewError(connect.CodeInvalidArgument, err)
	case codes.DeadlineExceeded:
		return connect.NewError(connect.CodeDeadlineExceeded, err)
	case codes.NotFound:
		return connect.NewError(connect.CodeNotFound, err)
	case codes.AlreadyExists:
		return connect.NewError(connect.CodeAlreadyExists, err)
	case codes.PermissionDenied:
		return connect.NewError(connect.CodePermissionDenied, err)
	case codes.ResourceExhausted:
		return connect.NewError(connect.CodeResourceExhausted, err)
	case codes.FailedPrecondition:
		return connect.NewError(connect.CodeFailedPrecondition, err)
	case codes.Aborted:
		return connect.NewError(connect.CodeAborted, err)
	case codes.OutOfRange:
		return connect.NewError(connect.CodeOutOfRange, err)
	case codes.Unimplemented:
		return connect.NewError(connect.CodeUnimplemented, err)
	case codes.Internal:
		return connect.NewError(connect.CodeInternal, err)
	case codes.Unavailable:
		return connect.NewError(connect.CodeUnavailable, err)
	case codes.DataLoss:
		return connect.NewError(connect.CodeDataLoss, err)
	case codes.Unauthenticated:
		return connect.NewError(connect.CodeUnauthenticated, err)
	default:
		return connect.NewError(connect.CodeInternal, err)
	}
}
