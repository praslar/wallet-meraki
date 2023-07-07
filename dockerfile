FROM alpine:3.8

WORKDIR /home/
COPY ../wallet-meraki.exec .
RUN chmod +x wallet-meraki.exec

EXPOSE 8080
CMD ["/wallet-meraki.exec"]
