#!/bin/bash

# root_dir=$1
# link_target_dir=$2
root_dir="$HOME/go/src/github.com/myokoyama0712/go-competitive/past_contests/AGC"
link_target_dir="$HOME/go/src/github.com/myokoyama0712/go-competitive/remark_link/AGC"

# 末尾のスラッシュ除去
# root_dir=$(echo $1 | tr -d '/')
# link_target_dir=$(echo $2 | tr -d '/')

# 最低限の引数チェック
if [[ !( -d $root_dir && -d $link_target_dir )  ]]; then
  echo 'arguments error!'
  exit 1
fi

files_array=$(find $root_dir -type f)
for file in $files_array; do
  path=$file
  declare -a a_dirs=()
  a_dirs=$(echo $path | tr '/' ' ') # /をスペースに置換

  for dir in ${a_dirs[0]}
  do
    if [ $dir = "README.md" ]; then
      # echo "SLINK FILE PATH: ${link_target_dir}/${before_dir}-${dir}"
      # echo "ORIGIN FULL PATH: $file"
      ln -fnsv "$file" "${link_target_dir}/${before_dir}-${dir}"
    fi

    before_dir=$dir
  done
done

exit 0
