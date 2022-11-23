// Compile Debug Mode modifed

package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/pkg/profile"
)

// const pprofMode = "pcpu"

const pprofMode = "pmem"

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func main() {

	// differert pprof mode
	switch pprofMode {
	case "cpu":
		f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	case "mem":
		f, _ := os.OpenFile("mem.pprof", os.O_CREATE|os.O_RDWR, 0644)
		defer f.Close()
		pprof.WriteHeapProfile(f)
	case "pcpu":
		defer profile.Start().Stop()
	case "pmem":
		defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	}

	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}
