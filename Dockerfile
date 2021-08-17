FROM centos
RUN mkdir -p /etc/webapp
COPY ./mytemplate/index.html /etc/webapp/index.html
COPY webserver /home/webserver
RUN ln -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
CMD /home/webserver
