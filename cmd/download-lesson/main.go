package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/hashicorp/go-getter"
)

const urlLocation = "http://localhost:4566/gameset-704289e53f01556-moneyduck-vertical-1080-hli/gameset.tar.gz?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=dummy%2F20250320%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250320T174521Z&X-Amz-Expires=1200&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=201d4b0222bb92694306b6a4df40a355f0970b7cbc2c22cbb9ea465cce12c3d0"

func main() {
	// usingGetter()

	usingGrabWithProgress()
}

func usingGetter() {
	c := getter.Client{
		Ctx:              context.Background(),
		Src:              urlLocation,
		Dst:              "./download",
		Pwd:              "..",
		Mode:             0,
		Umask:            0,
		Detectors:        nil,
		Decompressors:    nil,
		Getters:          nil,
		Dir:              true,
		ProgressListener: nil,
		Insecure:         false,
		DisableSymlinks:  false,
		Options:          nil,
	}

	fmt.Println("Downloading...")

	err := c.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}

func usingGrab() {
	fmt.Println("Downloading...")

	resp, err := grab.Get(".", urlLocation)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")
	fmt.Println("Download saved to", resp.Filename)
}

func usingGrabWithProgress() {
	client := grab.NewClient()
	req, _ := grab.NewRequest("b4ea0fe7-bfa3-4427-bece-ff83d9344b32", urlLocation)

	sum, err := hex.DecodeString("F9DF12BD3787437EB0AF40F97830B60CC0ADD2C48ACA41C7C2BB492FB6C5D016")
	if err != nil {
		panic(err)
	}
	req.SetChecksum(sha256.New(), sum, true)

	fmt.Println("Downloading... URL: ", req.URL())

	resp := client.Do(req)
	fmt.Printf("\tResponse status: %s\n", resp.HTTPResponse.Status)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf(
				"\ttransf. %v / %v bytes total (%.2f%%)\n", resp.BytesComplete(), resp.Size(), 100*resp.Progress(),
			)
		case <-resp.Done:
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf(
		"\ttransf. %v / %v bytes total (%.2f%%)\n", resp.BytesComplete(), resp.Size(), 100*resp.Progress(),
	)

	fmt.Printf("Download completed! Save to: ./%v\n", resp.Filename)
}
