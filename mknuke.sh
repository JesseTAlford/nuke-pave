#!/bin/bash
# set -e

go build
cf uninstall-plugin Nuke-and-Pave
cf install-plugin nuke-pave
