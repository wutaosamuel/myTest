#!/bin/bash

# USAGE: bash execute_from_argument.sh "echo \"hello world! \""
args=("$@")

for arg in "${args[@]}";
do
	echo "Entered cmd: $1"
	bash -c "$1"
	echo
done
