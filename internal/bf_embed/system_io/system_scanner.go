package system_io

import (
	"bufio"
	"fmt"
	"os"

	"butterfly_embed/runtime"
)

func init() {
	const moduleName = "Standard Input"

	runtime.BF__EventSubscribe("Sys Scan", moduleName, func(message runtime.BF__MessageContent) {
		var args []any
		var ok bool
		if args, ok = message["args"].([]interface{}); !ok {
			panic("Printf args are not an array!")
		}
		fmt.Scan(args...)
		if event, ok := message["shares"]; ok {
			runtime.BF__Dispatch(event.(string), runtime.BF__MessageCreate(moduleName, map[string]interface{}{
				"input": args,
			}))
		}
	})

	runtime.BF__EventSubscribe("Sys Scanf", moduleName, func(message runtime.BF__MessageContent) {
		var format string
		var args []any
		var ok bool
		if format, ok = message["format"].(string); !ok {
			panic("Printf message is not a string!")
		}
		if args, ok = message["args"].([]interface{}); !ok {
			panic("Printf args are not an array!")
		}
		fmt.Scanf(format, args...)
		if event, ok := message["shares"]; ok {
			runtime.BF__Dispatch(event.(string), runtime.BF__MessageCreate(moduleName, map[string]interface{}{
				"input": args,
			}))
		}
	})

	runtime.BF__EventSubscribe("Sys Scanln", moduleName, func(message runtime.BF__MessageContent) {
		var line, err = bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		if event, ok := message["shares"]; ok {
			runtime.BF__Dispatch(event.(string), runtime.BF__MessageCreate(moduleName, map[string]interface{}{
				"input": line,
			}))
		}
	})
}
