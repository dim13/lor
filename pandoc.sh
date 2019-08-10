#!/bin/sh

mkdir md

cd dump

for i in *.wiki; do
	pandoc -f mediawiki -t gfm -s $i -o ../md/${i%.*}.md
done
