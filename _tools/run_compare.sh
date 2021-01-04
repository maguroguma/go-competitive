#!/bin/bash

# run_compare.sh {testcase dir} {exe1} {exe2}

# help
if [ $# = 0 ] ; then
  echo '[Usage] run_compare.sh {testcase dir} {exe1} {exe2}'
  exit 0
fi

# target test cases
files=$(find "$1" -type f | xargs)
if [ -z "$files" ] ; then
  echo '[Error] no random test cases'
  exit 1
fi

# run tests
ac=$((0))
wa=$((0))
total=$((0))
for inputFile in $files ; do
  total=$((total + 1))
  echo 'Test: '"$inputFile"
  res1=$(go run $2 < $inputFile)
  res2=$(go run $3 < $inputFile)
  resDiff=$(diff <(echo $res1) <(echo $res2))
  if [ $? == 0 ] ; then
    ac=$((ac + 1))
    echo 'AC! Answer:'
    echo "$res1"
  else
    wa=$((wa + 1))
    echo 'Wrong Answer... Difference:'
    echo "$resDiff"
  fi
  echo '=== === === === === === === === === === === === '
done

# testing result
echo "AC: $ac/$total, WA: $wa/$total"
if [ $ac = $total ] ; then
  echo 'Success!'
else
  echo 'Failed...'
fi

