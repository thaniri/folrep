#!/bin/bash
# Script is used for running unit and integration tests
#

set -ex

go test -v ./...
