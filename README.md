# `memgo` - Mimicking memory management in operating systems

this repo is my attempt at mimicking the memory management behavior achieved in operating systems

## Architecture

this will start off as a very simple, barebone implementation of memory management. as things progress, I want to tack on more features bit by bit.

- client-server architecture: 
    - *client* would be equivalent to processes that make requests to read/write memory
    - *server* would be equivalent to the operating system that manages memory allocation to mukltiple processes

1. client (process) is spun up
2. server (OS) allocates a block of memory to be given to this new process
3. processes make requests the read memory at different virtual addresses (VA)
4. upon receiving the request, OS (the MMU, technically) consults the page table to translate this request and find the physical address (PA) requested
5. if success, report back to the client with the memory
6. otherwise, the page table reports that the allocated memory is stored on disk. raise a page fault
7. if no page fault, repeat from step 3 until process terminates
8. otherwise, decide which block of data to evict from physical RAM. store this info on disk, and then fetch the requested memory from disk and overwrite this evicted data block. Report back to the client with the memory

### Actors 
From the above steps, these are the actors involved:
- client (process)
- virtual memory
- server (CPU/MMU + OS)
- page table
- physical memory
- disk

### Relationships
- client (processes) create code instructions, which invoke memory accesses
- CPU receives client's instructions, executes them by consulting the page table
- page table returns the physical address to access
- OS handles any page faults raised from the CPU
- OS handles access to physical disk in response to page faults

### Use scenarios

- make memory reads
    - at random intervals, with varying request amounts?
    - log the server (OS) response, such as:
        - the physical address of requested memory
        - the page table entries consulted
        - page faults (if any), and the evicted data block


