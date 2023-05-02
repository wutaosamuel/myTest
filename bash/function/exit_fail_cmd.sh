#!/bin/bash

function CaptureStderr() {
	local cmd="python3 --version"
	local wrong_cmd="python --version"
	$cmd

	python --version

	echo "CaptureStderr done"
}

function ErrorExit() {
	if python3 --version ; then
		echo "command succeeded"
	else
		echo "command failed"
	fi

	if pytho --version ; then
		echo "command succeeded"
	else
		echo "command failed"
		exit 1
	fi
}

echo "start"
ErrorExit
echo "end"