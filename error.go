package connectgrpcerr

import (
	"slices"

	connect "connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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



	var result *connect.Error

	//nolint:exhaustive // codes.OK returns always a nil error handled at the first line of function.
	switch errorStatus.Code() {
	// case codes.OK:
	// The zero code in gRPC is OK, which indicates that the operation was a
	// success. We don't define a constant for it because it overlaps awkwardly
	// with Go's error semantics: what does it mean to have a non-nil error with
	// an OK status? (Also, the Connect protocol doesn't use a code for
	// successes.)
	case codes.Canceled:
		result = connect.NewError(connect.CodeCanceled, err)
	case codes.Unknown:
		result = connect.NewError(connect.CodeUnknown, err)
	case codes.InvalidArgument:
		result = connect.NewError(connect.CodeInvalidArgument, err)
	case codes.DeadlineExceeded:
		result = connect.NewError(connect.CodeDeadlineExceeded, err)
	case codes.NotFound:
		result = connect.NewError(connect.CodeNotFound, err)
	case codes.AlreadyExists:
		result = connect.NewError(connect.CodeAlreadyExists, err)
	case codes.PermissionDenied:
		result = connect.NewError(connect.CodePermissionDenied, err)
	case codes.ResourceExhausted:
		result = connect.NewError(connect.CodeResourceExhausted, err)
	case codes.FailedPrecondition:
		result = connect.NewError(connect.CodeFailedPrecondition, err)
	case codes.Aborted:
		result = connect.NewError(connect.CodeAborted, err)
	case codes.OutOfRange:
		result = connect.NewError(connect.CodeOutOfRange, err)
	case codes.Unimplemented:
		result = connect.NewError(connect.CodeUnimplemented, err)
	case codes.Internal:
		result = connect.NewError(connect.CodeInternal, err)
	case codes.Unavailable:
		result = connect.NewError(connect.CodeUnavailable, err)
	case codes.DataLoss:
		result = connect.NewError(connect.CodeDataLoss, err)
	case codes.Unauthenticated:
		result = connect.NewError(connect.CodeUnauthenticated, err)
	default:
		result = connect.NewError(connect.CodeInternal, err)
	}

	details := make([]*connect.ErrorDetail, 0, len(errorStatus.Details()))
	for detail := range slices.Values(errorStatus.Details()) {
		msg, ok := detail.(proto.Message)
		if !ok {
			continue
		}
		errorDetail, errorDetailErr := connect.NewErrorDetail(msg)
		if errorDetailErr == nil {
			details = append(details, errorDetail)
		}
	}
	for detail := range slices.Values(errorStatus.Details()) {
		msg, ok := detail.(proto.Message)
		if !ok {
			continue
		}
		errorDetail, errorDetailErr := connect.NewErrorDetail(msg)
		if errorDetailErr == nil {
			result.AddDetail(errorDetail)
		}
	}
	return result
}
