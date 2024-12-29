#!/usr/bin/env bash

readme=$(grep -c '\/description\/' README.md)
revise=$(wc -l todos/need-revise.md | awk '{print $1}')
solve=$(grep -c '\/description\/' todos/need-solve.md)

total=$((readme + revise + solve))

echo "Local: $total, Remote: https://leetcode.com/u/therenotomorrow/"
