package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"gohome.4gophers.ru/kovardin/goord"
)

func main() {
	client, err := goord.NewClientWithResponses(
		"https://api-sandbox.ord.vk.com",
		goord.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+os.Getenv("KEY"))

			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get("https://kodikapusta.ru/api/files/pbc_1125843985/6qnij2s67zzm6c8/gdsh_o8ubahqsaz.png")
	if err != nil {
		log.Fatal(fmt.Errorf("ошибка при скачивании изображения: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Errorf("ошибка HTTP при скачивании: %s", resp.Status))
	}

	respMedia, err := client.V1UploadMediaWithBodyWithResponse(context.Background(), "02", "file.png", resp.Body)
	if err != nil {
		log.Fatal(fmt.Errorf("ошибка при загрузке медиа: %w", err))
	}

	log.Println(respMedia.HTTPResponse.Status)
	log.Println(string(respMedia.Body))

	var (
		name        goord.CreativeName        = "example"
		brand       goord.CreativeBrand       = "brand"
		category    goord.CreativeCategory    = "category"
		description goord.CreativeDescription = "description"
		payType     goord.CreativePayType     = "cpm"
		form        goord.CreativeForm        = "text_graphic_block"
		targeting   goord.CreativeTargeting   = "Жители России"
		targetUrls  goord.CreativeTargetUrls  = []string{
			"https://example.com",
		}
		texts goord.CreativeTextsNullableTrue = []string{
			"example",
		}
		medias goord.CreativeMediaExternalIdsNullableTrue = []string{
			"02",
		}
	)

	respCreative, err := client.V3CreateCreativeWithResponse(context.Background(), "02", goord.V3CreateCreativeJSONRequestBody{
		Kktus:       []string{"30.15.1"},
		Name:        &name,
		Brand:       &brand,
		Category:    &category,
		Description: &description,
		PayType:     &payType,
		Form:        form,
		Targeting:   &targeting,
		TargetUrls:  &targetUrls,

		Texts:            &texts,
		MediaExternalIds: &medias,
		ContractExternalIds: []string{
			"4jvp4rojaeo-1jaocgov9",
		},
	})

	log.Println(respCreative.HTTPResponse.Status)
	log.Println(string(respCreative.Body))
	log.Println(respCreative.JSON200)
	log.Println(respCreative.JSON201)

}
