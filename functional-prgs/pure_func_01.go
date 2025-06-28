package main

import (
	"github.com/devak23/go/functional-prgs/utils"
	"os"
)

func main() {
	// The Add is defined in functional-prgs/utils/func_utils.go but wasn't really exported to github repository at the time
	// of writing this code. However, it was successfully referenced in this program due to 4 main corrections that were
	// made:
	// 1. The utils package was imported correctly i.e. line# 3.
	// 2. The Add function was defined with capital 'A' in the utils package.
	// 3. The utils package had a mod.go that was removed. And lastly -
	// 4. The module was defined correctly in functional-prgs/go.mod file.
	os.Exit(utils.Add(3, 4))
}
