#================= Message test.
log "Message box Testing." 
echo "Message box Testing." 
./cmdline --title "Welcome to Message box. Only OK." --msgbox "OK to Continue" 8 70
EXITCODE=$?
log Exitcode After count line is: $EXITCODE
if [ x$EXITCODE == x0 ] ; then 
      CHOICE="YES"
else
      CHOICE="NO"
fi
echo Msgbox box Returned: $CHOICE   Exitcode= $EXITCODE
log Msgbox Returned: $CHOICE Exitcode= $EXITCODE

