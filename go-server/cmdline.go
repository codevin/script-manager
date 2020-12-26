package main

import (
  // "fmt"
  "strings"
  "bufio"
  "os"
  "strconv"
  "path/filepath"
  // "io"
)

/*
  Algo: 
   Web server has started the script. Script is using this app (cmdline).
   During this time, stdin, stdout are in cmdline's control. Web app can't
   know that easily.

   So we follow a protocol. As long as we don't receive ":" as first char of 
   line, it is going to be input for cmdline script. As soon as we receive ":",    we stop reading the stdin.
*/

func log_output(s string) {
    // os.Stderr.WriteString("[cmdline]" + s + "\n")
}

func cmdline_output(s string) {
    os.Stderr.WriteString(s + "\n")
}

func return_to_script(s string) {
    // log_output("[ReturnValue For script] " + s)
    os.Stdout.WriteString(s)
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

func get_words(s string) []string {

   // Split by spaces, but not if we are within a quote.
   inQuote := false
   // single word char: false
   // space: true
   // quote: true, but within quote, ignore.
   var quotedStrings = func(c rune) bool {
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
   return MyFieldsFunc(s, quotedStrings)
}


func stdin_is_tty() bool {
   stat, _ := os.Stdin.Stat()
   if (stat.Mode() & os.ModeCharDevice) == 0 {
      return false 
   } else {
      return true 
   }
}
func get_exit_value(line string) (int, bool) {
        arr := get_words(line)
        exitValue := 0
        if len(arr) > 0 && arr[0] == ":exit" {
           if ( len(arr) > 1 ) {
              val, err := strconv.Atoi(arr[1])
              if err != nil {
                 exitValue = -1
              } else {
                 exitValue = val
              }
           }
           return exitValue, true
        }
        return 0, false
}

func scan_lines(lines *[]string, s string) bool {
     scanner := bufio.NewScanner(strings.NewReader(s))
     scanner.Split(bufio.ScanLines)
     n:=0
     for scanner.Scan() {
         line := scanner.Text()
         if ( line != "" ) {
             *lines = append(*lines, line)
             if (line[0] == ':') {
                return true
             }
         }
         n++
     }
     return false 
}

// Consume till we get back ":", or there is force close of input channel.  
func receive_from_webserver() ([]string, bool) {

     readBuf := make([]byte, 1024)
     lines := make([]string, 0, 10)
     for {
         n, err := os.Stdin.Read(readBuf)
         if ( n > 0 ) {
             done := scan_lines(&lines, string(readBuf[0:n-1]))
             if ( done ) {
                return lines, true
                break
             }
         } else if ( err != nil ) {
             // Say EOF
             return lines, false
         }
     }
     return lines, false
}

func get_lines_NOTUSED() ([]string, bool) {
        var readBuf []byte
        readBuf = make([]byte, 1024)

        lines := make([]string, 0, 20)
        ok:=false
        for {
            n, err := os.Stdin.Read(readBuf)
            if ( n == 0 && err == nil) {
                  continue;
            }
            if (err != nil || n < 0) {
                if (err != nil) {
                    log_output("Error reading from web")
                    log_output(err.Error())
                }
                ok = ( err.Error() == "EOF")
                break
            }
            text := string(readBuf)[0:n-1]

            if (text[0] == ':') {
                lines = append(lines, text)
                ok = true
                break
            } else {
                lines = append(lines, text)
            }
        }
        return lines, ok
}

/*
   Algorithm:

   cmdline is called from running script. Its purpose is to interact with
   web server, which in turn interacts with user. It gets back the results and
   passes them back to script using two means:
       - Controlling Exit code 
       - Producing stdout output, which is collected by main script (back ticks).

   1. Converts command line arguments into stdout/stderr string that starts with ":cmdline", sends to webserver.
   2. Wait for response from webserver on stdin. 
       - If it doesn't start with ":exit", then output will be collected and sent to main program as stderr. 
       - Final line should be with ":exit <exitcode>"
*/
func main() {

        log_output("UI request follows.")

        thisProg := filepath.Base(os.Args[0])
        output := ":cmdline"
        if (thisProg == "update")  {
           output = ":update"
        }

        n := len(os.Args[1:])
        if ( n >= 0 ) {
            for _, a := range os.Args[1:] {
               output = output + ` "` + a + `"`
            }
            cmdline_output(output)
        } else {
            // Nothing to do. Exit?
        }

        // Response sent on stdin
        lines, ok := receive_from_webserver()
        // lines, ok := get_lines()
        if (! ok) {
            log_output("Bad Response. Exiting with -1.")
            os.Exit(-1)
        }

        var no_of_lines = len(lines)

        log_output("No. of lines read=" + strconv.Itoa(no_of_lines))

        final_output := ""
        exitValue := -1
        if ( no_of_lines > 0 ) {
           for _, line := range(lines)  {
               if (line[0] == ':') {
                   val, found := get_exit_value(line) // last line
                   if found {
                      exitValue = val
                      break
                   }
               }
               final_output += line
               final_output += "\n"
           }
        }
        log_output("Normal Response. Exiting with value=" + strconv.Itoa(exitValue))
        return_to_script(final_output)
        os.Exit(exitValue)
}

