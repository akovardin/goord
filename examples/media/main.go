package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"gohome.4gophers.ru/kovardin/goord/ord"
)

var create = false

func main() {
	client, _ := ord.NewClient(
		ord.WithBase("https://api-sandbox.ord.vk.com"),
		ord.WithToken(os.Getenv("TOKEN")),
	)

	fmt.Println("Getting list of media...")

	mediaList, err := client.GetMediaList(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting media list: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d media items (total: %d)\n", len(mediaList.ExternalIDs), mediaList.TotalItemsCount)
		for i, id := range mediaList.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	externalID := "test-media-001"

	fmt.Println("Uploading media...")

	if create {
		img := generateTestImage()
		var buf bytes.Buffer
		err = png.Encode(&buf, img)
		if err != nil {
			log.Printf("Error encoding image: %v\n", err)
			return
		}

		sha256, err := client.UploadMedia(context.Background(), externalID, "test.png", &buf)
		if err != nil {
			log.Printf("Error uploading media: %v\n", err)
		} else {
			fmt.Printf("Media %s uploaded successfully. SHA256: %s\n", externalID, *sha256)
		}

		fmt.Println("Getting media info...")

		mediaInfo, err := client.GetMediaInfo(context.Background(), externalID)
		if err != nil {
			log.Printf("Error getting media info: %v\n", err)
		} else {
			fmt.Printf("Media info: %+v\n", mediaInfo)
			fmt.Printf("Filename: %s\n", mediaInfo.Filename)
			fmt.Printf("Size: %d\n", mediaInfo.Size)
			fmt.Printf("Content Type: %s\n", mediaInfo.ContentType)
		}
	}

	fmt.Println("Getting media binary data...")

	binaryData, err := client.GetMediaBinary(context.Background(), externalID)
	if err != nil {
		log.Printf("Error getting media binary data: %v\n", err)
	} else {
		fmt.Printf("Retrieved binary data of size: %d bytes\n", len(binaryData))
		fmt.Printf("First 100 bytes: %v\n", binaryData[:100])
	}

	fmt.Println("Getting batch media info...")

	mediaInfos, err := client.GetMediaInfoBatch(context.Background(), []string{externalID})
	if err != nil {
		log.Printf("Error getting batch media info: %v\n", err)
	} else {
		fmt.Printf("Retrieved info for %d media items\n", len(mediaInfos))
		for i, info := range mediaInfos {
			fmt.Printf("  %d. %s (%s)\n", i+1, info.ExternalID, info.Filename)
		}
	}
}

func generateTestImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			r := uint8(x * 255 / 100)
			g := uint8(y * 255 / 100)
			b := uint8((x + y) * 255 / 200)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	return img
}
