#!/bin/bash

function confirm() {
	return 1
}

function echoMessage() {
	echo "$*"
}

confirm
getConfirm=$?
echo "get confrom: ${getConfirm}"
echo ""

message=$(echoMessage "Are you sure?")
echo $message