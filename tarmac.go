//
// Package tarmac provides simple abstractions around the archive/tar package for
// common high-level use cases.
//
package tarmac

import (
    "os"
    "io"
	  "archive/tar"
)

// GetFileList returns a slice of tar.Header structs read
// from the specified tar file.
func GetFileList(tarfile string) ([]tar.Header, error) {
	// open tar file
	tf, err := os.Open(tarfile)
	if err != nil {
	    return nil, err
	}
	defer tf.Close()

	// tar reader
	tarc := tar.NewReader(tf)

	// return slice with a default capacity
	headers := make([]tar.Header, 0, 100)

	var hdr *tar.Header
	var rerr error = nil

	hdr, rerr = tarc.Next()
	for (rerr != io.EOF) {
	    headers = append(headers, *hdr)
	    hdr, rerr = tarc.Next()
    }

	return headers, nil
}