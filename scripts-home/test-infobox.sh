#!/bin/bash

echo "" > /tmp/test.out 

log() {
  echo $@ >> /tmp/test.out 
}
#================= Infobox test.
log "Infobox Testing." 
echo "Infobox Testing." 
./cmdline --title "Welcome to Info box" --infobox "Keep Waiting." 8 70
EXITCODE=$?
log Exitcode After count line is: $EXITCODE
echo "You should keep seeing the box."

sleep 10
./update --title "Title got updated..." --infobox "Still keep waiting" 8 70
EXITCODE=$?
log Exitcode After count line is: $EXITCODE
sleep 10

./cmdline --title "OK, All done. " --msgbox "Now you can quit.." 8 70


echo "Message box done." 
log "Message box done." 


