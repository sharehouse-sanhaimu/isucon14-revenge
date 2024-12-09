deploy:
	@cp -r /home/isucon/webapp/etc/nginx /etc/
	@cp -r /home/isucon/webapp/etc/mysql /etc/
	@systemctl restart nginx
	@systemctl restart mysql
	@systemctl daemon-reload
