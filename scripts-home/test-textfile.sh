#================= Progress test. TODO
log "Infobox Testing." 
echo "Infobox Testing." 
./cmdline --title "Progress test still not done." --infobox "Yes to Continue" 8 70
EXITCODE=$?
log Exitcode After count line is: $EXITCODE
if [ x$EXITCODE == x0 ] ; then 
      CHOICE="YES"
else
      CHOICE="NO"
fi
echo "Infobox box Returned:" $CHOICE 
log "Infobox box Returned:" $CHOICE 

