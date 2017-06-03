#!/usr/bin/bash
filename="songs.txt"
while read -r line
do
	url=`echo $line | awk -F'|' '{ print $1 }'`
	out=`echo $line | awk -F'|' '{ print $2 }'`
	wget "$url" -O "music/$out"

done < "$filename"