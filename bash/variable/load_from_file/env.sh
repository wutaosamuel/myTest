#!/bin/bash

shellDir="$(dirname $0)"
envFile="$shellDir"/env.txt

source "$envFile"

echo "MY_ENV=$MY_ENV"
echo "MY_ENV1=$MY_ENV1"
echo "MY_ENV2=$MY_ENV2"