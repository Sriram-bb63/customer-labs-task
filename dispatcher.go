package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

func dispatch(data map[string]any) {
	log.Println("Dispatcher called")
	outputPayload := make(map[string]any)

	outputPayload["event"] = data["ev"].(string)
	outputPayload["event_type"] = data["et"].(string)
	outputPayload["app_id"] = data["id"].(string)
	outputPayload["user_id"] = data["uid"].(string)
	outputPayload["message_id"] = data["mid"].(string)
	outputPayload["page_title"] = data["t"].(string)
	outputPayload["page_url"] = data["p"].(string)
	outputPayload["browser_language"] = data["l"].(string)
	outputPayload["screen_size"] = data["sc"].(string)
	outputPayload["attributes"] = make(map[string]map[string]string)
	outputPayload["traits"] = make(map[string]map[string]string)
	attributesMap := outputPayload["attributes"].(map[string]map[string]string)
	traitsMap := outputPayload["traits"].(map[string]map[string]string)

	attributePattern := regexp.MustCompile(`^atrk(\d+)$`)
	traitPattern := regexp.MustCompile(`^uatrk(\d+)$`)
	for key, val := range data {
		if isAttribute := attributePattern.MatchString(key); isAttribute {
			log.Println("ATTRIBUTE FOUND")
			attributeNum := key[4:]
			attributeKey := val.(string)
			attributeValue := data[fmt.Sprintf("atrv%v", attributeNum)].(string)
			attributeType := data[fmt.Sprintf("atrt%v", attributeNum)].(string)
			log.Println(attributeKey, attributeValue, attributeType)
			attributesMap[attributeKey] = map[string]string{
				"value": attributeValue,
				"type":  attributeType,
			}
		}
		if isTrait := traitPattern.MatchString(key); isTrait {
			log.Println("TRAIT FOUND")
			traitNum := key[5:]
			traitKey := val.(string)
			traitValue := data[fmt.Sprintf("uatrv%v", traitNum)].(string)
			traitType := data[fmt.Sprintf("uatrt%v", traitNum)].(string)
			log.Println(traitKey, traitValue, traitType)
			traitsMap[traitKey] = map[string]string{
				"value": traitValue,
				"type":  traitType,
			}
		}
	}
	jsonData, err := json.MarshalIndent(outputPayload, "", "  ")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	log.Println(string(jsonData))

}
