FROM golang:1.13.4-buster

RUN apt-get update && \
    apt-get install -y \
        libgirepository1.0-dev \
        libatk1.0-dev \
        libcairo2-dev \
        libglib2.0-dev \
        libgtk-3-dev \
        libgtksourceview-3.0-dev \
        libpango1.0-dev \
        libwebkit2gtk-4.0-dev

CMD cd /gobbi/internal && go run cmd/generate/*.go
