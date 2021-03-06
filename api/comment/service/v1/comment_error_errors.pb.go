// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsCommentServiceErrorReasonUnknownError(err error) bool {
	e := errors.FromError(err)
	return e.Reason == CommentServiceErrorReason_COMMENT_SERVICE_ERROR_REASON_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorCommentServiceErrorReasonUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CommentServiceErrorReason_COMMENT_SERVICE_ERROR_REASON_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}
