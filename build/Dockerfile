FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/cert-operator \
    USER_UID=1001 \
    USER_NAME=cert-operator \
    HOME=/home/cert-operator

# install operator binary
COPY build/_output/bin/cert-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

WORKDIR ${HOME}
ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
