package cp

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func File(dst string, src []byte) error {
	f, err := create(dst)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	w := bufio.NewWriter(f)
	fmt.Println("Writing to ", dst)
	nn, err := w.Write(src)
	if err == nil {
		err = w.Flush()
	}
	if err != nil {
		log.Fatalln(err)
		fmt.Fprintln(os.Stderr, err)
		return err
	} else {
		fmt.Printf("%d bytes written to %s\n", nn, dst)
	}
	return nil
}
