#!/bin/bash

multi_line=$(cat <<- END
Title here
This is the line 1
This is the line 2
END
)

echo "$multi_line"