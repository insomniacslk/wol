package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"syscall/js"

	"github.com/insomniacslk/wol"
)

func jsFromBytes() js.Func {
	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "need exactly one argument in call to wol"
		}
		mac := args[0].String()
		if err := wol.Send(mac); err != nil {
			return fmt.Errorf("failed to send WOL packet to %s: %w", mac, err).Error()
		}
		return "packet sent"
	})
	return fn
}

func main() {
	fmt.Println("Loaded wolweb WASM module")
	js.Global().Set("sendWOL", jsFromBytes())

	// keep the code running or it will fail with "Go program has already exited"
	select {}
}
