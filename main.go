package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/karrick/godirwalk"
	"gopkg.in/yaml.v2"
)

var P = flag.String("P", "/home/ftpd/glftpd/site/mp3/1969-01-01", "Full scan path")
var C = flag.Bool("C", false, "Cleanup (delete folders)")
var D = flag.Bool("D", false, "Debug mode")

type Config struct {
	Clean   []string
	Keepdir []string
}

func genre(file string) string {
	tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()
	return tag.Genre()
}

func traverse(path string, cleangenres []string, keepdirs []string) {
	var rmdirs []string
	err := godirwalk.Walk(path, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				// check if we match any of our keep directories
				for _, d := range keepdirs {
					match, _ := regexp.MatchString(d, string(de.Name()))
					if match == true {
						if *D == true {
							fmt.Printf("KEEPDIR: %s/%s\n", path, string(de.Name()))
						}
						return nil
					}
				}
				// traverse through the directory and get the genre
				root := path + "/" + string(de.Name())
				var files []string
				err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
					if filepath.Ext(path) == ".mp3" {
						files = append(files, path)
					}
					return nil
				})
				if err != nil {
					panic(err)
				}
				if len(files) > 0 {
					genre := genre(files[0])
					for _, g := range cleangenres {
						if strings.ToLower(g) == strings.ToLower(genre) {
							if *D == true {
								fmt.Printf("DELETE : %s/%s %s\n", path, string(de.Name()), genre)
								if *C == true {
									rmdir := path + "/" + string(de.Name())
									rmdirs = append(rmdirs, rmdir)
									//_ = os.RemoveAll(rmdir)
									return nil
								}
								return nil
							}
							return nil
						}
					}
					if *D == true {
						fmt.Printf("KEEP   : %s/%s %s\n", path, string(de.Name()), genre)
					}
				}
			}
			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// do the actual deletion of the files
	for _, rmdir := range rmdirs {
		if *C == true {
			fmt.Printf("rm -rf %s\n", rmdir)
			err = os.RemoveAll(rmdir)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func main() {
	flag.Parse()
	filename, _ := filepath.Abs("./bob-id3clean.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	traverse(*P, config.Clean, config.Keepdir)
}
