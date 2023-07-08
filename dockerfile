FROM alpine:3.8

WORKDIR /home/
COPY wallet_meraki.bin .
RUN chmod +x wallet_meraki.bin

EXPOSE 8081
CMD ["./wallet_meraki.bin"]
