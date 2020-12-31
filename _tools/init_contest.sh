#!/bin/bash

# 名前
#   init_contest.sh - 引数で渡した問題コード名で初期ファイルを作る
#
# 書式（コンテスト用ディレクトリに移動してから実行）
#   $GOCOMPE/init_contest.sh a b c d e f

readonly SCRIPT_NAME=${0##*/}

# make files for each problem directory
for dirName in "$@"
do
  mkdir $dirName
  mkdir "$dirName/random"
  cp "${GOCOMPE}/template/go/base-competitive.go" "${dirName}/${dirName}.go"
done

# make files for current directory
touch 'README.md'
cp "${GOCOMPE}/_tools/generate.go" "${PWD}/"
