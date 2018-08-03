package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/itchio/lzma"
)

// https://www.reddit.com/r/algotrading/comments/65kb14/dukascopy_forex_data/

// http://datafeed.dukascopy.com/datafeed/NZDCAD/2018/06/26/17h_ticks.bi5

// 32-bit integer: milliseconds since epoch
// 32-bit integer: Ask price
// 32-bit integer: Bid price
// 32-bit float: Ask volume
// 32-bit float: Bid volume

type Block struct {
	time   uint32
	ask    uint32
	bid    uint32
	askVol float32
	bidVol float32
}

func readBlock(r io.Reader, b *Block) error {

	data := make([]byte, 4*5)

	n, err := r.Read(data)
	if err != nil {
		return err
	}

	if n != len(data) {
		return fmt.Errorf("invalid count bytes %d", n)
	}

	buf := bytes.NewReader(data)

	binary.Read(buf, binary.BigEndian, &b.time)
	binary.Read(buf, binary.BigEndian, &b.ask)
	binary.Read(buf, binary.BigEndian, &b.bid)
	binary.Read(buf, binary.BigEndian, &b.askVol)
	binary.Read(buf, binary.BigEndian, &b.bidVol)

	return nil
}

func main() {

	file, err := os.Open("17h_ticks.bi5")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := lzma.NewReader(file)

	err = nil
	block := Block{}

	for err == nil {
		err = readBlock(reader, &block)
		fmt.Println(block)
		break
	}

	fmt.Println(err)
}
