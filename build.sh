#!/bin/bash
work_dir=$(pwd)
data_dir=$(pwd)/data
echo wokr dir is $work_dir
echo data dir is $data_dir

rm -rf ${data_dir}
git clone https://github.com/yishui01/chinese-poetry-Mysql-Elastic.git data

cd ${data_dir}
git clone https://github.com/chinese-poetry/chinese-poetry.git chinese-poetry
composer install
php ./Worker.php

cd ${work_dir}
rm ${work_dir}/poems*.db
go mod download
go run convert.go
file=poems.db
newfile=poems_$(openssl sha1 $file | awk '{print $2}').db
cp $file $newfile
