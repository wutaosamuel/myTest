#!/bin/bash

function MultipleLine() {
	echo \
		"line1
line2
line3"
	echo ""

	lines="new line
line 2
line 3"
	echo "$lines"
}

MultipleLine