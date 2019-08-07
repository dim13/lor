#!/bin/sh

for i in *.md; do
	#sed -i '' 's/\xd1/х/g' $i
	iconv -c -t utf-8 $i > xxx.md
	mv xxx.md $i
done
