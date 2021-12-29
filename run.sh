#!/bin/bash

docker run -e MYSQL_HOST=192.168.1.15 -e MYSQL_USER=root -e MYSQL_PASSWORD=473550 -e MYSQL_DBNAME=devcode -p 8090:3030 alfiantech/devcode-todo:1.0.19