#!/bin/bash
NUM=$(head -n 179 ./streets/Buckingham_Place | tail -n 1 | awk '{print $3}' | cut -c 2-)
echo "$NUM"
cat ./interviews/interview-"$NUM"
echo "$MAIN_SUSPECT"