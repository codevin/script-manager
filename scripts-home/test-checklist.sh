#!/bin/bash

echo "" > /tmp/test.out 
log() {
  echo $@ >> /tmp/test.out 
}

log Starting now

# Original COLOR=$(cmdline --title "Fav Color $count" --inputbox "What is your favorite Color?" 8 78 Blue --title "Example Dialog" 3>&1 1>&2 2>&3)

echo "Entering the script"
echo "Input was: "

CHOICE=`./cmdline --title "Menu example" --checklist "Choose user's permissions" 20 78 4 \
"NET_OUTBOUND" "Allow connections to other hosts" ON \
"NET_INBOUND" "Allow connections from other hosts" OFF \
"LOCAL_MOUNT" "Allow mounting of local devices" OFF \
"REMOTE_MOUNT" "Allow mounting of remote devices" OFF`
EXITCODE=$?
echo "Choice was: " $CHOICE
log Chosen: $CHOICE, Count: $count 
log Exitcode After count line is: $EXITCODE
