package main

import (
	tools "eximgolog/tools"
)

func main() {
	teste := tools.ReadLog("mainlog")
	for _, t := range teste {
		tools.InsertLogLine(t)
	}

}
