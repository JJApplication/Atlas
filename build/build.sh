#!/bin/bash

# 一次性模块打包脚本
current=$(pwd)
output="$current/modules"

i=0
dirs=$(ls -1 | grep module|grep -v modules)
for line in $dirs
do
  names[${i}]=$line
  ((i++))
done
echo "match dirs: ${#names[*]}"

echo "create output dir [modules]"
rm -rf "$output"
mkdir -p "$output"

for dir in "${names[@]}"
do
  echo "->build $dir"
  echo "cd $current/$dir"
  cd "$current/$dir" || exit
  go build -mod=mod -buildmode=plugin -trimpath

  opt=$(ls | grep ".so")
  if [ ! -z "$opt" ];then
    echo "build output: $opt"
    echo "move $opt -> $output"
    mv "$opt" "$output"
    res=$?
    if [ $res != 0 ];then
      exit 1
    fi
  fi
  echo "===== end ====="
done