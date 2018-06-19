package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Parse flags after load config to allow override options in config
	cmdLineConfig := parseCmdLineConfig()

	fmt.Printf(`
       /\      MEOW Proxy %s
   )  ( ')     https://renzhn.github.io/MEOW/
  (  /  )      https://github.com/netheril96/MEOW
  \ (__)|      https://github.com/sastation/MEOW

	`, version)
	fmt.Println()

	parseConfig(cmdLineConfig.RcFile, cmdLineConfig)

	initSelfListenAddr()
	initLog()
	initAuth()
	initStat()

	initParentPool()

	if config.JudgeByIP {
		initCNIPData()
	}

	if config.Core > 0 {
		runtime.GOMAXPROCS(config.Core)
	}

	var wg sync.WaitGroup
	wg.Add(len(listenProxy))
	for _, proxy := range listenProxy {
		go proxy.Serve(&wg)
	}
	wg.Wait()
}
