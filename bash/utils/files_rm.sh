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
	for file in "$@"
	do
		# Check file if exist
		if ! test -e $file ; then
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
	# Check rm command exist
	CmdExist rm

	# rm empty files
	RemoveFiles ${files[*]}

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