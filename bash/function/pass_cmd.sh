#!/bin/bash

function CmdFailExist() {
	if "$@"; then
		echo "$* --> execute succeeded" 
	else
		echo "$* --> execute failed"
		exit 1
	fi
}

echo "Start"
CmdFailExist python3 --version
echo "End"