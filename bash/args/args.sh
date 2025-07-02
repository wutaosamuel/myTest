#!/bin/bash

# default value 
input_file=""
all_proxy="http://172.16.63.20:10089"

# parse arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        -i|--input)
            input_file="$2"
            shift 2
            ;;
				--all-proxy)
            all_proxy="$2"
            shift 2
            ;;
        *)
            echo "Unknown argument: $1"
            exit 1
            ;;
    esac
done

echo "input file: $input_file"
echo "all proxy: $all_proxy"
