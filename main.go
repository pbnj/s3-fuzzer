package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Sirupsen/logrus"
	open "github.com/petermbenjamin/go-open"
)

var (
	fileFlag  = kingpin.Flag("file", "Path to file.").Short('f').String()
	nameFlag  = kingpin.Flag("name", "Bucket name.").Short('n').String()
	openFlag  = kingpin.Flag("open", "Open URL in default browser.").Short('o').Bool()
	debugFlag = kingpin.Flag("debug", "Debug mode.").Short('d').Action(func(c *kingpin.ParseContext) error {
		logrus.SetLevel(logrus.DebugLevel)
		return nil
	}).Bool()
)

func main() {
	kingpin.Parse()
	var wg sync.WaitGroup
	if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			log.Fatalf("could not open file: %+v", err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			wg.Add(1)
			go fuzz(scanner.Text(), &wg)
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("could not scan file: %+v", scanner.Err())
		}
	}

	if *nameFlag != "" {
		wg.Add(1)
		go fuzz(*nameFlag, &wg)
	}
	wg.Wait()
}

func fuzz(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	logrus.Debugf("checking %+v", name)
	resp, err := http.Get("http://" + name + ".s3.amazonaws.com")
	if err != nil {
		logrus.Errorf("could not GET bucket: %+v", err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("could not read response body: %+v", err)
	} else {
		if resp.StatusCode == http.StatusOK {
			logrus.Infof("Bucket found: %s.s3.amazonaws.com", name)
			if *openFlag {
				open.Open("http://" + name + ".s3.amazonaws.com")
			}
		}
	}
}
