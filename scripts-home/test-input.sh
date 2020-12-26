#!/bin/bash

echo "" > /tmp/test.out 

log() {
  echo $@ >> /tmp/test.out 
}

log "Input box Testing." 
echo "Input box Testing." 
CHOICE=`./cmdline --inputbox "What is your favorite Color?" 8 78 Blue --title "Example Dialog"` 
EXITCODE=$?
log Exitcode After count line is: $EXITCODE
log CHOICE is $CHOICE
echo "Infobox box Returned:" $CHOICE and Exitcode is $EXITCODE

