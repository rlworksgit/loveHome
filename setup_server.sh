redis-server ./conf/redis.conf
#启动trackerd  restart表示重新启动
fdfs_trackerd ./conf/tracker.conf restart
#启动storaged
fdfs_storaged ./conf/storage.conf restart
