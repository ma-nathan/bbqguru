#!/bin/bash
cd "$(dirname "$0")"

LOGFILE=/home/nb/go/src/bbq/logfile

./bbq > $LOGFILE 2>&1
