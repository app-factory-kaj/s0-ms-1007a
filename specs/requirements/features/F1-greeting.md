# Greeting

## Purpose

Return a personalized JSON greeting for a name supplied on the request.

## User Stories

- F1.1 As an API caller, I send a GET request to /hello with a `name` query
parameter and receive a JSON greeting that includes that name. \[idea\]
- F1.2 As an API caller, when I omit the `name` parameter, I receive a default
JSON greeting addressed to "World" rather than an error. *assumed*

## Decisions

- The service is a single Go HTTP service named `greeter`, built to the
conventions of the `app-factory-kaj/e2e-reference` repository. \[idea\]
- The response is JSON, containing at least the greeting message. \[idea\]

## Out of Scope

- Any greeting language or format other than a single JSON message.
- Rate limiting or quotas on the endpoint.