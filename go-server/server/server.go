package server

import (
    scr "../script"
    log "github.com/sirupsen/logrus"
    "fmt"
    "time"
    // "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    // "net"
    // "os"
    // "io"
    // "io/ioutil"
    // "path/filepath"
    "strings"
    "encoding/json"
)

var updateTime bool
var systime time.Time

var minversion string = "111"  // Filled by linker.

func CacheControlWrapper(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
        etag := minversion 
        w.Header().Set("Etag", etag)
        h.ServeHTTP(w, r)
    })
}


func handleLogAllRequestsFn(f http.HandlerFunc)  http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
       log.Info("HttpRequest: ", r.URL)
       f(w, r)
   }
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        s:=r.RequestURI
        logMe := true
        switch  {
          case strings.HasPrefix(s, "/vue"):
                logMe = false
          case strings.HasPrefix(s, "/_asset"):
                logMe = false
          case s == "/update":
                logMe = false
          case s == "/interact":
                logMe = false
          case s == "/log":
                logMe = false
        }
        if (logMe) {
           log.Info("Http Request: ", r.RequestURI)
        }
        next.ServeHTTP(w, r)
    })
}

func Begin() {
    updateTime = true

    router := mux.NewRouter()

    router.Use(loggingMiddleware)

    router.HandleFunc("/update", handleUpdateAll)
    router.HandleFunc("/updateScript", handleUpdateScript)
    router.HandleFunc("/interact", handleInteract)
    router.HandleFunc("/logs", handleLogs)

    static_dir := GetConfig().StaticDir
    log.Info("Serving static files from: ", static_dir) 

    // // Static files come here. 
    // fileServer := http.FileServer(http.Dir(static_dir))
    // // Finally, check static files.
    // router.PathPrefix("/").Handler(fileServer)
    router.PathPrefix("/").Handler(
            http.StripPrefix("/", CacheControlWrapper(http.FileServer(http.Dir(static_dir)))))

    portNum := GetConfig().Port
    log.Info("HTTP Server: (Gorilla/router) Starting server at port: ", portNum)

    server := http.Server{
      Handler: router,
      Addr:  ":" + portNum,
      WriteTimeout: (15 * time.Second),
      ReadTimeout: (15 * time.Second),
    }

    err := server.ListenAndServe()
    if ( err != nil ) {
       fmt.Println("Error: Server closed connection with error. Exiting.:", err)
       log.Fatal("Error: Server closed connection with error. Exiting.:", err)
    }
    // err := http.ListenAndServe(":" + strconv.Itoa(portNum), router); 
}

func handleLogs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Connection", "close")
    defer r.Body.Close()

    switch r.Method {
    case "POST":
         err := r.ParseMultipartForm(32 << 20) // Shortcut for 32MB
         if err != nil {
             log.Error("logs: Bad form content:", err)
             http.Error(w, err.Error(), http.StatusInternalServerError)
             return
         }
         script_name := r.FormValue("ScriptName")
         command := ":logs"
         content := ""

         if (script_name == "" || command == "" ) {
             log.Error("logs: interact: received empty args: script_name, command:", script_name, command) 
         }

         log.Info("logs: command=", command, " content=", content)
         response := scr.ProcessMessageFromUI(script_name, command, content) 
         log.Info("logs: response=", response.MessageType) 
         out, err := json.Marshal(response)
         if (err != nil ) {
             log.Error("logs: ERROR in JSON'ing output:", err) 
             out2 := `{"ResultCode": 1}`
             fmt.Fprintln(w, out2)
             return
         }
         fmt.Fprintln(w, string(out))
         // log.Info("logs: response=", string(out))
         return

    default:
        log.Info("GET Request not expected for /script")
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}


func handleInteract(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Connection", "close")
    defer r.Body.Close()

    switch r.Method {
    case "POST":
         err := r.ParseMultipartForm(32 << 20) // Shortcut for 32MB
         if err != nil {
             log.Error("interact: Bad form content:", err)
             http.Error(w, err.Error(), http.StatusInternalServerError)
             return
         }
         script_name := r.FormValue("ScriptName")
         command := r.FormValue("Command")
         content := r.FormValue("Content")

         if (script_name == "" || command == "" ) {
             log.Error("server: interact: received empty args: script_name, command:", script_name, command) 
         }

         log.Info("interact: command=", command, " content=", content)
         response := scr.ProcessMessageFromUI(script_name, command, content)
         log.Info("interact: response=", response.MessageType, " subtype=", response.MessageSubType)
         out, err := json.Marshal(response)
         if (err != nil ) {
             log.Error("interact: ERROR in JSON'ing output:", err) 
             out2 := `{"ResultCode": 1}`
             fmt.Fprintln(w, out2)
             return
         }
         fmt.Fprintln(w, string(out))
         // log.Info("interact: response=", string(out))
         return

    default:
        log.Info("GET Request not expected for /script")
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func handleUpdateAll(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Connection", "close")
    defer r.Body.Close()

    switch r.Method {
    case "GET":
            fallthrough
    case "POST":
            scripts_list, err := json.Marshal(scr.GetConfig().Scripts)
            if (err != nil) {
                log.Error("update: Bad form content:", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            // log.Info(string(scripts_list))
            fmt.Fprintln(w, string(scripts_list))
            return
    default:
        log.Info("GET Request not expected for /script")
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func handleUpdateScript(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Connection", "close")
    defer r.Body.Close()

    switch r.Method {
    case "GET":
            fallthrough
    case "POST":
            p:=scr.ScriptInfo{}
            err := json.NewDecoder(r.Body).Decode(&p)
            if err != nil {
                 http.Error(w, err.Error(), http.StatusBadRequest)
                 return
            }
            script := scr.LookupScript(p.Id)
            if script == nil {
               log.Error("updateScript: script not found: script_name=", p.Id)
               http.Error(w, "Requested Script Not found: " + p.Id, http.StatusInternalServerError)
               return
            }
            json, err := json.Marshal(script)
            if (err != nil) {
                log.Error("update: Json error", err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            fmt.Fprintln(w, string(json))
            return
    default:
        log.Info("GET Request not expected for /script")
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

