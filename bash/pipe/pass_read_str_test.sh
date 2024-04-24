#!/bin/bash

myDir="$(dirname $0)"
testSh="$myDir/pass_read_str.sh"

echo "this is a test string" | bash $testSh
