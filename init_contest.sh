#!/bin/bash

# 名前
#   init_contest.sh - 引数で渡した問題コード名で初期ファイルを作る
#
# 書式（コンテスト用ディレクトリに移動してから実行）
#   $GOCOMPE/init_contest.sh a b c d e f

readonly SCRIPT_NAME=${0##*/}

for dirName in "$@"
do
  mkdir $dirName
  cp "${GOCOMPE}/_TEMPLATE.go" "${dirName}/${dirName}.go"
done

touch 'README.md'

