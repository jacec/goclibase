#!/bin/sh -e

#check that godep is installed, if not, do it!
command -v godep >/dev/null 2>&1 || {
  echo "godep is required, but not avilable.\nPlease install it and make sure it's on your PATH.\nfor more information please see https://github.com/tools/godep" >&2;
  exit 1;
}
