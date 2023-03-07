package model

import (
	"errors"
)

const (
	Success Result = iota
	ErrInvalidSignature
	ErrInternal
	ErrTransactionAlreadySynced
	ErrNotExist
	ErrUserExpired
	ErrInvalidRequest
	ErrExceedRefundLimitTime
	ErrTransactionAlreadyRefunded
	ErrInvalidDurationForTransaction
	ErrInvalidPackageForTransaction
	ErrInvalidPartner
	ErrInvalidPurchaseDateForTransaction
	ErrInvalidRefundDate
)

var (
	ErrorRecordNotFound = errors.New("record not found")
)

type BasicResponse struct {
	Code    int         `json:"msg_code"`
	Message string      `json:"msg_content"`
	Data    interface{} `json:"msg_data"`
}

type SimpleErrorResponse struct {
	Code    int    `json:"msg_code"`
	Message string `json:"msg_content"`
}

func (s SimpleErrorResponse) Error() string {
	return s.Message
}

type Result int

func (s Result) StringMess() string {
	switch s {
	case ErrInvalidSignature:
		return "invalid signature"
	case ErrUserExpired:
		return "expired"
	case ErrNotExist:
		return "not exist"
	case ErrInternal:
		return "something went wrong with server"
	case ErrTransactionAlreadySynced:
		return "transaction had been synchronized"
	case ErrInvalidRequest:
		return "invalid request"
	case ErrExceedRefundLimitTime:
		return "refund date exceed limit time"
	case ErrTransactionAlreadyRefunded:
		return "transaction already refunded"
	case ErrInvalidDurationForTransaction:
		return "invalid duration for transaction"
	case ErrInvalidPackageForTransaction:
		return "invalid package for transaction"
	case ErrInvalidPartner:
		return "invalid partner"
	case ErrInvalidPurchaseDateForTransaction:
		return "invalid purchase date"
	case ErrInvalidRefundDate:
		return "invalid refund date"
	}
	return "unknown"
}
