#!/bin/bash
cd main
go run . > /dev/null &
cd ..
python show.py > /tmp/board.svg
