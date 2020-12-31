#!/bin/bash

# gen_random.sh {testcase dir} {n} {exe}

# help
if [ $# = 0 ] ; then
  echo '[Usage] gen_random.sh {testcase dir} {n} {exe}'
  exit 0
fi

for ((i = 1; i <= $2; i++)) ; do
  echo "[Run] go run $3 > $1/random-$i.in"
  eval "go run $3 > $1/random-$i.in"
done

