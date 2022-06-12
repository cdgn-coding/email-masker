package sendgrid

import "errors"

var SendGridOutboundEmailError = errors.New("error requesting to send an email with SendGrid")
