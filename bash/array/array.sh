#!/bash/bin

array=(
	11
	12
	13
	14
	15
)

for val in ${array[@]}
do
	echo $val
done

for i in ${!array[@]}
do
	echo $i
done
