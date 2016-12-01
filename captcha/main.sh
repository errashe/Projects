convert img.jpg  -format %c -depth 8  histogram:info:histogram_image.txt
String_color=$(sort -n histogram_image.txt | tail -1)
[[ $String_color =~ (.[0-9A-zА-я]{6}) ]] && Code_color=${BASH_REMATCH[0]}
rm -rf histogram_image.txt

convert img.jpg -fill black -fuzz 80% +opaque $Code_color img2.jpg
convert img2.jpg -fill white -fuzz 20% +opaque "#000000" img3.jpg

# convert img3.jpg -blur 0x2 -level 40%,80% img4.jpg

convert img3.jpg -fill white -draw 'color 40,40 floodfill' img4.jpg