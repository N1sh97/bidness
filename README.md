# bidness
for bro

## Tooling

### Project (not provisional)

- Build tool -> [make](https://www.gnu.org/software/make/)

### Frontend (provisional)

- Programming language -> [typescript](https://www.typescriptlang.org/)
- JS runtime -> [node](https://nodejs.org/en)
- Package manager -> [pnpm](https://pnpm.io/)
- Framework -> [react](https://react.dev/) + [react compiler](https://react.dev/learn/react-compiler)
- Build tool -> [vite](https://vite.dev/)
- Testing -> [vitest](https://vitest.dev/)
- Linting + formatting -> [oxc](https://oxc.rs/#feature-minifier)

### Backend (provisional)

- Programming language -> [go](https://go.dev/)
- Testing -> [testing](https://pkg.go.dev/testing@go1.27.3)
- HTTP server -> [net/http](https://pkg.go.dev/net/http)
- Router -> [chi](https://github.com/go-chi/chi)
- Database -> [sqlite](https://www.sqlite.org/)
- Migrations -> [goose](https://github.com/pressly/goose)
- SQL tooling -> [sqlc](https://sqlc.dev/?utm_source=chatgpt.com)
- Logging -> [slog](https://pkg.go.dev/log/slog)
- Profiling -> [pprof](https://pkg.go.dev/net/http/pprof)

## Monorepo

This is currently behaving as a monorepo as both apps are in their own folders under one git history.

To manage some common coordination from the root folder we can use `make`.

I've got some commands already in there -> run the api, build the api, run the api binary, run a vite dev server.

Just run `make <command name>` to try it out.

The two apps are currently not talking to each other but they do run fine independently.