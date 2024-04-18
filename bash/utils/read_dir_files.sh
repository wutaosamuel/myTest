#!/bin/bash

function dirExist() {
	if [ ! -d $1 ]; then
		return 0
	fi
	return 1
}

function readDirFiles() {
	# check directory exist
	dirExist $1
	local isExist=$?
	if [ $isExist != 1 ]; then
		echo "$1 is not exist or not a directory"
		exit 1
	fi

	# read directory files
	for file in "$1"/*; do
		echo ""
		echo "$file"
		local dirname=$(dirname $file)
		local filename=$(basename ${file%.*})
		local extention=${file##*.}
		local jpgname="$dirname"/jpg/"$filename".jpg
		echo "$dirname"
		echo "$filename"
		echo "$extention"
		echo "$jpgname"
	done
}

readDirFiles $1
