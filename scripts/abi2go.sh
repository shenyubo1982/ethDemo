#!/bin/zsh

# 说明概述：将abi转换成go
# 作者：沈宇波
# 创建时间：2022.06.01
# 当前版本：v1.0
# 版本：v1.0
# 版本内容:可以编译abi文件生成go
# Todo: 将abi文件加下所有的abi编译成go（不用参数传递）
# 


cd ../abi
abiname=$1
# echo $abiname
packagename=${(L)abiname} 
# echo $packagenam
abigen --abi=$abiname.abi --pkg=$packagename --out=$abiname.go
echo "Done!"

