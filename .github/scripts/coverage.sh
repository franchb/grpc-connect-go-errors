#!/bin/bash

# ANSI color codes
BOLD="\033[1m"
RED="\033[31m"
RESET="\033[0m"

if [ $# -lt 2 ]
then
    echo -e "${BOLD}Usage: coverage.sh [threshold] [go-coverage-report]${RESET}"
    exit 1
fi

threshold=$1
report=$2

output=$(go tool cover -func "${report}")
percent_coverage_string=$(echo "$output" | tail -n 1 | awk '{print $(NF)}')
percent_coverage=${percent_coverage_string::-1}

echo -e "${BOLD}Coverage: ${percent_coverage}${RESET}"

if (( $(echo "$percent_coverage < $threshold" | bc -l) ))
then
    echo -e "${BOLD}${RED}Coverage below threshold of $threshold${RESET}"
    exit 1
fi