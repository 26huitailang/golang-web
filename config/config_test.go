package config

import (
	"os"
	"runtime/pprof"
	"testing"
)

func BenchmarkInitTheme(b *testing.B) {
	var conf *Configuration
	conf.initConfiguration()
	f, err := os.Create("cpu.prof")
	if err != nil {
		b.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	for i := 0; i < b.N; i++ {
		conf.InitTheme()
	}
}
