FROM golang:1.15

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build

RUN chmod +x start.sh

EXPOSE 8001

CMD ["./start.sh"]
