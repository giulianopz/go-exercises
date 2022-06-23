#!/bin/sh

podman run -d --name=mysql -e MYSQL_ROOT_PASSWORD=pwd -v $(pwd):/opt/db:Z -p=3306:3306 mysql
podman exec -it mysql /bin/sh -c "mysql -u root -ppwd </opt/db/initdb.sql"
