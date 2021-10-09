#!/bin/sh -l
ls input/
eval "/annotator -i $1 -o $2 -p '$3'"
