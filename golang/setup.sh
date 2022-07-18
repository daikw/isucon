#!/bin/bash

make
chown isucon app

systemctl reload nginx
mv /var/log/nginx/access.log /var/log/nginx/access.log.$(date +%T)
nginx -s reopen

systemctl restart mysql
rm /var/log/mysql/mysql-slow.log
mysqladmin flush-logs

systemctl restart isu-go

echo "Application restart done, run your benchmarker ..."

sleep 5
sudo query-digester

ls /tmp/slow_query_*.digest | xargs -IXXX sudo mv XXX /home/isucon/digests/$1
