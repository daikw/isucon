# installation

```
## common
sudo apt install -y git unzip graphviz tree
sudo apt install -y percona-toolkit apache2-utils dstat

## alp
wget https://github.com/tkuchiki/alp/releases/download/v1.0.3/alp_linux_amd64.zip
unzip alp_linux_amd64.zip ; sudo mv alp /usr/local/bin/alp ; rm alp_linux_amd64.zip

## query-digester

git clone git@github.com:kazeburo/query-digester.git
cd query-digester
sudo install query-digester /usr/local/bin

## redis
sudo add-apt-repository -y ppa:chris-lea/redis-server ; sudo apt install -y redis-server
```

------------------------------------------------------------------------------------------------------------------

# tools

- ab
- pt-query-digest
- alp 

------------------------------------------------------------------------------------------------------------------

# snippets

## ab

```
ab -c 1 -n 10 http://localhost/
```

## bench

```
sudo su - isucon
/home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://172.31.44.240
```

## mysql

```
mysql -uroot -p -e \
  "set global slow_query_log_file = '/tmp/slow-query.log';\
   set global long_query_time = 0;\
   set global slow_query_log = ON;"
```


```
sudo mysql -uroot

explain $query

show create table $table

alter table $table add index $idx_name($column)
```

## logrotate

```sh
sudo logrotate -f /etc/logrotate.conf
sudo mv /var/log/nginx/access.log{,.$(date +%T)}


sudo pt-query-digest /var/log/mysql/mysql-slow.log > digest_query.log.$(date +%T)
sudo rm /var/log/mysql/mysql-slow.log
```

## analyze

```
sudo cat /var/log/nginx/access.log | alp json -r -m "/image/+,/posts/+"
```

```
sudo pt-query-digest /tmp/slow-query.log
```

## nginx

```
sudo nginx -t
```


## restart & snapshot-logs

```sh
#!/bin/sh

systemctl reload nginx
mv /var/log/nginx/access.log /var/log/nginx/access.log.$(date +%T)
nginx -s reopen
```


Recommended to replace with `query-digester`

```sh
pt-query-digest /var/log/mysql/mysql-slow.log > digest_query.log.$(date +%T)
rm /var/log/mysql/mysql-slow.log
systemctl restart mysql
mysqladmin flush-logs
```

