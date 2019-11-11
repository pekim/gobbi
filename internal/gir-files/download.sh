#!/bin/bash

set -e
set -u

get () {
    VER="bionic"
    MIRRORS="https://packages.ubuntu.com/$VER/amd64/$1-dev/download"

    if [ $# -lt 2 ]; then
        MIRROR="mirrors.kernel.org"
    else
        MIRROR="$2"
    fi

    echo $MIRROR $1

    wget -q --show-progress -O tmp.html "$MIRRORS"
    URL=`cat tmp.html | grep -oP "http://$MIRROR/[^\"]+" | tail -n 1`
    rm tmp.html

    wget -q --show-progress -O tmp.deb "$URL"
    ar x tmp.deb data.tar.xz
    rm tmp.deb
    tar xf data.tar.xz --strip-components 4 ./usr/share/gir-1.0
    rm data.tar.xz
}

cd "$(dirname "$0")"

get libatk1.0
get libgirepository1.0
get libgdk-pixbuf2.0
get libgtk-3
get libgtksourceview-3.0
get libpango1.0 security.ubuntu.com
