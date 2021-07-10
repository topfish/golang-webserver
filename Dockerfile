FROM centos
COPY webserver /home/webserver
RUN ln -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
CMD /home/webserver