#!/bin/bash

make
chown isucon app

systemctl reload nginx
mv /var/log/nginx/access.log /var/log/nginx/access.log.$(date +%T)
nginx -s reopen

systemctl restart mysql
rm /var/log/mysql/mysql-slow.log
mysqladmin flush-logs

sleep 5
sudo query-digester

ls /tmp/slow_query_*.digest /home/isucon/digests/$1
