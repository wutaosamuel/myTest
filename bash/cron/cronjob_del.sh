#!/bin/bash

homedir=$HOME
displayFile="${homedir}/cronjob.log"
cmd="/usr/bin/echo \"Add this line at $(/usr/bin/date)\" >> ${displayFile}"
cronjob="* * * * * ${cmd}"

(crontab -l 2>/dev/null; grep -v "${cronjob}") | crontab -