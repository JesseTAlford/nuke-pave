#!/bin/bash
# set -e

go build main.go
cf uninstall-plugin Nuke-and-Pave
cf install-plugin nuke-pave
