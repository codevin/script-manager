#================= input test.
log "Infobox Testing." 
echo "Infobox Testing." 
CHOICE=`whiptail --inputbox "What is your favorite Color?" 8 78 Blue --title "Example Dialog"` 
EXITCODE=$?
log Exitcode After count line is: $EXITCODE
log CHOICE is $CHOICE
echo "Infobox box Returned:" $CHOICE and Exitcode is $EXITCODE

