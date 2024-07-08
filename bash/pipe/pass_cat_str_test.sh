#!/bin/bash

myDir="$(dirname $0)"
testSh="$myDir/pass_cat_str.sh"

message1="this is a test string"
message2="this is a test line 1
this is a test line 2
this is a test line 3"

echo "single line cat string, pass"
printf "%s" "$message1" | bash $testSh
echo ""
echo ""

echo "multiple line cat string, pass"
printf "%s" "$message2" | bash $testSh
echo ""
