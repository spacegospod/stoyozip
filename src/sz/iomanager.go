package sz

import (
	"bufio"
	"log"
	"os"
)

type iomanager struct {
	writer   *bufio.Writer
	out_file *os.File
}

func NewIomanager(out_path string) *iomanager {
	iom := new(iomanager)

	of, err := os.Create(out_path)
	if err != nil {
		log.Fatal(err.Error())
	}

	iom.out_file = of
	iom.writer = bufio.NewWriter(of)

	return iom
}

func (iom *iomanager) writeChunk(chunk []byte) {
	if _, err := iom.writer.Write(chunk); err != nil {
		log.Println(err.Error())
	}

	if iom.writer.Available() < 128 {
		iom.writer.Flush()
	}
}

func (iom *iomanager) flush() {
	iom.writer.Flush()
	iom.out_file.Close()
}
