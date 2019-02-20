#!/bin/bash

link_target_dir="$HOME/go/src/github.com/myokoyama0712/go-competitive/remark_link/ABC"

link_files_array=$(find $link_target_dir -type l | grep "README.md")
for file in $link_files_array
do
  unlink $file
done

