# Cài đặt database 
```
docker run \
--name postgres-17 \
--rm -e POSTGRES_USER=microlap \
-e POSTGRES_PASSWORD=123456 \
-p 5432:5432 -it \
-d postgres:17  
```

Windows
```
 docker run --name postgres-17 --rm -e POSTGRES_USER=microlap -e POSTGRES_PASSWORD=123456 -p 5432:5432 -it -d postgres:17
```