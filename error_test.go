package connectgrpcerr_test

import (
	"database/sql"
	"errors"
	"testing"

	connect "connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	connectgrpcerr "github.com/franchb/grpc-connect-go-errors"
)

func TestFromGRPCError(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		err  error
		code connect.Code
	}{
		{name: "NoError", err: nil, code: 0},
		{name: "OK", err: status.Error(codes.OK, ""), code: 0},
		{name: "Cancelled", err: status.Error(codes.Canceled, ""), code: connect.CodeCanceled},
		{name: "Unknown", err: status.Error(codes.Unknown, ""), code: connect.CodeUnknown},
		{name: "InvalidArgument", err: status.Error(codes.InvalidArgument, ""), code: connect.CodeInvalidArgument},
		{name: "DeadlineExceeded", err: status.Error(codes.DeadlineExceeded, ""), code: connect.CodeDeadlineExceeded},
		{name: "NotFound", err: status.Error(codes.NotFound, ""), code: connect.CodeNotFound},
		{name: "AlreadyExists", err: status.Error(codes.AlreadyExists, ""), code: connect.CodeAlreadyExists},
		{name: "PermissionDenied", err: status.Error(codes.PermissionDenied, ""), code: connect.CodePermissionDenied},
		{name: "ResourceExhausted", err: status.Error(codes.ResourceExhausted, ""), code: connect.CodeResourceExhausted},
		{name: "FailedPrecondition", err: status.Error(codes.FailedPrecondition, ""), code: connect.CodeFailedPrecondition},
		{name: "Aborted", err: status.Error(codes.Aborted, ""), code: connect.CodeAborted},
		{name: "OutOfRange", err: status.Error(codes.OutOfRange, ""), code: connect.CodeOutOfRange},
		{name: "Unimplemented", err: status.Error(codes.Unimplemented, ""), code: connect.CodeUnimplemented},
		{name: "Internal", err: status.Error(codes.Internal, ""), code: connect.CodeInternal},
		{name: "Unavailable", err: status.Error(codes.Unavailable, ""), code: connect.CodeUnavailable},
		{name: "DataLoss", err: status.Error(codes.DataLoss, ""), code: connect.CodeDataLoss},
		{name: "Unauthenticated", err: status.Error(codes.Unauthenticated, ""), code: connect.CodeUnauthenticated},
		{name: "UnknownCode", err: status.Error(codes.Code(1000), ""), code: connect.CodeInternal},
		{name: "OtherError", err: status.Error(codes.OK, ""), code: 0},
		{name: "nil", err: nil, code: 0},
		{name: "sql.ErrNoRows", err: sql.ErrNoRows, code: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			result := connectgrpcerr.FromGRPCError(testCase.err)
			if testCase.err == nil {
				if result != nil {
					t.Errorf("expected nil error on nil input, got %v", result)
				}
				return
			}
			var connectErr *connect.Error
			ok := errors.As(result, &connectErr)
			if ok {
				if _, isStatusErr := status.FromError(testCase.err); !isStatusErr && !errors.Is(result, testCase.err) {
					t.Errorf("passed non-status.Error (%v), expected this error wrapped, got %v", testCase.err, result)
				}
			} else if !(errors.Is(result, testCase.err) && result.Error() == testCase.err.Error()) {
				t.Errorf("non-status error passed in FromGRPCError should be wrapped with connect.Error, got %v", result)
			}
		})
	}
}

func TestFromGRPCErrorCodeOK(t *testing.T) {
	t.Parallel()
	original := status.Error(codes.OK, "")
	result := connectgrpcerr.FromGRPCError(original)
	if result != nil {
		t.Errorf("got %v, want nil", result)
	}
}
