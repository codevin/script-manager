#!/bin/bash

if [ x$1 == x ]
then
  echo Usage: $0 \<vue_component.vue\>
  exit
fi

code=`awk '/<\/script>/ {exit} /<script>/ {f=1; next} f' $1` 
# > /tmp/vue-tmp.js && node /tmp/vue-tmp.js

echo "var xyz=Function('exports', 'require', 'httpVueLoader', 'module', $code); console.log(xyz)" | node
