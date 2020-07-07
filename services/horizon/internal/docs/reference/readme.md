---
title: Overview
---

Horizon is an API server for the Stellar ecosystem.  It acts as the interface between [dfc-core](https://github.com/dfc/dfc-core) and applications that want to access the Stellar network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Stellar ecosystem](https://www.dfc.org/developers/guides/) for details of where Horizon fits in. You can also watch a [talk on Horizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Stellar.org developer Scott Fleckenstein:

[![Horizon: API webserver for the Stellar network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Horizon: API webserver for the Stellar network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Horizon provides a RESTful API to allow client applications to interact with the Stellar network. You can communicate with Horizon using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Stellar SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.dfc.org/developers/js-dfc-sdk/learn/index.html) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net: [https://horizon-testnet.dfc.org/](https://horizon-testnet.dfc.org/) and one that is connected to the public Stellar network:
[https://horizon.dfc.org/](https://horizon.dfc.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/dfc/js-dfc-sdk)
- [Java](https://github.com/dfc/java-dfc-sdk)
- [Go](https://github.com/dfc/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/dfc/ruby-dfc-sdk)
- [Python](https://github.com/StellarCN/py-dfc-base)
- [C#](https://github.com/elucidsoft/dotnet-dfc-sdk)
