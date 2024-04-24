#!/bin/bash

myDir="$(dirname $0)"
testFile="$myDir/printf_cat.sh"

printf "%b" "$(cat $testFile)"