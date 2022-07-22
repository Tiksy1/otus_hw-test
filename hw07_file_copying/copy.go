package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotCreated        = errors.New("file not created")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 0
	}
	fromFileStat, err := os.Stat(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	if offset > fromFileStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	fromFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fromFile.Close()

	toFile, err := os.Create(toPath)
	if err != nil {
		return ErrFileNotCreated
	}
	defer toFile.Close()

	if limit == 0 || limit+offset > fromFileStat.Size() {
		limit = fromFileStat.Size() - offset
	}

	var bar Bar
	fromFile.Seek(offset, 0)
	buf := make([]byte, 32)
	var read, bufSum int
	bar.NewOption(offset, limit)
	for bufSum < int(limit) {
		read, err = fromFile.Read(buf)
		if !errors.Is(err, io.EOF) && err != nil {
			return err
		}
		if bufSum+read > int(limit) {
			read = int(limit) - bufSum
		}
		toFile.Write(buf[:read])
		bufSum += read
		bar.Play(int64(bufSum))
	}
	bar.Finish()

	return nil
}

// Bar ...
type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 20 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Print("\nКопирование завершено\n\n")
}
