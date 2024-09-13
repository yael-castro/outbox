# Transactional outbox pattern + Hexagonal pattern
This repository is a proof of concept (PoC) in which I implement the [transactional outbox pattern](https://microservices.io/patterns/data/transactional-outbox.html)
with [the hexagonal pattern](https://alistair.cockburn.us/hexagonal-architecture/) to have the `service/sender` and
the `message relay` in the same code base to avoid having different projects sharing databases.
(Also this project can be used as a template)

![Component diagram](./docs/images/components.svg)

As the diagram shows, this project is compiled in two parts:
1. `outbox-http` It is the binary in charge of manage the http requests related with purchases (service/sender)
2. `outbox-relay` It is the binary in charge of read and send the messages (message relay)

## Architecture decisions
###### Go project layout standard
I decided to follow the [Go project layout standard](https://github.com/golang-standards/project-layout).
###### Package tree
I built the package tree following the concepts of the [hexagonal architecture pattern](https://alistair.cockburn.us/hexagonal-architecture/).
```
.
├── cmd
└── internal
    ├── app
    │   ├── business (Use cases, rules, data models and ports)
    │   ├── input    (Everything related to "drive" adapters)
    │   └── output   (Everything related to "driven" adapters)
    ├── container (DI container)
    └── pkg       (Public and global code, potencially libraries)
```
###### Compile only what is required
According to the theory of hexagonal architecture, it is possible to have *n* adapters for different external signals (http, gRPC, command line).

So I decided to compile a binary to handle each signal.

### Database schema
The database schema is described by the `.sql` files in the [sql](./scripts/sql) directory.

To build the database run the following command.
```shell
make database
```

### Docs
[See the OpenAPI specification](./docs/OpenAPI.json)

[See the required environment variables](.env.example)

### How to install
###### Purchase service (sender)
```shell
go install -tags http github.com/yael-castro/outbox/cmd/outbox-http
```
###### Message relay
```shell
go install -tags relay github.com/yael-castro/outbox/cmd/outbox-relay
```
### How to use from source
All compiled binaries will put in the `build` directory
###### Purchase service (sender)
```shell
make http
./build/outbox-http
```
###### Message relay
```shell
make relay
./build/outbox-relay
```
