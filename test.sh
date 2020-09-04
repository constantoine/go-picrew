prev=;
for i in out/*.png;
do
	if [ -n "$prev" ]
	then
        	diff "$prev" "$i"
	fi
	prev=$i
done
