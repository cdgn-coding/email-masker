package emails

import (
	"errors"
)

var MaskAddressNotFoundError = errors.New("mask address not found")

var OutboundEmailError = errors.New("error while sending the email")

var UserNotFoundError = errors.New("user not found")
