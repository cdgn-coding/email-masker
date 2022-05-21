# Email Masker v1

## Summary

As stated in the [inception document](./Inception.md), Email Masker
is an open-source project to help users have control of who is capable to
reach them through emails. This document describes the first implementation of Email Masker.

## Domain abstraction

## Receiving emails

SendGrid is an all-in-one service for managing emails.
Email Masker will use SendGrid's [Inbound Parse Webhook](https://docs.sendgrid.com/for-developers/parsing-email/setting-up-the-inbound-parse-webhook)
to receive emails as POST requests.

To accomplish this, Email Masker should have:

* A registered domain
* A hosted zone with a subdomain MX record pointing to SendGrid, the inbound subdomain.
* An endpoint capable of receiving POST requests from the internet, called the receiving endpoint.

### Receiving endpoint security

Because this endpoint can receive a request from the internet, it should
verify that the request comes from SendGrid, otherwise, the communication should be
terminated as soon as possible.

To achieve this, Email Masker will use the [Signed Webhook Feature](https://docs.sendgrid.com/for-developers/tracking-events/getting-started-event-webhook-security-features).
This feature adds a Header `X-Twilio-Email-Event-Webhook-Signature` to the post request which allows us to assert that the request
was originated from Twilio by using an asymmetric encryption method (Elliptic Curve Digital Signature Algorithm).
Twilio signs the request using a private key and the endpoint verifies the headers using the public key.
Twilio's [documentation](https://docs.sendgrid.com/for-developers/tracking-events/getting-started-event-webhook-security-features#verify-the-signature) includes much detail about how to do this verification.

### Receiving endpoint parameters

## Email masks mapping

## Infrastructure
