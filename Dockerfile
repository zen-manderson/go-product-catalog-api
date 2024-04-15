# Build --------------------------------
FROM us-central1-docker.pkg.dev/zen-dev-166315/zen-docker-images/golang/1.22-bookworm:latest as build

WORKDIR /app

ARG RELEASE_VERSION
ARG REPO_NAME
ARG go_gh_read_all_token

RUN git config --global url."https://${go_gh_read_all_token}@github.com/".insteadOf "https://github.com/"
ENV GOPRIVATE="github.com/zenbusiness/*"

# Copy all files over
COPY . ./

# Log the Go compiler version for reference and verification
RUN go version

RUN go mod download

# Pass the release version and app name to the build so that it can be used in the buildinfo package for logging and metrics
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X github.com/zenbusiness/go-toolkit/buildinfo.version=${RELEASE_VERSION} -X github.com/zenbusiness/go-toolkit/buildinfo.appName=${REPO_NAME}" -o ./server ./main.go

# Create a 2nd state that just copys over the server
FROM us-central1-docker.pkg.dev/zen-dev-166315/zen-docker-images/debian:bookworm-slimlatest  AS application

WORKDIR /app

COPY --from=build /app/server ./

ARG RELEASE_VERSION=unset
ARG RELEASE_BRANCH=unset
ARG RELEASE_TAG=unset

ENV RELEASE_VERSION=$RELEASE_VERSION
ENV RELEASE_BRANCH=$RELEASE_BRANCH
ENV RELEASE_TAG=$RELEASE_TAG

EXPOSE 50051
EXPOSE 3001

# Run
ENTRYPOINT ["/app/server"]
