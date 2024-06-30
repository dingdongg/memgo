package main

import "fmt"

const VIRTUAL_MEM_ADDR_SPACE = 32
const PHYSICAL_MEM_ADDR_SPACE = 24
const DISK_SIZE = 1 << 50
const PAGE_OFFSET_SIZE = 14 // each page = 16KB in size

type PageTable map[uint]uint

type Memory struct {
	physical []byte			// equivalent to RAM. RAM serves as a fast cache for the 100,000X slower disk storage
	pageTable *PageTable	// page table
	secondary []byte 		// not sure what the appropriate type would be, just do this for now
}

func NewMemory() *Memory {
	pageTable := make(PageTable)
	return &Memory{
		physical: make([]byte, 1 << PHYSICAL_MEM_ADDR_SPACE),
		pageTable: &pageTable,
		secondary: make([]byte, DISK_SIZE),
	}
}

func (m *Memory) Read(addr uint, n int) []byte {
	offsetMask := uint((1 << PAGE_OFFSET_SIZE) - 1)
	offset := addr & uint(offsetMask)
	pageNum := m.pageNum(addr)

	// page num is used to index into the page table.
	// afterrwards, combine with offset to get physical address
	ppn := m.getPPN(pageNum)
	physAddr := ppn | offset
	return m.physical[physAddr : physAddr+4]
}

func (m *Memory) pageNum(addr uint) uint {
	offsetMask := uint((1 << PAGE_OFFSET_SIZE) - 1)
	bitmask := ^offsetMask
	return (addr & bitmask) >> PAGE_OFFSET_SIZE
}

func (m *Memory) getPPN(pageNum uint) uint {
	return 0
}

func (m *Memory) Write(addr uint, data []byte) error {
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