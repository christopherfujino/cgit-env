#!/usr/bin/env bash

set -euo pipefail

fallback_to_shell() {
  echo "Trapped SIGINT. Falling back to /bin/bash..." 1>&2
  exec /bin/bash
}

trap fallback_to_shell INT

# -D - Don't get to background
lighttpd -D -f /etc/lighttpd/lighttpd.conf
