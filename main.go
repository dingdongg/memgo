package main

import "fmt"

const VIRTUAL_MEM_SIZE = 1 << 40
const PHYSICAL_MEM_SIZE = 1 << 32
const DISK_SIZE = 1 << 50

type VirtAddr uint64
type PhysAddr uint64

type PageTable map[VirtAddr]PhysAddr

type Memory struct {
	physical []byte			// equivalent to RAM. RAM serves as a fast cache for the 100,000X slower disk storage
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

func (m *Memory) Read(addr VirtAddr, n int) []byte {
	return []byte{}
}

func (m *Memory) Write(addr VirtAddr, data []byte) error {
	return nil
}

// main function can simulate the memory access requests from the CPU
func main() {
	fmt.Println("HELLO WORLD")

	mem := NewMemory() // upon process startup, memory is allocated to the process
	fmt.Printf("%+v\n", mem)

	// bunch of read/write requests to memory
}


/*
- client (process)
- virtual memory
- server (CPU/MMU + OS)
- page table
- physical memory
- disk -> will probably represent as a text file; 
*/