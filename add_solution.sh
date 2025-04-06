#!/bin/bash

# Get the problem number and name from user input
read -p "Enter problem number: " num
read -p "Enter problem name (use-dash): " name

folder=$(printf "%04d-%s" $num $name)
mkdir "$folder"
touch "$folder/main.go"
touch "$folder/README.md"

echo "## LeetCode $num: ${name//-/ }" > "$folder/README.md"
echo "// Solution for LeetCode $num: ${name//-/ }" > "$folder/main.go"
echo "Created folder: $folder"
