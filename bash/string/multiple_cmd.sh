#!/bin/bash

String1="this string from varialbe"

multi_line="$(cat <<- END
echo "title here"
echo "hello world!"
echo "$String1"
END
)"

echo "commands: "
echo "$multi_line"
echo