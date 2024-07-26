#!/bin/bash

tarFile=$1
sourceFile=$2

tar -czvf "$tarFile" "$sourceFile"
