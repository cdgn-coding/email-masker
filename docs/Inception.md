# Email Masker Inception

## Table of contents

* Overview
* Goals and non-goals
* Requirements

## Overview

Nowadays email marketing is a powerful tool for companies
to reach us, potential customers. However, as users of the internet,
we do not have many tools to effectively filter all this information.
Email regulations are heterogeneous, and often useless. It is tough to trust companies
to keep our email addresses private and respect the contact opt-out.

Email masker pretends to solve this issue by turning the problem around. Instead
of using a primary email address to sign up, we propose to use disposable email addresses, called masks.
The existence of the mask is in the control of the user and can be deleted or disabled at any time.
This way, the user is in control of the communication.

## Goals and non-goals

The main goal of this project is to provide the capability of creating reliable and disposable email addresses.
This project should leverage existing technologies for this, for example, it should not create an SMTP or POP3 client
unless it is necessary for the main goal.

## Requirements

### Features

* Create, delete and edit masks
* Ease of remembering mask addresses
* Capability to disable and enable masks
* Allow the user to easily identify mask usage

### Values

* Scalability
* Cost-effectiveness
* Security
* Privacy
* Cloud independence
* Availability
* Failure tolerance
