FROM golang:latest AS builder

WORKDIR /app
COPY api/index.go /app/api/index.go
COPY build.go /app/build.go
COPY assets/ /app/assets/
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

# Forgive me for this, vercel forced us to stick with a code structure
# so as a workaround i did this ;-;
RUN sed -i \
  -e 's/package handler/package main/' \
  -e '/var app \*gin.Engine/d' \
  -e '/func init() {/,/^}/d' \
  -e '/func Handler(/,/^}/d' \
  api/index.go && \
  printf '\nfunc main() {\n  gin.SetMode(gin.ReleaseMode)\n  app := gin.New()\n  r := app.Group("/api")\n  iconRoute(r)\n\n  decoder := json.NewDecoder(strings.NewReader(iconsJSON))\n  if err := decoder.Decode(&icons); err != nil {\n    panic(err)\n  }\n\n  for key := range icons {\n    iconNameList = append(iconNameList, strings.Split(key, "-")[0])\n    if strings.Contains(key, "-light") || strings.Contains(key, "-dark") || strings.Contains(key, "-auto") {\n      themedIcons = append(themedIcons, strings.Split(key, "-")[0])\n    }\n  }\n\n  app.Run()\n}\n' >> api/index.go

RUN go mod tidy
RUN go run build.go

# i swear there is a command for more optimitzation but i cant remember aaghhhhhhh
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -ldflags="-s -w" -o skillicons ./api/index.go

FROM scratch

COPY --from=builder /app/skillicons /skillicons

EXPOSE 8080

ENTRYPOINT ["/skillicons"]