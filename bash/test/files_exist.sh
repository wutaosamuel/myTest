#!/bin/bash

function CheckFiles() {
	for file in "$@"
	do
		# Check file if exist
		if ! test -e $file ; then
			echo "$file is not exist"
			break
		fi
	done
}

CheckFiles ./file1 ./file2