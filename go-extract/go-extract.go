package extract

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		fmt.Println(1, err)
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Println(2, err)
	}
	defer archive.Close()

	writer, err := os.Create(target)
	if err != nil {
		fmt.Println(3, err)
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err

}

func Untar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return err
}
