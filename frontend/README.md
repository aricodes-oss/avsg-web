# frontend

This is the static assets that compose the frontend portion of `avsg`

## Requirements

### Docker

Have a modern version of [docker](https://docs.docker.com/engine/install/) installed and running on your system.

### Native

We're using [bun](https://bun.sh) as our package manager and [vite](https://vitejs.dev) as our bundler.

## Starting the Dev Server

### Docker

Run `docker compose up --build` in the root directory of the project.

### Native

Run `bun install` to download dependencies, then `bun dev` to start the dev server.
Make sure you've configured the port number (default is `8082`) and provided that number to the backend so it can proxy requests in development.

## Building for production

### Docker

Covered in the root README.

### Native

Run `bun run build`
