FROM golang:1.12.13 
ENV GO111MODULE "on" 
ENV GOPROXY "https://goproxy.cn"
WORKDIR $GOPATH/src/github.com/asynccnu/food_service
COPY . $GOPATH/src/github.com/asynccnu/food_service
RUN make
EXPOSE 8080
CMD ["./main", "-c", "conf/config.yml"]
