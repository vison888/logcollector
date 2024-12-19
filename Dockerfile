FROM alpine:3.14
RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
WORKDIR /app
COPY  logcollector /app/logcollector

RUN mkdir -p temp
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
ENTRYPOINT ["./logcollector"]
