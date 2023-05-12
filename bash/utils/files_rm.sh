#!/bin/bash

function CmdExist() {
	for cmd in "$@"
	do
		if ! command -v $cmd &> /dev/null; then
			echo "($cmd) is not exist"
			echo "Exit"
			exit 1
		fi
	done
}

function RemoveFiles() {
	# chech the number of files larger than 0
	if [[ $# -eq 0 ]] ; then
		echo "no files"
		return
	fi

  # remove all files
	for file in "$@"
	do
		# Check file if exist
		if [[ ! -e $file ]] ; then
			echo "$file is not exist"
			break
		fi

		# Delete file
		if ! rm $file &> /dev/null ; then
			echo "($file) fail to remove"
		fi
	done
}

function files_rm() {
	local files=(
		file0.tmp
		file1.tmp
		file2.tmp
	)
	local emptyFiles=()
	# Check rm command exist
	CmdExist rm

	echo "Start remove empty file array"
	#rm empty files
	RemoveFiles ${emptyFiles[*]}
	echo "done, must output (no files)"

	echo "Start remove"
	# touch files
	for file in ${files[*]}
	do
		touch $file
	done

	#rm files
	RemoveFiles ${files[*]}
	echo "done"
}


files_rm