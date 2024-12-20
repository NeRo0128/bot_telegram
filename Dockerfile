FROM golang:alpine AS builder

LABEL authors="Rodolfo"

ENTRYPOINT ["top", "-b"]