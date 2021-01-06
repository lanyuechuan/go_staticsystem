FROM centos:latest
MAINTAINER lan
WORKDIR /root
ADD ./test01 ./test01
ENTRYPOINT ["./test01"]
