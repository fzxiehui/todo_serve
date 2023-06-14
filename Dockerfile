# Build Stage
FROM fzxiehui/todo_serve:1.13 AS build-stage

LABEL app="build-todo_serve"
LABEL REPO="https://github.com/fzxiehui/todo_serve"

ENV PROJPATH=/go/src/github.com/fzxiehui/todo_serve

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/fzxiehui/todo_serve
WORKDIR /go/src/github.com/fzxiehui/todo_serve

RUN make build-alpine

# Final Stage
FROM fzxiehui/todo_serve

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/fzxiehui/todo_serve"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/todo_serve/bin

WORKDIR /opt/todo_serve/bin

COPY --from=build-stage /go/src/github.com/fzxiehui/todo_serve/bin/todo_serve /opt/todo_serve/bin/
RUN chmod +x /opt/todo_serve/bin/todo_serve

# Create appuser
RUN adduser -D -g '' todo_serve
USER todo_serve

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/todo_serve/bin/todo_serve"]
