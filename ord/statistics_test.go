package ord

import (
	"context"
	"testing"
)

func TestClient_Statistics(t *testing.T) {
	client, _ := NewClient(
		WithBase("https://api-sandbox.ord.vk.com"),
		WithToken("test-token"),
	)

	t.Run("CreateStatisticsV2", func(t *testing.T) {
		statistics := StatisticsV2ItemsArray{
			Items: []StatisticsV2Item{
				{
					CreativeExternalID: "creative-1",
					PadExternalID:      "pad-1",
					ShowsCount:         100,
					DateStartActual:    "2023-01-01",
					DateEndActual:      "2023-01-31",
				},
			},
		}

		_, err := client.CreateStatisticsV2(context.Background(), statistics)
		if err == nil {
			t.Error("Expected error due to invalid token, got nil")
		}
	})

	t.Run("CreateStatisticsV3", func(t *testing.T) {
		statistics := StatisticsV3ItemsArray{
			Items: []StatisticsV3Item{
				{
					StatisticsV2Item: StatisticsV2Item{
						CreativeExternalID: "creative-1",
						PadExternalID:      "pad-1",
						ShowsCount:         100,
						DateStartActual:    "2023-01-01",
						DateEndActual:      "2023-01-31",
					},
				},
			},
		}

		_, err := client.CreateStatisticsV3(context.Background(), statistics)
		if err == nil {
			t.Error("Expected error due to invalid token, got nil")
		}
	})

	t.Run("GetStatisticsList", func(t *testing.T) {
		_, err := client.GetStatisticsList(context.Background(), 0, 10)
		if err == nil {
			t.Error("Expected error due to invalid token, got nil")
		}
	})

	t.Run("DeleteStatisticsV3", func(t *testing.T) {
		deleteReq := DeleteStatisticsRequest{
			Items: []struct {
				CreativeExternalID string `json:"creative_external_id"`
				PadExternalID      string `json:"pad_external_id"`
				DateStartActual    string `json:"date_start_actual"`
			}{
				{
					CreativeExternalID: "creative-1",
					PadExternalID:      "pad-1",
					DateStartActual:    "2023-01-01",
				},
			},
		}

		err := client.DeleteStatisticsV3(context.Background(), deleteReq)
		if err == nil {
			t.Error("Expected error due to invalid token, got nil")
		}
	})
}
