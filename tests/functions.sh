#!/bin/bash

function retry {
  local n=0
  local max=20
  local delay=5
  while true; do
    ((n=n+1))
    "$@" && break || {
      echo "Command failed. Attempt $n/$max."
      if [[ $n -ge $max ]]; then
        return 1
      fi
      sleep $delay;
    }
  done
}
