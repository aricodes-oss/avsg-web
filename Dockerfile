## Frontend build
FROM oven/bun:latest as frontend

WORKDIR /code
COPY frontend .

RUN bun install
RUN bun run build


## Backend build
FROM golang:latest as backend

WORKDIR /code

# Take the frontend slug and copy it into this image
COPY --from=frontend /embeds/dist ./embeds/dist

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN go build

## Ship!
FROM golang:latest

COPY --from=backend /code/avsg .

ENV GIN_MODE=release

CMD ["./avsg"]
