#!/bin/bash
# [] 中可用的比较运算符只有 == 和 !=，两者都是用于字符串比较的，不可用于整数比较，整数比较只能使用 -eq，-gt 这种形式，[[]]
target="fatal"
for ((i=1;i<=10;i++))
do
	echo "now trying connect num: $i"
	d=$(date)
	echo $d
	result="$(git push origin -f)"
	res=$?
	echo $res
	if [[ $res != *[^0-9]* ]]&&[[ $res != 0* ]];
	then
		echo $res is int
	else
		echo $res is string
	fi
	if [[ $res == 0 ]];
	then
		echo "break!" && break
	else
		echo -e "not yet submitted, continue~\n"
	fi
	#echo "***"
	#echo $result
	#echo "---"
	#[[ "$result" != *"$target"* ]] && break
	#echo $target
	#if [[ $result =~ $target ]];
	#then
	#	echo "not yet submitted, continue~"
	#else
	#	echo "break!" && break
	#fi

done
