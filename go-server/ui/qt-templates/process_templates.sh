#!/bin/bash

set -e
echo "go generate compile_template ... "
go generate compile_templates.go || {
   echo "Had errors. Not deploying."
   exit 
}

echo "go run compile_template ... "
go run compile_templates.go || {
   echo "Had errors. Not deploying."
   exit 
}

qtc -dir=./widgets
 
echo "Deploying ..."
set -x
cp ./out/index-autogen.html ../../http-static/index.html 
cp ./views/index-cleaned.js ../../http-static/_assets/js/index-cleaned.js
cp ./out/index-autogen.js ../../http-static/_assets/js/index-autogen.js
cp ./out/index-autogen.css ../../http-static/_assets/css/index-autogen.css
cp ./widgets/*.go ../widgets


echo "Deploying server vue components .."
mkdir -p ../../http-static/vue/components
cp ./vue-server-components/* ../../http-static/vue/components
