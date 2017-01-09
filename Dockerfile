FROM mysql:5.5.47

RUN sed -e 's/^ *user *= *mysql$/user = root/' -i /etc/mysql/my.cnf