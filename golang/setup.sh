#!/bin/bash

# isu-go
make
chown isucon app
systemctl restart isu-go

# nginx
systemctl reload nginx
mv /var/log/nginx/access.log /var/log/nginx/access.log.$(date +%T)
nginx -s reopen

# mysql
systemctl restart mysql
rm /var/log/mysql/mysql-slow.log
mysqladmin flush-logs

# redis
# ...

echo "Application restart done, run your benchmarker ..."

sleep 5
sudo query-digester

filename=/home/isucon/digests/$1
ls /tmp/slow_query_*.digest | xargs -IXXX sudo mv XXX $filename
echo "Query digest saved as: " $filename
