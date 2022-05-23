package emails

import "fmt"

type MaskAddressNotFoundError struct {
	Err     error
	Address string
}

func NewMaskAddressNotFoundError(address string, err error) *MaskAddressNotFoundError {
	return &MaskAddressNotFoundError{
		Address: address,
		Err:     err,
	}
}

func (e *MaskAddressNotFoundError) Error() string {
	return fmt.Sprintf("Mask %s does not exist. %v", e.Address, e.Err)
}

type OutboundEmailError struct {
	Err error
}

func NewOutboundEmailError(err error) *MaskAddressNotFoundError {
	return &MaskAddressNotFoundError{
		Err: err,
	}
}

func (e *OutboundEmailError) Error() string {
	return fmt.Sprintf("Error sending email. %v", e.Err)
}
