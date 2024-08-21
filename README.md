# Axiom Verge Save Game Utilities

This is a web application that allows users to encrypt and decrypt their Axiom Verge save files.

(live URL pending)

## Requirements

### Docker

Have a modern version of [docker](https://docs.docker.com/engine/install/) installed and running on your system.

### Native

**Backend**:
[Go v1.18+](https://go.dev)

[gowatch](https://github.com/silenceper/gowatch) (optional, used for hot reloading in dev)

**Frontend**:
We're using [bun](https://bun.sh) as our package manager + interpreter

## Starting the Dev Server

Regardless of which launch method you use, the application will be available at [http://localhost:8081](http://localhost:8081) when it starts up.

### Docker

```sh
docker compose up --build
```

### Native

Have two terminal windows ready (or be familiar with [Linux job control](https://linuxcommand.org/lc3_lts0100.php)).

In the root of the repository, run `gowatch`. In the `frontend` directory, run `bun run dev`.

## Building for Production

### Docker

`docker build -t avsg-web:latest .`

### Native

1. Go to the frontend folder and run `bun run build` to create the frontend bundle
2. In the root of the repository, run `go build` to create the application executable
3. Deploy the `avsg` binary wherever you please!

When building natively the frontend assets are embedded into the application binary and do not need to be shipped separately.

## FAQ

**Q: Why another save game decryptor?**

**A:** The existing ones work excellently, but require some technical knowledge to compile and run. The hope with this application is that save game editing is more widely accessible.

**Q: Why use a bundler for a plain HTML frontend?**

**A:** Mostly for hot reloading in development and the ability to load versioned assets (namely [Bootstrap](https://getbootstrap.com)) from NPM instead of a CDN. Page load times and application responsiveness are important to me.
