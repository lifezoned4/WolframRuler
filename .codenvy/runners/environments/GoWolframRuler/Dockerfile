FROM codenvy/go

ENV GOPATH /home/user/application

RUN mkdir -p /home/user/application/src

ADD $src$ /home/user/app/

WORKDIR /home/user/application/src

# 1. Get dependencies from source code of application
# 2. Start application

CMD go get -d
CMD go run /home/user/app/WolframRulerGo/$executable:-main.go$