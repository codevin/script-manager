package main

import (
        "fmt"
        "./views"
        "io/ioutil"
)

/* go:generate go get -u github.com/valyala/quicktemplate/qtc */

//go:generate qtc -dir=./views 

type Page struct {
   Name string
   Title string
   Fn   func(s string)string
}

func Pages() []Page {
    return  []Page{ {"index", "Welcome to GCare", views.HtmlForIndexPage} }
}

func Javascripts() []Page {
    return  []Page{ {"index", "Welcome to GCare", views.JSCodeForIndexPage} }
}

func CSS() []Page {
    return  []Page{ {"index", "Welcome to GCare", views.CSSForIndexPage} }
}

func createIndexFile(page *Page) {
     output := []byte( page.Fn(page.Title) )
     ioutil.WriteFile("./out/" + page.Name + "-autogen.html", output, 0644) 
     fmt.Println("Wrote to: ", "./out/" + page.Name + "-autogen.html")
}

func createJavascriptFile(page *Page) {
     output := []byte( page.Fn(page.Title) )
     ioutil.WriteFile("./out/" + page.Name + "-autogen.js", output, 0644) 
     fmt.Println("Wrote to: ", "./out/" + page.Name + "-autogen.js")
}

func createCSSFile(page *Page) {
     output := []byte( page.Fn(page.Title) )
     ioutil.WriteFile("./out/" + page.Name + "-autogen.css", output, 0644) 
     fmt.Println("Wrote to: ", "./out/" + page.Name + "-autogen.css")
}



func main() {
    for _, page:=range(Pages()) {
       createIndexFile(&page)
    }
    for _, page:=range(Javascripts()) {
       createJavascriptFile(&page)
    }
    for _, page:=range(CSS()) {
       createCSSFile(&page)
    }
}

