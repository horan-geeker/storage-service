FROM scratch
ADD PRC /etc/localtime
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/
ADD ./storage-service /app
ADD ./.env /.env
EXPOSE 80
CMD ["/app"]