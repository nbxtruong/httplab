package main

import (
	"encoding/json"
	"log"
	"os"
)

// Helper function to import json from file to map
func importJSONFromFile(fileName string, result interface{}) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to reading file, err %v", err)
	}
	
	err = json.Unmarshal(content, result)
	if err != nil {
		log.Fatalf("Failed to unmarshal file, err %v", err)
	}
}
