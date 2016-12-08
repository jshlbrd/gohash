//Simple directory scan and hashing package
//Contains options to specify input directory, output file, and maximum file size for files to be hashed
//Josh Liburdi 2016

package main

import (
	"crypto/md5"
	"fmt"
	"path/filepath"
	"encoding/hex"
	"io"
	"os"
	"flag"
)

//http://www.mrwaggel.be/post/generate-md5-hash-of-a-file/
func hash_file_md5(filePath string) (string, error) {
	var hashString string
	file, err := os.Open(filePath)
	if err != nil {
		return hashString, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashString, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	hashString = hex.EncodeToString(hashInBytes)
	return hashString, nil
}


func main() {
	inputDirPtr := flag.String("scanDir","","The directory that will be scanned")
	outputFilePtr := flag.String("outputFile","/tmp/gohash.dat","The file where MD5 hashes will be output")
	fsizePtr := flag.Int("maxSize",10000000,"The maximum size of files to be hashed")
	flag.Parse()

    	files := []string{}
    	filepath.Walk(*inputDirPtr, func(path string, f os.FileInfo, err error) error {
		if f.Size() <= int64(*fsizePtr) {
	        	files = append(files, path)
		}
		return nil
    	})

	out, err := os.Create(*outputFilePtr)
	if err != nil {
		fmt.Println(out, err)
	}

    	for _, file := range files {	
		hash, err := hash_file_md5(file)
		if err == nil {
			var outString string = file + ": " + hash + "\n"
			write, err := io.WriteString(out,outString)
			if err != nil {
				fmt.Println(write, err)
			}
		}
    	}
	
	out.Close()
}
