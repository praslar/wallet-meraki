FROM alpine:3.8

WORKDIR /home/
COPY ../wallet-meraki.bin .
RUN chmod +x wallet-meraki.bin

EXPOSE 8081
CMD ["/wallet-meraki.bin"]
