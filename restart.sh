#!/bin/sh

systemctl reload nginx
mv /var/log/nginx/access.log /var/log/nginx/access.log.$(date +%T)
nginx -s reopen

systemctl restart mysql
rm /var/log/mysql/mysql-slow.log
mysqladmin flush-logs
