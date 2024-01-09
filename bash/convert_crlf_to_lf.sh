#!/bin/bash

for file in $@; do
	if [ -f $file ]; then
		sed -i 's/\r//g' $file
		echo "${file} convert to LF done"
	fi
done
