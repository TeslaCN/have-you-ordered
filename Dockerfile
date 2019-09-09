#FROM alpine:latest
# musl-libc 无法正常运行本应用
#FROM frolvlad/alpine-glibc:latest
# glibc 版本的 alpine不知道会不会有问题
FROM centos:latest
# 用 centos 比较稳
WORKDIR /
ENV PARAMS=""
# -e PARAMS="-config /path/to/config.json"
ADD orderserver /
ADD configs /configs
EXPOSE 60001
ENTRYPOINT ["sh", "-c", "/orderserver $PARAMS"]
