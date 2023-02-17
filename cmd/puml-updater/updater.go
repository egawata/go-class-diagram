package main

import (
	"flag"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	goplantuml "github.com/jfeliu007/goplantuml/parser"
)

var (
	srcDir  string
	outFile string
)

func init() {
	flag.StringVar(&srcDir, "watch", "", "directory to be watched")
	flag.StringVar(&outFile, "o", "./out.puml", "output file (.puml)")
}

func main() {
	flag.Parse()
	log.Printf("srcdir: %s, outfile: %s", srcDir, outFile)

	// 起動時に1回だけ puml ファイル生成をしておく
	update()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Printf("modified file: %s", event.Name)
					update()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()

	err = watcher.Add(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

func update() {
	res, err := render(srcDir)
	if err != nil {
		log.Println(err)
		return
	}
	err = writeToFile(outFile, res)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func render(dir string) (string, error) {
	dirs := []string{dir}
	result, err := goplantuml.NewClassDiagram(dirs, nil, true)
	if err != nil {
		return "", err
	}

	return result.Render(), nil
}

func writeToFile(out string, puml string) error {
	f, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write([]byte(puml))
	return nil
}
