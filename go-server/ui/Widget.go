package ui

import (
  log "github.com/sirupsen/logrus"
  . "./types"
  // "./widgets"
  "strconv"
)

func InitWidgets() {
/*   WidgetsMap = make(map[string](func(*WidgetInfo)string), 50)
   WidgetsMap["InfoBox"] = widgets.InfoBox
   WidgetsMap["InputBoxWithOkCancel"] = widgets.InputBoxWithOkCancel
   WidgetsMap["MessageBox"] = widgets.MessageBox
   WidgetsMap["YesNoBox"] = widgets.YesNoBox
   WidgetsMap["YesNoBoxVue"] = widgets.YesNoBoxVue */
}


func getQuotedStrings(str string) []string {

       // Split by spaces, but not if we are within a quote.
       inQuote := false
       // single word char: false
       // space: true
       // quote: true, but within quote, ignore.
       quotedStrings := func(c rune) bool {
           switch {
             case c == '"' && !inQuote:
               inQuote = true
               return true
             case c == '"' && inQuote:
               inQuote = false
               return true
             case inQuote:
               return false
             default:
               return c == ' '
          }
       }

   return MyFieldsFunc(str, quotedStrings)
}

// f: true == field separator char. false == field char.
func MyFieldsFunc(str string, f func(c rune) bool) []string {

  r := []string{}
  part := ""
  // String may start with quote, so first char may be false.
  haveStr := false
  for _, rune := range str {
    if f(rune) {
      if (haveStr) {
         r = append(r, part)
         haveStr = false
      }
      part = ""
    } else {
      part = part + string(rune)
      haveStr = true
    }
  }
  if len(part) > 0 {
    r = append(r, part)
  }
  return r
}



/*
  Widget Invocation cmdlines:

  Number args: 
   - 8 70   -- height, width
   - 8 70 1  -- Here, 3rd argument is init value for the box, which could be a number instead of text. 

  MessageBox:  
         whiptail --title "Example Title" --msgbox "This is an example message box. Press OK to continue." 8 70
  InfoBox:  
         whiptail --title "Example Title" --infobox "This is an example info box." 8 70
  InputBox:  
         COLOR=$(whiptail --inputbox "What is your favorite Color?" 8 78 Blue --title "Example Dialog" 3>&1 1>&2 2>&3)
  YesNoBox:  
         (whiptail --title "Example Title" --yesno "This is an example yes/no box." 8 70); 
         -- exitvalue 0 == yes, 1== no

  Passwordbox: 
     PASSWORD=$(whiptail --passwordbox "Enter your new password" 8 70 --title "New Password" 3>&1 1>&2 2>&3)
     And exitstatus can be 0 or 1 (0 == OK, 1 == Cancel)

  Menu:
       whiptail --title "Menu example" --menu "Choose an option" 25 78 16 \
"<-- Back" "Return to the main menu." \
"Add User" "Add a user to the system." \
"Modify User" "Modify an existing user." \
"List Users" "List all users on the system." \
"Add Group" "Add a user group to the system." \
"Modify Group" "Modify a group and its list of members." \
"List Groups" "List all groups on the system."

  Checklist:

   whiptail --title "Check list example" --checklist \
"Choose user's permissions" 20 78 4 \
"NET_OUTBOUND" "Allow connections to other hosts" ON \
"NET_INBOUND" "Allow connections from other hosts" OFF \
"LOCAL_MOUNT" "Allow mounting of local devices" OFF \
"REMOTE_MOUNT" "Allow mounting of remote devices" OFF

  Radiolist:

  whiptail --title "Radio list example" --radiolist \
"Choose user's permissions" 20 78 4 \
"NET_OUTBOUND" "Allow connections to other hosts" ON \
"NET_INBOUND" "Allow connections from other hosts" OFF \
"LOCAL_MOUNT" "Allow mounting of local devices" OFF \
"REMOTE_MOUNT" "Allow mounting of remote devices" OFF

  ProgressGuage: 

Original, it was done like this:
#!/bin/bash
{
    for ((i = 0 ; i <= 100 ; i+=5)); do
        sleep 0.1; echo $i
    done
} | 
whiptail --gauge "Please wait while we are sleeping..." 6 50 0

Now we support this: 
    cmdline --gauge "Please wait while we are sleeping..." 6 50 0
And later: 
    cmdline --gauge_update 45
    // Progress in percent.

Common Options - Preceding widget specific options
================================================
  --clear - Clear screen on exit
  --defaultno  -- Open dialog with "No" as default
  --nocancel -- Open dialog without a  "No/Cancel" button 
  --yesbutton text  -- Set text of yes button
  --nobutton text  -- Set text of no button
  --okbutton text  -- Set text of ok button
  --cancelbutton text  -- Set text of cancel button
  --noitem
    The menu, checklist and radiolist widgets will display tags only, not the item strings. The menu widget still needs some items specified, but checklist and radiolist expect only tag and status. 
  --separate-output
    For checklist widgets, output result one line at a time, with no quoting. This facilitates parsing by another program. 
  --title title
    Specifies a title string to be displayed at the top of the dialog box. 
  --backtitle backtitle
    Specifies a backtitle string to be displayed on the backdrop, at the top of the screen. 
  --scrolltext
    Force the display of a vertical scrollbar. 

*/
// Output json object with equivalent info in cmdline.
// Only if cmdline starts with ":"
func parseCmdLine(script_name string, cmdline string, wi *WidgetInfo) {

   args := getQuotedStrings(cmdline)
   // fmt.Println("Length of args Collected=", len(args))
   // for _, item:=range(args) {
   //    fmt.Printf("===Start:%s:==END\n", item)
   // }

   getSubOption := func(idx *int) string {
     if (*idx + 1)  < len(args) {
          subopt :=  args[*idx + 1]
          if ( subopt[0] == '-' ) {
             // not a subopt
             return ""
          }
          *idx = *idx + 1
          return subopt
     }
     return ""
   }

   get_common := func(widget_name string, i *int) {
     wi.WidgetName = widget_name;
     wi.Message = getSubOption(i);
     wi.Height = getSubOption(i);
     wi.Width = getSubOption(i);
   }

   // Now process the args.
   tagsAreItems := false 
   // scrollText := false 
   i := 1  // drop :cmdline which is first arg
   for {
      if ( i >= len(args) ) {
           break
      }
      switch args[i] {
         case "--msgbox":
               get_common("MessageBoxVue", &i)

         case "--infobox":
               get_common("InfoBoxVue", &i)
               wi.WidgetType = "NewUpdatableVueWidget"

         case "--yesno":
               get_common("YesNoBoxVue", &i)

         case "--inputbox":
               get_common("InputBoxVue", &i)
               wi.DefaultValue = getSubOption(&i)

         case "--menu":  // --menu text height width menu-height [tag item]... 
               get_common("MenuBoxVue", &i)
               wi.MenuHeight = getSubOption(&i)

               wi.Items= make([]Item, 0, 20)
               for {
                  item := Item{}
                  tag := getSubOption(&i)
                  if (tag == "") {
                      break;
                  }
                  item.Value = tag
                  item.Label= tag
                  if ( !tagsAreItems) {
                     text := getSubOption(&i)
                     item.Label= text
                  }
                  wi.Items = append(wi.Items, item)
               }
         case "--passwordBox":
               get_common("PasswordBox", &i)

         case "--checklist":
               get_common("ChecklistBoxVue", &i)
               wi.MenuHeight = getSubOption(&i)

               wi.Items= make([]Item, 0, 20)
               for {
                  tag := getSubOption(&i)
                  if (tag == "") {
                      break;
                  }
                  item := Item{}
                  item.Value = tag
                  item.Label = tag
                  if ( !tagsAreItems) {
                      text := getSubOption(&i)
                      item.Label = text
                  }
                  status := getSubOption(&i)
                  if ( status == "on" ) {
                     item.Checked = true 
                  }
                  wi.Items = append(wi.Items, item)
               }

         case "--radiolist":
               get_common("RadiolistBoxVue", &i)
               wi.MenuHeight = getSubOption(&i)

               wi.Items= make([]Item, 0, 20)
               for {
                  tag := getSubOption(&i)
                  if (tag == "") {
                      break;
                  }
                  item := Item{}
                  item.Value = tag
                  item.Label = tag
                  if ( !tagsAreItems) {
                      text := getSubOption(&i)
                      item.Label= text
                  }
                  status := getSubOption(&i)
                  if ( status == "on" ) {
                     item.Checked = true 
                  }
                  wi.Items = append(wi.Items, item)
               }

         case "--gauge":
               get_common("ProgressBoxVue", &i)
               p, _ := strconv.Atoi(getSubOption(&i))
               wi.PercentCompleted = p
               wi.WidgetType = "NewUpdatableVueWidget"

         case "--gauge_update":
               p, _ := strconv.Atoi(getSubOption(&i))
               wi.PercentCompleted = p

         case "--clear":
               wi.CloseWidgetAtEnd = true

         case "--defaultno":
               wi.DefaultIsNo = true

         case "--nocancel":
               wi.NoCancelButton = true

         case "--yesbutton":  
               wi.YesButtonText =  getSubOption(&i)

         case "--nobutton":  
               wi.NoButtonText =  getSubOption(&i)

         case "--okbutton": 
               wi.OkButtonText =  getSubOption(&i)

         case "--cancelbutton": 
               wi.CancelButtonText =  getSubOption(&i)

         case "--noitem ":  // Only tags are not provided in radiobox, checkbox etc. Assume tags as items.
               tagsAreItems = true // TODO Don't know this

         case "--separate-output":
               wi.SendOneOutputAtATime = true

         case "--title": 
               wi.Title = getSubOption(&i)

         case "--backtitle":
               wi.SubText = getSubOption(&i)

         case "--scrolltext":
               // scrollText = true

         case "--default-item":
               wi.DefaultValue = getSubOption(&i)

         default :
               log.Warning("Widgets: Unknown option received when parsing:", args[i])
       }
       i = i + 1
     }
     return
}

//  cmdline starts with ":cmdline "
func CreateNewWidgetRequest(script_name string, cmdline string) *WidgetInfo {
    wi := CreateWidgetInfo(script_name)
    wi.WidgetType = "NewVueWidget" 
    // For updatable widgets, use NewUpdatableVueWidget
    parseCmdLine(script_name, cmdline, wi)
    return wi
}


//  cmdline starts with ":update "
func CreateUpdateForWidget(script_name string, cmdline string) *WidgetInfo {
    wi := CreateWidgetInfo(script_name)
    parseCmdLine(script_name, cmdline, wi)
    // Note, we place this line after parsing.
    wi.WidgetType = "WidgetUpdate"
    return wi
}


