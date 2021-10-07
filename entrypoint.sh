#!/bin/sh -l
echo "$1"
eval "/annotator -i $1 -o $2"