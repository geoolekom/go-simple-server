screen -S postgres -dm sudo -u postgres /usr/lib/postgresql/9.6/bin/postgres -D /var/lib/postgresql/9.6/main -c config_file=/etc/postgresql/9.6/main/postgresql.conf
sleep 2
go-simple-server
