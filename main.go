package main

import (
	"fmt"
	"sync"
	"time"
)

const VIRTUAL_MEM_ADDR_SPACE = 32
const PHYSICAL_MEM_ADDR_SPACE = 24
const DISK_SIZE = 1 << 40
const PAGE_OFFSET_SIZE = 14 // each page = 16KB in size

type PageTable map[uint]uint

// https://ocw.mit.edu/courses/6-004-computation-structures-spring-2017/pages/c16/c16s2/c16s2v2/

type PageFault struct {
	callback any // will be some sort of callback function to execute after CPU exits kernel mode
}

type Memory struct {
	physical []byte			// equivalent to RAM. RAM serves as a fast cache for the 100,000X slower disk storage
	pageTable *PageTable	// page table
	secondary []byte 		// not sure what the appropriate type would be, just do this for now
	pageFaultQueue chan *PageFault
}

func NewMemory() *Memory {
	pageTable := make(PageTable)
	return &Memory{
		physical: make([]byte, 1 << PHYSICAL_MEM_ADDR_SPACE),
		pageTable: &pageTable,
		secondary: make([]byte, DISK_SIZE),
		pageFaultQueue: make(chan *PageFault),
	}
}

func (m *Memory) Read(addr uint, n int) []byte {
	offsetMask := uint((1 << PAGE_OFFSET_SIZE) - 1)
	offset := addr & uint(offsetMask)
	pageNum := m.getVPN(addr)

	// page num is used to index into the page table.
	// afterrwards, combine with offset to get physical address
	ppn := m.getPPN(pageNum)
	physAddr := ppn | offset
	return m.physical[physAddr : physAddr+4] 
}

func (m *Memory) getVPN(addr uint) uint {
	offsetMask := uint((1 << PAGE_OFFSET_SIZE) - 1)
	bitmask := ^offsetMask
	return (addr & bitmask) >> PAGE_OFFSET_SIZE
}

func (m *Memory) getPPN(pageNum uint) uint {
	ppn, exists := (*m.pageTable)[pageNum]

	fmt.Println(ppn)
	if !exists {
		// TODO: implement page fault + fetching from disk
		fmt.Println("pagefault")
	}

	/*
	(more research required on this)
	so  far, it seems like page faults are implemented via CPU interrupts in the hardware.
		- this means there is no "event queue" or "polling" involved
	I can't think of a way to replicate this hardware implementation,
	so for now I think I will spawn a goroutine to concurrently listen for & handle
	page fault events (ie. an event queue lol)
	*/
	return 0
}

func (m *Memory) Write(addr uint, data []byte) error {
	return nil
}

// workaround for the CPU interrupt mechanism
// TODO: implement using m.pageFaultQueue
func (m *Memory) listenForPageFaults() {
	for {
		time.Sleep(time.Second)
		fmt.Println("listenin...")
	}
}

// main function can simulate the memory access requests from the CPU
func main() {
	fmt.Println("HELLO WORLD")

	mem := NewMemory() // upon process startup, memory is allocated to the process
	// fmt.Printf("%+v\n", mem)

	var wg sync.WaitGroup

	wg.Add(1)
	go mem.listenForPageFaults()
	// bunch of read/write requests to memory
	wg.Wait()
}


/*
- client (process)
- virtual memory
- server (CPU/MMU + OS)
- page table
- physical memory
- disk -> will probably represent as a text file; 
*/