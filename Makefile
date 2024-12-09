deploy:
	@cp -r /home/isucon/webapp/etc/nginx /etc/
	@cp -r /home/isucon/webapp/etc/mysql /etc/
	@systemctl restart nginx
	@systemctl restart mysql
	@systemctl daemon-reload

backup:
	@cp -r /etc/mysql /home/isucon/webapp/etc
	@cp -r /etc/nginx /home/isucon/webapp/etc
	@chown isucon:isucon -R .
