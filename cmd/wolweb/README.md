# wasm

WebAssembly wrapper for the wol library.

## Build

```
GOOS=js GOARCH=wasm go build -o wol.wasm
```

## Use

The output file `wol.wasm` needs to be integrated in your web application.
You need:
* `wol.wasm`, obviously
* HTML code that will load `wol.wasm`
* Go's `wasm_exec.js`

The HTML code has to load the Go WASM stub `wasm_exec.js`. You can find it in
`$(go env GOROOT)/misc/wasm/wasm_exec.js`.

Then you can use `wol()` from JavaScript.

An example of HTML can be found at [index.html.example](index.html.example).
