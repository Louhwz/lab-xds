FROM envoyproxy/envoy-alpine-dev:latest

RUN mkdir /code
ADD ./http_server /code
ADD ./lightest /code
ADD ./hello-start.sh /usr/local/bin/hello-start.sh
RUN chmod u+x /code/http_server
RUN chmod u+x /code/lightest
RUN chmod u+x /usr/local/bin/hello-start.sh

ENTRYPOINT /usr/local/bin/hello-start.sh
