#!/bin/bash

echo "" > /tmp/test.out 
log() {
  echo $@ >> /tmp/test.out 
}

log Starting now

# Original COLOR=$(cmdline --title "Fav Color $count" --inputbox "What is your favorite Color?" 8 78 Blue --title "Example Dialog" 3>&1 1>&2 2>&3)

echo "Entering the script"
echo "Input was: "

CHOICE=`./cmdline --title "Menu example" --radiolist "Choose user's permissions" 20 78 4 \
"NET_OUTBOUND" "Allow connections to other hosts" ON \
"NET_INBOUND" "Allow connections from other hosts" OFF \
"LOCAL_MOUNT" "Allow mounting of local devices" OFF \
"REMOTE_MOUNT" "Allow mounting of remote devices" OFF`
EXITCODE=$?
echo "Choice was: " $CHOICE
log Chosen: $CHOICE, Count: $count 
log Exitcode After count line is: $EXITCODE
if [ x$EXITCODE == x0 ] ; then 
         log $CHOICE chosen continue in loop
else
         log Escape pressed, no choice chosen breaking loop
         break
fi
log Showing new second UI widget 
./cmdline --title "OK, All good, Ending." --yesno "We will anyway end now." 8 70 
EXITCODE=$?
log echoing final log exit code is $EXITCODE 
echo "Second Exit code:" $EXITCODE
log Done. 
sync

