package main

import (
	"fmt"
	"log"
	"os"
)

type Stream struct {
	Signature      []byte // First 4 bytes
	MetaDataBlocks []byte
	AudioFrames    []byte
}

type MetadataBlockHeader struct {
	IsLast   bool
	Type     MetadataBlockDataType
	DataSize uint
}

// METADATA_BLOCK_DATA types

type MetadataBlockDataType uint8

const (
	METADATA_BLOCK_STREAMINFO MetadataBlockDataType = iota
	METADATA_BLOCK_PADDING
	METADATA_BLOCK_APPLICATION
	METADATA_BLOCK_SEEKTABLE
	METADATA_BLOCK_VORBIS_COMMENT
	METADATA_BLOCK_CUESHEET
	METADATA_BLOCK_PICTURE
)

// this array of bytes is what is in the first 4 bytes of a flac file
var FLAC_LABEL = []byte{'f', 'L', 'a', 'C'}

func main() {
	file, err := os.ReadFile("music.flac")
	if err != nil {
		log.Fatalln(err)
	}

	// if the beginning of the file does not have the "fLaC" array of bytes,
	// it is not a flac
	if string(file[:4]) != string(FLAC_LABEL) {
		panic("File is not a .flac")
	}
	fmt.Println(ReadBlockHeader(file[4:8]))
}

func ReadBlockHeader(blockHeader []byte) MetadataBlockHeader {
	isLast := uint(blockHeader[0] & 1)
	return MetadataBlockHeader {
		Type:     MetadataBlockDataType(blockHeader[0] & 0x7F),
		IsLast:   isLast != 0,
		DataSize: GetSizeOfBlockData(blockHeader[0:4]),
	}
}

func GetSizeOfBlockData(blockHeader []byte) uint {
	// Convert bytes to an integer, for the byte length of the metadata block
	var byteSize uint
	for i := 0; i < 3; i++ {
		byteSize |= uint(blockHeader[0:4][i]) << (8 * uint(i))
	}
	return byteSize
}
