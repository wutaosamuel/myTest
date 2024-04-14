#!/bin/bash

function dirExist() {
	if [ ! -d $1 ]; then
		echo "$1 is not exist"
		return 0
	fi
	return 1
}