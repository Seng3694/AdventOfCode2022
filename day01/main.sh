#!/bin/sh
paste -sd '+' input.txt | sed -r 's/\+{2}/\n/g' | bc | sort -n | tail -n 1
paste -sd '+' input.txt | sed -r 's/\+{2}/\n/g' | bc | sort -n | tail -n 3 | paste -sd '+' | bc
