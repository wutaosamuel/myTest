#!/bin/bash

function confirm() {
	while true; do
		read -r -p "$* [y/n](default: n): " input
		case $input in
		y | Y | yes | YES | Yes)
			return 1
			;;
		*)
			return 0
			;;
		esac
	done
}

confirm "Are you sure?"
getConfirm=$?
echo "get confrom: ${getConfirm}"
