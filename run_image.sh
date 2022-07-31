# 自动删除容器的--rm自然也是用不了
#docker run --name goback -p 8888:8888 aikan/golang:0.1
#docker run --link mysql --name goback -p 8888:8888 aikan/golang:0.1
#docker run --name goback2 -p 8888:8888 aikan/golang:0.1
docker run  -tid --name goback -p 18888:8888 fogoo/aikan:1.2
