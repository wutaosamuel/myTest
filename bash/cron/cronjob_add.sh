#!/bin/bash

homedir=$HOME
displayFile="${homedir}/cronjob.log"
cmd="Add this line"
cronjob="* * * * *"

echo $cmd

# https://stackoverflow.com/questions/42198960/how-to-add-a-crontab-job-to-crontab-using-a-bash-script
# https://stackoverflow.com/questions/4880290/how-do-i-create-a-crontab-through-a-script