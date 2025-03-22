FROM ubuntu:latest


ENV DEBIAN_FRONTEND=noninteractive

RUN apt  update && apt install -y golang-go 

WORKDIR /app

COPY . .

RUN go mod init myapp && go mod tidy && go build -o myapp

EXPOSE 8081

CMD ["./myapp"]
