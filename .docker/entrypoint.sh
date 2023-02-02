#!/bin/sh

go mod tidy
apt-get update -y
apt-get install sqlite3 -y

tail -f /dev/null