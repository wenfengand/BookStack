FROM truthhun/bookstack:env
MAINTAINER "wenfengand <wenfengand@gmail.com>"


ENV LANG en_US.utf8

# 将程序拷贝进去
COPY . /www/BookStack/

# 将程序拷贝进去
COPY ./lib/time/zoneinfo.zip /usr/local/go/lib/time/

RUN chmod 0777 -R /www/BookStack/

WORKDIR /www/BookStack/

CMD [ "./run.sh" ]
