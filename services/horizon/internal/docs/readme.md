---
title: Horizon
---

Horizon is the server for the client facing API for the Stellar ecosystem.  It acts as the interface between [dfc-core](https://www.dfc.org/developers/learn/dfc-core) and applications that want to access the Stellar network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Stellar ecosystem](https://www.dfc.org/developers/guides/) for more details.

You can interact directly with horizon via curl or a web browser but SDF provides a [JavaScript SDK](https://www.dfc.org/developers/js-dfc-sdk/learn/) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net [https://horizon-testnet.dfc.org/](https://horizon-testnet.dfc.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/dfc/js-dfc-sdk)
- [Java](https://github.com/dfc/java-dfc-sdk)
- [Go](https://github.com/dfc/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/dfc/ruby-dfc-sdk)
- [Python](https://github.com/StellarCN/py-dfc-base)
- [C# .NET 2.0](https://github.com/QuantozTechnology/csharp-dfc-base)
- [C# .NET Core 2.x](https://github.com/elucidsoft/dotnetcore-dfc-sdk)
- [C++](https://bitbucket.org/bnogal/dfcqore/wiki/Home)
