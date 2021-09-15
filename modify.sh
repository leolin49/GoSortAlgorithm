#!/bin/bash

grepStr="package sort"
sedStr="package sortalgorithm"

filelist=$(ls *.go)

for file in ${filelist}
do
	echo "modify "$file
	#grep ${grepStr} -rl ${file} | xargs -r sed -i 's#'${grepStr}'#'${sedStr}'#g'
	sed -i 's#'${grepStr}'#'${sedStr}'#g' $file
done
