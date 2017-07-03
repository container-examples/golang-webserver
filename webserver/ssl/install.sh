#!/bin/bash
openssl req -nodes -new -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -subj "/C=FR/ST=IDF/L=Paris/O=My Inc/OU=DevOps/CN=localhost/emailAddress=example@gmail.com"
