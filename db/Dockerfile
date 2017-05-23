FROM mysql:8.0 

ENV MYSQL_DATABASE=todo
ENV MYSQL_ROOT_PASSWORD=bob

ADD createTodoTable.sql /docker-entrypoint-initdb.d/
EXPOSE 3306 
