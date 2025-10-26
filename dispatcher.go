package main

import (
	"fmt"
	"log"
	"regexp"
)

func dispatch(data map[string]string) {
	log.Println("Dispatcher called")
	outputPayload := make(map[string]any)

	outputPayload["event"] = data["ev"]
	outputPayload["event_type"] = data["et"]
	outputPayload["app_id"] = data["id"]
	outputPayload["user_id"] = data["uid"]
	outputPayload["message_id"] = data["mid"]
	outputPayload["page_title"] = data["t"]
	outputPayload["page_url"] = data["p"]
	outputPayload["browser_language"] = data["l"]
	outputPayload["screen_size"] = data["sc"]
	outputPayload["attributes"] = make(map[string]map[string]string)
	outputPayload["traits"] = make(map[string]map[string]string)
	attributesMap := outputPayload["attributes"].(map[string]map[string]string)
	traitsMap := outputPayload["traits"].(map[string]map[string]string)

	attributePattern := regexp.MustCompile(`^atrk(\d+)$`)
	traitPattern := regexp.MustCompile(`^uatrk(\d+)$`)
	for key, val := range data {
		if isAttribute := attributePattern.MatchString(key); isAttribute {
			attributeNum := key[4:]
			attributeKey := val
			attributeValue := data[fmt.Sprintf("atrv%v", attributeNum)]
			attributeType := data[fmt.Sprintf("atrt%v", attributeNum)]
			attributesMap[attributeKey] = map[string]string{
				"value": attributeValue,
				"type":  attributeType,
			}
		}
		if isTrait := traitPattern.MatchString(key); isTrait {
			traitNum := key[5:]
			traitKey := val
			traitValue := data[fmt.Sprintf("uatrv%v", traitNum)]
			traitType := data[fmt.Sprintf("uatrt%v", traitNum)]
			traitsMap[traitKey] = map[string]string{
				"value": traitValue,
				"type":  traitType,
			}
		}
	}
}
