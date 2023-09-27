package main

import (
	"fmt"
	"log"
	"os"
)

type Stream struct {
    Signature []byte    // First 4 bytes
    MetaDataBlocks []byte
    AudioFrames []byte
}

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
    fmt.Println(file[4:8])
    // MetaData Block Header
    fmt.Println(uint(file[4] & 1))
    fmt.Println(uint(file[4] & 0x7F))
    fmt.Println(file[5:8])
    // Convert bytes to an integer, for the byte length of the metadata block
    var byteSize uint32
    for i := 0; i < 3; i++ {
        byteSize |= uint32(file[5:8][i]) << (8 * uint(i))
    }

}
