package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var inputARTCC string
	flag.StringVar(&inputARTCC, "input", "", "Path to ARTCC JSON file")
	flag.Parse()

	// Assume were in the CRC directory.
	currentDir, _ := os.Getwd()
	artccDir := filepath.Join(currentDir, "ARTCCs", inputARTCC+".json")
	file, err := os.Open(artccDir)
	if err != nil {
		log.Fatalf("Error opening ARTCC file: %v", err)
	}
	defer file.Close()

	artcc := ARTCC{}
	json.NewDecoder(file).Decode(&artcc)

	output := ERAMMapGroups{}

	for _, geoMap := range artcc.Facility.EramConfiguration.GeoMaps {
		group := ERAMMapGroup{}
		for i, videoMapID := range geoMap.VideoMapIds {
			// Get the video map file and load it:
			videoMapFile, err := os.Open(filepath.Join(currentDir, "VideoMaps", inputARTCC, videoMapID+".geojson"))
			if err != nil {
				log.Fatalf("Error opening video map file: %v", err)
			}
			defer videoMapFile.Close()

			h := ERAMMap{
				BcgName:    geoMap.BcgMenu[i],
				LabelLine1: geoMap.LabelLine1,
				LabelLine2: geoMap.LabelLine2,
				Name:       geoMap.Name,
			}

			var gj GeoJSON
			body, err := io.ReadAll(videoMapFile)
			if err != nil {
				log.Fatalf("Error reading video map file: %v", err)
			}
			err = UnmarshalJSON(body, &gj)
			if err != nil {
				log.Fatalf("Error unmarshalling video map file: %v", err)
			}

			for _, feature := range gj.Features {
				if feature.Type != "Feature" {
					continue
				}
				if feature.Geometry.Type != "LineString" {
					continue
				}
				h.Lines = append(h.Lines, feature.Geometry.Coordinates)
			}
			group = append(group, h)
		}
		output[geoMap.Name] = group
	}
}
