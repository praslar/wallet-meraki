FROM alpine:3.8

WORKDIR /home/
COPY wallet.bin .
RUN chmod +x wallet.bin

EXPOSE 8080
CMD ["./wallet.bin"]