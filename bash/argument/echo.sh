#!/bin/bash

function printAry() {
	arr=("$@")
	echo "arr length: ${#arr[@]}"

	for arg in "${arr[@]}";
	do
		echo $arg
	done
}

function echoArg1() {
	echo "from echoArg1"
}

function echoArg2() {
	echo "from echoArg2"
}

arg1=$(echoArg1)

array=(
	"from echo Arg3"
	"$arg1"
	"$(echoArg2)"
)

printAry "${array[@]}"
echo ""
printAry "from cmd" "$arg1" "$(echoArg2)"