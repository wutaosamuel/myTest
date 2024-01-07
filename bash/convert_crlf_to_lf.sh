#!/bin/bash

sed -i 's/\r//g'

for file in $@
do
	if [ -f $file ]; then
		sed -i 's/\r//g' $file
	fi
done