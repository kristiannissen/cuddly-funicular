This is for fun
https://gobyexample.com/

### Docker file for GoLang

FROM ubuntu:latest
WORKDIR /root
RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get -y install vim
RUN apt-get -y install wget

COPY .vimrc /root

RUN wget -q https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:{$PATH}"

### .vimrc

syntax on
set number
set tabstop=2
set autoindent
set smartindent
set expandtab
set softtabstop=2
