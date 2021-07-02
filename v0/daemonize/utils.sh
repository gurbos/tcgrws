#! /bin/bash

function FindStaticDataDir() {
    echo `df -T | grep efs | tr -s ' ' | rev | cut -f1 -d' ' | rev`
}
