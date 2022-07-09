#!/bin/sh

set -e 
set -x

podman run -d --name=mysql -e MYSQL_ROOT_PASSWORD=pwd -v $(pwd):/opt/db:Z -p=3306:3306 mysql

while [[ $(podman inspect mysql -f {{.State.Running}}) != "true" ]]; 
do 
	sleep 0.1; 
done

podman exec -it mysql /bin/sh -c "mysql -u root -ppwd </opt/db/initdb.sql"

export DBUSER=root
export DBPASS=pwd
