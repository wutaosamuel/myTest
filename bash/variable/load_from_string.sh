#!/bin/bash

envString=$(cat <<- END
Name=Users
Age=18
END
)

name=
age=

if [ -z "$name"]; then
	echo "name is empty now before assign the value"
fi

source <(echo "$envString")
name=$Name
age=$Age

if [ -z "$name" ]; then
  echo "wrong, name is not empty now after assign the value"
else
	echo "After assign the value"
	echo "Name = $name"
	echo "Age = $age"
fi
