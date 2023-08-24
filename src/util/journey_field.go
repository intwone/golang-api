package util

import "go.uber.org/zap"

func CreateJourneyField(journey string) zap.Field {
	return zap.String("journey", journey)
}
