# We establish a separate stage for building the app.
FROM golang:1.22-bullseye AS build

WORKDIR /app 

# We optimize our path to discovery, selecting only the files required to install dependencies. 🧭
# With this choice, we unlock the potential of better layer caching, improving our image's efficiency.
COPY go.mod go.sum ./

# Cache mounts speed up the installation of existing dependencies,
# empowering our image to sail smoothly through vast dependency seas.
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

FROM build AS dev

# Setup air and delve, via go install enhances the development
# with hot reload capabilities and powerful debugging prowess
RUN go install github.com/cosmtrek/air@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

CMD ["air", "-c", ".air.toml"]

FROM build AS build-production

# Add non-root user
RUN useradd -u 1001 crocoder

COPY . .

# During this stage, we compile our application ahead of time, avoiding any runtime surprises.
# The resulting binary, web-app-golang, will be our steadfast companion in the final leg of our journey.
# We strategically add flags to statically link our binary.
RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o web-app-golang

# The scratch base image welcomes us as a blank canvas for our prod stage.
FROM scratch

WORKDIR /

# We copy the passwd file, essential for our non-root user
COPY --from=build-production /etc/passwd /etc/passwd

# We transport the binary to our deployable image
COPY --from=build-production /app/web-app-golang web-app-golang
COPY --from=build-production /app/templates templates
COPY --from=build-production /app/static static

# Use non-root user
USER crocoder

# By exposing port 3000, we signal to the Docker environment the intended entry point for our application.
EXPOSE 3000

CMD ["/web-app-golang"]