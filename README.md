# script-manager

Allows you to:
* start/stop scripts from browser. Scripts listed in config.yaml.
* Get UI on browser, when you use whiptail like command line to drive user dialogs. See the usage below.

## How to add dialogs for getting UI on web
* Lookup 'whiptail' command line syntax. We support almost all the UI widgets such as message box, info box, checklists, menus, input box, password, progress gauge and so on.
* Use 'cmdline' instead of whiptail. (Or create a symlink if you don't want to change the script)
* Use 'update', which is symlink to 'cmdline', to send updates to widgets which support updates, such as progress and infobox. 

## Examples

See ./scripts-home/ directory to see examples. Usually exit code tells if there is success (exit code 0) or failure (other than 0):


```
echo "This is script. This line will be shown in logs."
if ./cmdline --title "Do you like to go to Moon?" --yesno "OK to Continue, Cancel to cancel." 8 70
then
  echo "We got Yes as an answer"
else
  echo "We got No as an answer"
fi
```

Here is an example of use of cmdline which communicates results back to main script (in case of say, checklists, menu, input etc.) by sending on stdout:

```
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
```


## How to Build

Front end is Vue/Vutify/Vuex based. Do 'yarn build' to build and populate ./dist static files directory first. 
And then run the go server:

```
cd go-server
./build.sh 
```

## How it works

When script uses:
```
./cmdline --title "Welcome to Message box. Only OK." --msgbox "OK to Continue" 8 70
```
the args are output on stderr, with prefix ":cmdline --title ...". Web server, which started this script, watches stderr. When it receives
a line with ":cmdline", it will interpret it as widget request, and then drives the UI to show the widget to user. (We use Vue for final widgets.)
And response is transmitted on stdin, back to cmdline, which is currently incharge of stdin, stdout, stderr.
All other lines are considered as script logs. 

