package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type block struct {
	free   bool
	fileId int
}

type diskMap struct {
	disk      []block
	lastUsed  int
	firstFree int
}

func readData(inputFile string) diskMap {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileId := 0
	free := false
	var disk diskMap

	for _, c := range data {
		blockLength := int(c) - int('0')
		for j := 0; j < blockLength; j++ {
			if !free {
				disk.lastUsed = len(disk.disk)
				disk.disk = append(disk.disk, block{false, fileId})
			} else {
				disk.disk = append(disk.disk, block{true, 0})
			}

		}
		if !free {
			fileId++
		}
		free = !free
	}
	for i, b := range disk.disk {
		if b.free {
			disk.firstFree = i
			break
		}
	}

	return disk
}

func updateFirstFree(data *diskMap) {
	for i := (*data).firstFree; i < len((*data).disk); i++ {
		if (*data).disk[i].free {
			(*data).firstFree = i
			return
		}
	}
}

func updateLastUsed(data *diskMap) {
	for i := (*data).lastUsed - 1; i >= 0; i-- {
		if !(*data).disk[i].free {
			(*data).lastUsed = i
			return
		}
	}
	panic("No used blocks found")
}

func nextLastFile(data *diskMap, lastFile int) int {
	var fileId int
	foundFile := false

	for i := lastFile - 1; i >= 0; i-- {
		if foundFile && ((*data).disk[i].free || (*data).disk[i].fileId != fileId) {
			return i + 1
		}
		if !foundFile && !(*data).disk[i].free {
			foundFile = true
			fileId = (*data).disk[i].fileId
		}
	}
	return -1
}

func getNextFreeBlock(data *diskMap, freeBlock int) int {
	foundFile := false

	for i := freeBlock; i < len((*data).disk); i++ {
		if i == -1 || !(*data).disk[i].free {
			foundFile = true
		} else if foundFile && (*data).disk[i].free {
			return i
		}
	}
	panic("No free block")
}

func getSize(data *diskMap, start int) int {
	var freeBlocks bool
	var fileId int

	freeBlocks = (*data).disk[start].free
	fileId = (*data).disk[start].fileId

	for i := start; i < len((*data).disk); i++ {
		if freeBlocks && !(*data).disk[i].free {
			return i - start
		} else if !freeBlocks && fileId != (*data).disk[i].fileId {
			return i - start
		}
	}
	return len((*data).disk) - start
}

func runPart1(data diskMap) int {
	for {
		if data.firstFree > data.lastUsed {
			break
		}
		data.disk[data.firstFree].fileId = data.disk[data.lastUsed].fileId
		data.disk[data.firstFree].free = false
		data.disk[data.lastUsed].free = true
		updateFirstFree(&data)
		updateLastUsed(&data)
	}

	checksum := 0
	for i, b := range data.disk {
		if b.free {
			break
		}
		checksum += i * b.fileId
	}

	return checksum
}

func runPart2(data diskMap) int {
	var lastFile = len(data.disk)
	for {
		lastFile = nextLastFile(&data, lastFile)
		if lastFile == -1 {
			break
		}
		fileSize := getSize(&data, lastFile)
		nextFreeBlock := -1
		for {
			nextFreeBlock = getNextFreeBlock(&data, nextFreeBlock)
			if nextFreeBlock > lastFile {
				break
			}

			freeSize := getSize(&data, nextFreeBlock)
			if freeSize >= fileSize {
				for i := 0; i < fileSize; i++ {
					data.disk[nextFreeBlock+i].fileId = data.disk[lastFile].fileId
					data.disk[nextFreeBlock+i].free = false
					data.disk[lastFile+i].free = true
				}
				break
			}
		}
	}

	checksum := 0
	for i, b := range data.disk {
		if !b.free {
			checksum += i * b.fileId
		}
	}

	return checksum
}

func main() {
	var inputFile = flag.String("input", "inputs/day09.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	data = readData(*inputFile)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}
