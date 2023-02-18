#!/bin/bash

# 一次性模块打包脚本
current=$(pwd)
output="$current/modules"

i=0
dirs=$(ls -1 | grep module)
for line in $dirs
do
  names[${i}]=$line
  ((i++))
done
echo "match dirs: ${#names[*]}"

echo "crate output dir [modules]"
mkdir -p "$output"

for dir in "${names[@]}"
do
  echo "build $dir"
  cd "$current/$dir" || exit
  go build -mod=mod -buildmode=plugin

  opt=$(ls | grep ".so")
  if [ ! -z "$opt" ];then
    echo "build output: $opt"
    echo "move $opt -> $output"
    mv "$opt" "$output"
  fi
done