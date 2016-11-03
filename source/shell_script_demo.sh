#!/bin/sh

for((i=0;i<$1;i++));do
        ps -fe|grep "groupOrderAutoConfirm.php $i$" |grep -v grep
        if [ $? -ne 0 ]
        then
        /usr/local/php56/bin/php /data/httpd/qmfx/\!scripts/groupOrderAutoConfirm.php $i >> /tmp/groupOrderAutoConfirm &
        fi
    done;
