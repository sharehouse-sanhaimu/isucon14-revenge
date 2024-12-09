deploy:
	@cp /home/isucon/webapp/env.sh /home/isucon/env.sh
	@cp -r /home/isucon/webapp/etc/nginx /etc/
	@cp -r /home/isucon/webapp/etc/mysql /etc/
	@systemctl restart nginx
	@systemctl restart mysql
	@cd /home/isucon/webapp/go && go build -o isuride
	@systemctl restart isuride-go
	systemctl daemon-reload

backup:
	@cp /home/isucon/env.sh /home/isucon/webapp/env.sh
	@cp -r /etc/mysql /home/isucon/webapp/etc
	@cp -r /etc/nginx /home/isucon/webapp/etc
	@chown isucon:isucon -R .
