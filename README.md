# Gollang-Memory-And-CPU

# Configure your builds
```
When reading data, the internal data registers of a modern computer's CPU can hold and process 64-bit.
This is called word size. It is usually 32-bit or 64-bit.

When we don't align our data to fit word sizes, the next field is a word size
padding is added to properly align the fields in memory so that it can start at an offset with multiples.
Modern CPU hardware reads and writes to memory most efficiently when data is naturally aligned.

Go compiler to make sure memory stored side by side is accessible using a common floor
uses the required alignment. Its value is equal to the size of the memory required for the largest space in 
the structure.

When a build is created in Go, a contiguous memory block is allocated for it. go memory allocator,
it doesn't optimize for data structure alignment, so rearrange the fields of your structure to fill it
You can reduce your memory usage.
```
# Example;
![carbon (5)](https://user-images.githubusercontent.com/92402372/186905545-1527dc9c-8ba0-4582-8058-3315524706f5.png)


With our struct editing, we regained 64 bits, 8 bytes, and reduced our memory usage.
