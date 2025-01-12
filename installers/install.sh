#!/usr/bin/env sh
# Copyright 2022 ActiveState Software Inc. All rights reserved.

# URL to fetch installer archive from
BASE_FILE_URL="https://state-tool.s3.amazonaws.com/update/state"
# Name of the installer executable and archive.
INSTALLERNAME="state-installer"
# Channel the installer will target
CHANNEL='release'
# the download exetension
DOWNLOADEXT=".tar.gz"
# the installer extension
BINARYEXT=""

SESSION_TOKEN_VERIFY="{TOKEN""}"
SESSION_TOKEN="{TOKEN}"
SESSION_TOKEN_VALUE=""

if [ "$SESSION_TOKEN" != "$SESSION_TOKEN_VERIFY" ]; then
  SESSION_TOKEN_VALUE=$SESSION_TOKEN
fi

parseChannel() {
  i=0
  for arg in ${@}; do
    i=$((i + 1)) && [ "${arg}" != "-b" ] && continue
    echo "${@}" | cut -d' ' -f$((${i} + 1)) && return
  done
  echo $CHANNEL
}
CHANNEL=$(parseChannel "$@")

if [ -z "${TERM}" ] || [ "${TERM}" = "dumb" ]; then
  OUTPUT_OK=""
  OUTPUT_ERROR=""
  OUTPUT_END=""
else
  OUTPUT_OK=`tput setaf 2`
  OUTPUT_ERROR=`tput setaf 1`
  OUTPUT_END=`tput sgr0`
fi

progress () {
  printf "• %s... " "$1"
}

progress_done() {
  echo "${OUTPUT_OK}✔ Done${OUTPUT_END}"
}

progress_fail() {
  echo "${OUTPUT_ERROR}x Failed${OUTPUT_END}"
}

error () {
  echo "$OUTPUT_ERROR${1}$OUTPUT_END"
}

# Determine the current OS.
case `uname -s` in
Linux)
  OS="linux"
  ;;
*BSD)
  OS=`uname -s | tr '[A-Z]' '[a-z]'`
  error "BSDs not supported yet"
  exit 1
  ;;
Darwin)
  OS="darwin"
  ;;
MINGW*|MSYS*)
  OS="windows"
  DOWNLOADEXT=".zip"
  BINARYEXT=".exe"
  ;;
*)
  error "Unsupported OS: `uname -s`"
  exit 1
  ;;
esac

# Determine a fetch method
if [ ! -z "`command -v wget`" ]; then
  FETCH="wget -q -O"
elif [ ! -z "`command -v curl`" ]; then
  FETCH="curl -sS -o"
else
  error "Either wget or curl is required to download files"
  exit 1
fi

# Determine the tmp directory.
if [ -z "$TMPDIR" ]; then
  TMPDIR="/tmp"
fi

progress "Preparing Installer for State Tool Package Manager"
STATEURL="$BASE_FILE_URL/$CHANNEL/$OS-amd64/$INSTALLERNAME$DOWNLOADEXT"
ARCHIVE="$INSTALLERNAME$DOWNLOADEXT"
if ! $FETCH $TMPDIR/$ARCHIVE $STATEURL ; then
  progress_fail
  error "Could not fetch the State Tool installer at $STATEURL. Please try again."
  exit 1
fi

if [ $OS = "windows" ]; then
  # Work around bug where MSYS produces a path that looks like `C:/temp` rather than `C:\temp`
  TMPDIRW=$(echo $(cd $TMPDIR && pwd -W) | sed 's|/|\\|g')
  powershell -command "& {&'Expand-Archive' -Force '$TMPDIRW\\$ARCHIVVE' '$TMPDIRW'}"
else
  tar -xzf $TMPDIR/$ARCHIVE -C $TMPDIR || exit 1
fi
chmod +x $TMPDIR/$INSTALLERNAME$BINARYEXT
progress_done
echo ""

ACTIVESTATE_SESSION_TOKEN=$SESSION_TOKEN_VALUE $TMPDIR/$INSTALLERNAME$BINARYEXT "$@" --source-installer="install.sh"