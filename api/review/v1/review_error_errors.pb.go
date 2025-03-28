// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsNeedLogin(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NEED_LOGIN.String() && e.Code == 401
}

// 为某个枚举单独设置错误码
func ErrorNeedLogin(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_NEED_LOGIN.String(), fmt.Sprintf(format, args...))
}

func IsDbFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DB_FAILED.String() && e.Code == 500
}

func ErrorDbFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DB_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsOrderReviewd(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ORDER_REVIEWD.String() && e.Code == 400
}

func ErrorOrderReviewd(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_ORDER_REVIEWD.String(), fmt.Sprintf(format, args...))
}

func IsResultNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RESULT_NOT_FOUND.String() && e.Code == 404
}

func ErrorResultNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_RESULT_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsReviewReplyAlreadyExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ReviewReplyAlreadyExist.String() && e.Code == 410
}

func ErrorReviewReplyAlreadyExist(format string, args ...interface{}) *errors.Error {
	return errors.New(410, ErrorReason_ReviewReplyAlreadyExist.String(), fmt.Sprintf(format, args...))
}

func IsStoreIDNotMatch(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_StoreIDNotMatch.String() && e.Code == 411
}

func ErrorStoreIDNotMatch(format string, args ...interface{}) *errors.Error {
	return errors.New(411, ErrorReason_StoreIDNotMatch.String(), fmt.Sprintf(format, args...))
}
