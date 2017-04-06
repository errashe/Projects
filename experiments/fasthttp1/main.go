package main

import "fmt"
import "github.com/valyala/fasthttp"
import "runtime"

func HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q.",
		ctx.Path())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fasthttp.ListenAndServe(":8000", HandleFastHTTP)
}
