#!/bin/bash

echo "" > /tmp/test.out 

log() {
  echo $@ >> /tmp/test.out 
}

log Starting now

echo "Entering the script"
echo "These should go to logs."
count=1
while ./cmdline --title "Current count is $count" --yesno "Yes to Continue or No to Quit." 8 70
do
   EXITCODE=$?
   log Exitcode After count line is: $EXITCODE
   if [ x$EXITCODE == x0 ] ; then 
      CHOICE="YES"
      log YES chosen continue in loop
   else
      CHOICE="NO"
      log NO chosen breaking loop
      break
   fi
   count=`expr $count + 1`
done

log Out of loop 
log echoing again to UI logs 
echo "First Choice:" $CHOICE 
echo "And It returned: " $RESULT


log Showing new second UI widget 
./cmdline --title "Second Example" --yesno "This is Second example yes/no box." 8 70 
EXITCODE=$?
log echoing final log exit code is $EXITCODE 
echo "Second Exit code:" $EXITCODE
log Done. 
sync

