package main

import (
	"log"
	"syscall"
	"time"
)

// ByteSize type
type ByteSize float64

// nolint
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB

	SleepWaiting = 1
)

// DiskStatus type
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// DiskUsage of path/disk
func DiskUsage(path string) (*DiskStatus, error) {
	fs := &syscall.Statfs_t{}

	if err := syscall.Statfs(path, fs); err != nil {
		return nil, err
	}

	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)

	return &DiskStatus{
		All:  all,
		Free: free,
		Used: all - free,
	}, nil
}

func main() {
	var (
		err error
		ds  *DiskStatus
	)

	for {
		if ds, err = DiskUsage("/"); err != nil {
			log.Fatalf("main: %s", err)
		}

		log.Printf("All: %.2f GB\n", float64(ds.All)/float64(GB))
		log.Printf("Used: %.2f GB\n", float64(ds.Used)/float64(GB))
		log.Printf("Free: %.2f GB\n", float64(ds.Free)/float64(GB))

		time.Sleep(SleepWaiting * time.Hour)
	}
}
