#! /bin/bash

function staticDataDir() {
    echo `df -T | grep efs | tr -s ' ' | rev | cut -f1 -d' ' | rev`
}
