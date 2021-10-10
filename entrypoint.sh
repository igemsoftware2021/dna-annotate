#!/bin/sh -l
echo "Hello World"
eval "/annotator -i $1 -o $2 -p '$3'"
