# syntax=docker/dockerfile:1

FROM docker.io/alpine:3.20 AS base

RUN apk add --no-cache \
    tzdata sudo bash git aws-cli curl ca-certificates jq \
    && rm -rf /var/cache/apk/* \
    && mkdir -p /var/cache/apk \
    && ln -s /var/cache/apk /etc/apk/cache

ENV TZ="Asia/Singapore" \
    LANG="en_US.UTF-8" \
    LANGUAGE="en_US:en"

USER root
