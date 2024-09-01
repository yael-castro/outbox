# Transactional Outbox Pattern + Hexagonal pattern

This repository is a proof of concept (PoC) in which I implement [the transactional outbox pattern](https://microservices.io/patterns/data/transactional-outbox.html)
with [the hexagonal pattern](https://alistair.cockburn.us/hexagonal-architecture/) to have the `server` and
the `message relay` in the same code base to avoid having different projects sharing databases.
(Also this project can be used as a template)