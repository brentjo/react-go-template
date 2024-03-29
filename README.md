## React Go Base

A template repository to create simple React web apps with a Go powered backend. The output is a single binary with your static JS/HTML/CSS assets embedded within it, for easy distribution and running on low spec'd hardware like Raspberry Pis or other devboards.

<p align="center">
  <img src="https://github.com/brentjo/react-go-template/assets/6415223/e8551741-50ee-4b51-aa02-1cc31d4c8bb7" width="400px">
  <img src="https://github.com/brentjo/react-go-template/assets/6415223/76fc8c46-22ec-493d-a7b7-4260df7717f8" width="400px">
</p>

### Setup and usage

Clone this repository:
```
git clone https://github.com/brentjo/react-go-template && cd react-go-template
```

Build the project:
```
make
```

Run the binary:
```
./go-react-example
```

### Requirements

- **Go**: https://go.dev/dl/
- **esbuild**: installed and accessible in your `$PATH` — `npm install -g esbuild`
- **make**: `brew install make` if you don't already have it, and use homebrew

## How does it work?
1. You write React components within [`src/components`](src/components)
    - There are two examples in this repository:
      - [`counter.jsx`](src/components/counter.jsx): a click-counter that's purely client side driven.
      - [`time.jsx`](src/components/time.jsx): a component that fetches the current time from the `/api/time` route on the Go backend.
    - Update the root [`app.jsx`](src/app.jsx) to use the new components you write. There is also minimal client side routing support, so you can [render different components per path](https://github.com/brentjo/react-go-template/blob/main/src/app.jsx#L25-L34).
2. Write any needed backend APIs within Go
    - For example, this handler for fetching the current time:
      - [`customizations.go`](https://github.com/brentjo/react-go-template/blob/main/customizations.go#L10-L11)
    - If you want any other paths besides the root `/` to render your React app, be sure to [add them to `allowedPathsForSPA`](https://github.com/brentjo/react-go-template/blob/main/customizations.go#L13-L14)
3. At [build time](https://github.com/brentjo/react-go-template/blob/1b4bef9c465cd39c071d65e95f64a74e723cc938/Makefile#L3-L7), the JSX is transpiled with `esbuild` and copied into a directory for publishing, along with the static React library dependencies within the `/static` folder
4. These contents are placed into an [embedded](https://pkg.go.dev/embed) file system at compile time, so the final output is self-contained within the binary, and there's no need to copy around HTML/JS/CSS separately.
5. Run the built binary, and you have your React app being served, along with any API routes you wrote.


## Why?
In the past I would use Next.js or Ruby with Sinatra when I needed to create a simple local-network web app, but I always found working with them on devboards/mini-PCs a bit painful: the devboard might not support the runtime, deploying and managing the dependencies with `npm` / `gem` was tough and prone to breakage depending on how it was set up, the apps would often run out of memory, etc.

Go is fantastic on all these fronts: it's compiled to a single binary for easy deployment, the language has great hardware support, it's very fast and you're in control of your memory footprint while still working in a memory-safe language, etc.

You would never want to build out a React app with any amount of complexity with this. See https://react.dev/learn/start-a-new-react-project for why you generally want to use a 'batteries included framework' when working with React, but for the use-cases like 'simple single page IoT dashboard' that I just want to style with plain CSS and use minimal 3rd party libraries, I found this pattern pleasant to work with.
