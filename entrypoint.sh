#!/bin/sh -l
ls
eval "/annotator -i $1 -o $2 -p '$3'"
