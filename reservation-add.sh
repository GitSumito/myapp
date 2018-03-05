#!/bin/bash
ymd=$1
yyyy=${ymd:0:4}
mm=${ymd:4:2}
dd=${ymd:6:2}

:> /tmp/$ymd.sql

for ii in `seq 0 23`;do
 for i in "00" "30" ;do
	echo "INSERT INTO \`reservation\` VALUES ('${yyyy}-${mm}-${dd} ${ii}:${i}:00',1,0,0,0,0,0,0,now(),now());" >> /tmp/$ymd.sql
 done
done


mysql -uroot  -h127.0.0.1 -ppassword booking < /tmp/$ymd.sql
