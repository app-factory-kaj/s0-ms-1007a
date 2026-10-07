# Greeter

## Problem Statement

Teams building services on this platform need a small, working reference
service that shows the expected shape of a Go HTTP service end to end — so
they have something concrete to compare a new service against, rather than
starting from a blank page.

## Solution

Greeter is a minimal Go HTTP service that returns a personalized JSON greeting
for a name passed on the request. It exists as a small, complete example:
simple enough to read in full, complete enough to show the real conventions.

## Actors

- API Caller — any client (human or system) that calls the Greeter HTTP API.

## Features

- F1 [Greeting](features/F1-greeting.md)

## Product-wide

See [Product-wide](product-wide.md).

## Out of Scope

- A user interface. Greeter is an API-only service.
- Any sign-in, accounts or per-user data.
- Storing or recalling past greetings.

## Further Notes

- The service follows the conventions of the `app-factory-kaj/e2e-reference`
repository for its structure and implementation. \[idea\]

