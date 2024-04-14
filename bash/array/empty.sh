#!/bash/bin

array=(
)

for ((i=0; i<10; i++))
do
	array[$i]=$i

	if [ $i -eq 2 ]; then
		array[$i]=""
	fi
	if [ $i -eq 4 ]; then
		array[$i]=""
	fi
done

for index in ${!array[@]}
do
	echo "$index: ${array[$index]}"
done
