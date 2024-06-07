#!/bin/bash

thisDir="$(dirname $0)"
targetFile="$(dirname $thisDir)"/utils/is_str_equal_func.sh

source "$targetFile"

Str1="Equal String"
Str2="Equal String"
Str3="Unequal String"

echo "Test equal string --> pass"
IsStringEqual "$Str1" "$Str2"
echo
echo "Test unequal string --> pass"
IsStringEqual "$Str1" "$Str3"