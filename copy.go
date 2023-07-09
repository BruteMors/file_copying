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
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.OpenFile(fromPath, os.O_RDWR, 0755)
	if err != nil {
		return fmt.Errorf("problem when opening from file: %v", err)
	}
	defer fromFile.Close()
	stat, err := fromFile.Stat()
	if err != nil {
		return err
	}

	sizeFile := stat.Size()

	if offset > sizeFile {
		return ErrOffsetExceedsFileSize
	}

	if limit > sizeFile || limit == 0 {
		limit = sizeFile - offset
	}

	toFile, err := os.OpenFile(toPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		//if os.IsNotExist(err) {
		//	toFile, err = os.Create("out.txt")
		//	if err != nil {
		//		return err
		//	}
		//} else {
		//	return
		//}
		return fmt.Errorf("problem when opening to file: %v", err)
	}
	defer toFile.Close()

	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	var bufSize int64 = 1024
	if bufSize > limit {
		bufSize = limit
	}
	buf := make([]byte, bufSize)

	for limit > 0 {
		if limit < bufSize {
			bufSize = limit
		}
		read, err := fromFile.Read(buf[:bufSize])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		_, err = toFile.Write(buf[:read])
		if err != nil {
			return err
		}
		limit -= int64(read)
	}

	return nil
}
