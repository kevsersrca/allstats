FROM golang:1.8.1
MAINTAINER LittleRobot daisukeayanami@gmail.com

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH  

# Install beego & bee
RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee
RUN go get github.com/tools/godep


EXPOSE 8080


ADD / /go/src/allstats

ADD run.sh /

RUN chmod +x /run.sh

CMD ["/run.sh"]





