package main

import "fmt"

const VIRTUAL_MEM_SIZE = 1 << 40
const PHYSICAL_MEM_SIZE = 1 << 32
const DISK_SIZE = 1 << 62

type VirtAddr uint64
type PhysAddr uint64

type PageTable map[VirtAddr]PhysAddr

type Memory struct {
	physical []byte			// equivalent to RAM
	pageTable *PageTable	// page table
	secondary []byte 		// not sure what the appropriate type would be, just do this for now
}

func NewMemory() *Memory {
	pageTable := make(PageTable)
	return &Memory{
		physical: make([]byte, PHYSICAL_MEM_SIZE),
		pageTable: &pageTable,
		secondary: make([]byte, DISK_SIZE),
	}
}

func main() {
	fmt.Println("HELLO WORLD")
}


/*
- client (process)
- virtual memory
- server (CPU/MMU + OS)
- page table
- physical memory
- disk -> will probably represent as a text file; 
*/