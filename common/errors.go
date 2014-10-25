package common

import (
	"errors"
)

var (
	// common error message
	ErrUnknown       = errors.New("unknown error")
	ErrInvalidParam  = errors.New("invalid input param")
	ErrTimeout       = errors.New("request timeout")
	ErrInvalidStatus = errors.New("status exception")
	ErrNullValue     = errors.New("value is NULL")
	ErrNotAllowed    = errors.New("operation not allow")
	ErrInvalidData   = errors.New("invalid data info")
	ErrNoPrivelige   = errors.New("no privelige")

	// server inner related error message
	ErrEntryExist    = errors.New("entry already exist")
	ErrEntryNotExist = errors.New("entry not exist")
	ErrMsgTooLarge   = errors.New("message payload too large")
	ErrInvalidMsg    = errors.New("check msg format invalid")
	ErrDecryptMsg    = errors.New("decrypt the msg failed")
	ErrEncryptMsg    = errors.New("encrypt the msg failed")

	// account related error message
	ErrInvalidEmail    = errors.New("email address invalid")
	ErrInvalidPhone    = errors.New("phone number invalid")
	ErrAccountExist    = errors.New("account already exist")
	ErrAccountNotExist = errors.New("account not exist")
	ErrPasswordWrong   = errors.New("account password wrong")
	ErrInvalidName     = errors.New("login name invalid")
	ErrInvalidSign     = errors.New("check signature invalid")
	ErrSignTimeout     = errors.New("check signature timeout")

	// dev related error message
	ErrDeviceIsSlave  = errors.New("device is slave")
	ErrDeviceIsMaster = errors.New("device is master")
	ErrAlreadyBinded  = errors.New("device already binded")
	ErrNotYetBinded   = errors.New("device not yet binded")
	ErrMasterNotExist = errors.New("master device not exist")
	ErrSlaveNotExist  = errors.New("slave device not exist")
	ErrInvalidDevice  = errors.New("device info invalid")

	// project related error message
	ErrProjectNotExist = errors.New("project not exist")
	ErrEntryNotDir     = errors.New("entry is not a directory")
)
