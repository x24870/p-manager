package diskwalker

import (
	"fmt"
)

// DiskWalker is a interface for walking a disk
type DiskWalker interface {
	Walk() error
}

// DiskWalkerImpl is a struct for walking a disk
type DiskWalkerImpl struct{}

// NewDiskWalker returns a new DiskWalker
func NewDiskWalkerImpl() DiskWalkerImpl {
	return DiskWalkerImpl{}
}

func (d *DiskWalkerImpl) Walk() error {
	fmt.Println("Walking the disk")
	return nil
}
