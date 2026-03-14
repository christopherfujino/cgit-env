# Distributed Compute

`A <==> B`

Based on https://wiki.archlinux.org/title/Cgit, with some paths changed for
Debian.

Depends on:

* cgit
* lighttpd
* docker
* go (for running tools)

## Principles

- Open source
- Minimal 3P dependencies

- All communication 2-way encrypted
- Persistent data:
    - All data encrypted at rest
    - Persistent data centrally stored
    - Persistent data backed up remotely, completely encrypted (maybe daily snapshots?)
