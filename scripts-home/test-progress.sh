#!/bin/bash

echo "" > /tmp/test.out 

log() {
  echo $@ >> /tmp/test.out 
}

log "Progress Gauge testing."
echo "Progress Gauge testing."
./cmdline --title "This is a progress Gauge" --gauge "Progress of the nation..." 6 50 5  

count=10
while true 
do
sleep 3 
./update --title "This is a progress Gauge" --gauge "Progress of the nation..." 6 50 $count  
count=`expr $count + 10`

if [ x$count == x100 ] ; then 
   break
fi

done

./cmdline --title "OK, All done. " --msgbox "Now you can quit.." 8 70

