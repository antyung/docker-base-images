# syntax=docker/dockerfile:1

FROM docker.io/alpine:3.20 AS base

RUN apk add --no-cache \
    tzdata sudo bash git aws-cli curl ca-certificates jq \
    && rm -rf /var/cache/apk/* \
    && cp /usr/share/zoneinfo/Asia/Singapore /etc/localtime \
    && echo "Asia/Singapore" > /etc/timezone \
    && mkdir -p /var/cache/apk \
    && ln -s /var/cache/apk /etc/apk/cache

ENV LANG="en_US.UTF-8" \
    LANGUAGE="en_US:en" \
    LC_ALL="en_US.UTF-8"

USER root
