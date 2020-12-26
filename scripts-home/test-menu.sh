#!/bin/bash

echo "" > /tmp/test.out 
log() {
  echo $@ >> /tmp/test.out 
}

log Starting now

echo "Entering the script"
CHOICE=`./cmdline --title "Menu example" --menu "Choose an option" 25 78 16 "<-- Back" "Return to the main menu." \
"Add User" "Add a user to the system." \
"Modify User" "Modify an existing user." \
"List Users" "List all users on the system." \
"Add Group" "Add a user group to the system." \
"Modify Group" "Modify a group and its list of members." \
"List Groups" "List all groups on the system."`
EXITCODE=$?
echo "Choice was: " $CHOICE  " Exitcode was" $EXITCODE
log Exitcode After count line is: $EXITCODE
sync


