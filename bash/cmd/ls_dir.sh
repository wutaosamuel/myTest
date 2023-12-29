#!/bin/bash

function lsDir() {
	local dirs=(
		testDir0
		testDir1
		testDir2
	)

	# make test directories
	for dir in ${dirs[*]}
	do
		mkdir $dir
	done

	# list all directories
	for d in $(ls -d */)
	do
		echo "$d"
	done

	# remove directories
	for dir in ${dirs[*]}
	do
		rmdir $dir
	done
}

lsDir