# ROS

ROS, pronounced "rose", is a remote, multi-user system that is intended to be
hosted on a secure network.

User features:

- Multiple user accounts
- Arbitrary file storage, with user permissions
- A text editor
- A file browser
- An image viewer

System features:

- Support for off-site, encrypted data backup
- All communication 2-way encrypted
- Persistent data:
    - All data encrypted at rest
    - Persistent data centrally stored
    - Persistent data backed up remotely, completely encrypted (maybe daily snapshots?)

## Licensing

There are exist two projects, "ros-open" and "ros-proprietary". "ros-open" is
the [open-source](LICENSE) core codebase. "ros-proprietary" is a commercial product that
can be licensed with support for a fee.

## Dependencies

Depends on:

* cgit
* lighttpd
* docker
* go (for running tools)

## Principles

- Open source
- Minimal 3P dependencies


