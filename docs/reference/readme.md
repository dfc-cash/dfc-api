---
title: Overview
---

The Go SDK contains packages for interacting with most aspects of the dfc ecosystem.  In addition to generally useful, low-level packages such as [`keypair`](https://godoc.org/github.com/dfc/go/keypair) (used for creating dfc-compliant public/secret key pairs), the Go SDK also contains code for the server applications and client tools written in go.

## Godoc reference

The most accurate and up-to-date reference information on the Go SDK is found within godoc.  The godoc.org service automatically updates the documentation for the Go SDK everytime github is updated.  The godoc for all of our packages can be found at (https://godoc.org/github.com/dfc/go).

## Client Packages

The Go SDK contains packages for interacting with the various dfc services:

- [`horizon`](https://godoc.org/github.com/dfc/go/clients/horizon) provides client access to a horizon server, allowing you to load account information, stream payments, post transactions and more.
- [`dfctoml`](https://godoc.org/github.com/dfc/go/clients/dfctoml) provides the ability to resolve Stellar.toml files from the internet.  You can read about [Stellar.toml concepts here](../../guides/concepts/dfc-toml.md).
- [`federation`](https://godoc.org/github.com/dfc/go/clients/federation) makes it easy to resolve a dfc addresses (e.g. `scott*dfc.org`) into a dfc account ID suitable for use within a transaction.

