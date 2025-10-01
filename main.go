package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func main() {
	log.Println("=== CRC ERAM Map Processor Starting ===")

	var inputARTCC string = "ZNY"
	flag.StringVar(&inputARTCC, "artcc", "", "ARTCC to get files for")
	flag.Parse()

	if inputARTCC == "" {
		log.Fatal("Error: ARTCC parameter is required. Use -artcc flag to specify ARTCC (e.g., ZNY)")
	}

	log.Printf("Processing ARTCC: %s", inputARTCC)

	// Assume were in the CRC directory.
	currentDir, _ := os.Getwd()
	log.Printf("Current working directory: %s", currentDir)

	artccDir := filepath.Join(currentDir, "ARTCCs", inputARTCC+".json")

	log.Printf("ARTCC file path: %s", artccDir)
	file, err := os.Open(artccDir)
	if err != nil {
		log.Fatalf("Error opening ARTCC file: %v", err)
	}
	defer file.Close()

	log.Println("Reading and parsing ARTCC file...")
	artcc := ARTCC{}
	err = json.NewDecoder(file).Decode(&artcc)
	if err != nil {
		log.Fatalf("Error parsing ARTCC JSON file: %v", err)
	}

	log.Printf("Successfully loaded ARTCC: %s (ID: %s)", artcc.Facility.Name, artcc.Facility.ID)

	output := ERAMMapGroups{}

	log.Printf("Found %d geomaps in ERAM configuration", len(artcc.Facility.EramConfiguration.GeoMaps))

	for i, geoMap := range artcc.Facility.EramConfiguration.GeoMaps {
		log.Printf("Processing geomap %d/%d: %s (ID: %s)", i+1, len(artcc.Facility.EramConfiguration.GeoMaps), geoMap.Name, geoMap.ID)
		log.Printf("  - Label: %s / %s", geoMap.LabelLine1, geoMap.LabelLine2)
		log.Printf("  - Video map count: %d", len(geoMap.VideoMapIds))
		log.Printf("  - BCG menu items: %d", len(geoMap.BcgMenu))
		group := ERAMMapGroup{}
		for j, filterMenu := range geoMap.FilterMenu {

			log.Printf("  Processing filter menu %d/%d: %s %s", j+1, len(geoMap.FilterMenu), filterMenu.LabelLine1, filterMenu.LabelLine2)

			// Skip unnamed/blank filters
			if filterMenu.LabelLine1 == "" && filterMenu.LabelLine2 == "" {
				continue
			}

			// Aggregate lines across all video maps for this filter
			var aggregatedLines [][]Point2LL
			bcg := ""
			// Prefer BCG label aligned with the filter index
			if j >= 0 && j < len(geoMap.BcgMenu) && geoMap.BcgMenu[j] != "" {
				bcg = geoMap.BcgMenu[j]
			}

			for _, videoMapID := range geoMap.VideoMapIds {
				// log.Printf("  Processing video map %d/%d: %s", i+1, len(geoMap.VideoMapIds), videoMapID)
				file, err := os.Open("C:/Users/Michael/AppData/Local/CRC/VideoMaps/ZNY/" + videoMapID + ".geojson")
				if err != nil {
					log.Fatalf("Error opening video map file %s: %v", videoMapID, err)
				}

				var gj GeoJSON
				err = json.NewDecoder(file).Decode(&gj)
				file.Close()
				if err != nil {
					log.Fatalf("Error decoding video map file %s: %v", videoMapID, err)
				}

				// Collect per-file defaults from special features
				lineDefaults := GeoJSONProperties{}
				for _, f := range gj.Features {
					if f.Properties == nil {
						continue
					}
					if f.Properties.IsLineDefaults {
						lineDefaults = *f.Properties
					}
					// Note: text/symbol defaults are not needed for line extraction
				}

				// Process features with fallback to defaults
				for _, feature := range gj.Features {
					if feature.Type != "Feature" {
						continue
					}
					// Skip defaults features themselves
					if feature.Properties != nil && (feature.Properties.IsLineDefaults || feature.Properties.IsTextDefaults || feature.Properties.IsSymbolDefaults) {
						continue
					}

					// Only extract lines for output
					if feature.Geometry.Type != "LineString" {
						// log.Printf("    Skipping non-LineString feature: %s. Current %v. Len %v", feature.Geometry.Type, k, len(gj.Features))
						continue
					}

					// Determine effective properties by applying defaults
					eff := GeoJSONProperties{}
					if feature.Properties != nil {
						eff = *feature.Properties
					}
					// Apply defaults where missing
					if eff.Bcg == 0 && lineDefaults.Bcg != 0 {
						eff.Bcg = lineDefaults.Bcg
					}
					if len(eff.Filters) == 0 && len(lineDefaults.Filters) != 0 {
						eff.Filters = append([]int(nil), lineDefaults.Filters...)
					}
					if eff.Style == "" && lineDefaults.Style != "" {
						eff.Style = lineDefaults.Style
					}
					if eff.Thickness == 0 && lineDefaults.Thickness != 0 {
						eff.Thickness = lineDefaults.Thickness
					}

					// Filter membership: require this filter index
					if !slices.Contains(eff.Filters, j) {
						continue
					}

					aggregatedLines = append(aggregatedLines, feature.Geometry.Coordinates)

					if eff.Bcg-1 >= 0 && eff.Bcg-1 < len(geoMap.BcgMenu) {
						// Only use element BCG if no filter-index BCG was set
						if bcg == "" {
							bcg = geoMap.BcgMenu[eff.Bcg-1]
						}
					}
				}
			}

			// Only append a map entry if we found any lines for this filter
			if len(aggregatedLines) > 0 {
				group.Maps = append(group.Maps, ERAMMap{
					BcgName:    bcg,
					LabelLine1: filterMenu.LabelLine1,
					LabelLine2: filterMenu.LabelLine2,
					Name:       geoMap.Name,
					Lines:      aggregatedLines,
				})
			}

		}
		output[geoMap.Name] = group
	}

	// Write the output to a file
	log.Println("Preparing to write output file...")
	log.Printf("Output contains %d geomap groups", len(output))

	// Calculate some statistics
	totalMaps := 0
	totalLines := 0
	for groupName, group := range output {
		log.Printf("  %s: %d maps", groupName, len(group.Maps))
		totalMaps += len(group.Maps)
		for _, mapItem := range group.Maps {
			totalLines += len(mapItem.Lines)
		}
	}
	log.Printf("Total maps processed: %d", totalMaps)
	log.Printf("Total LineString features extracted: %d", totalLines)

	outputFile, err := os.Create("output.json")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	log.Println("Writing output to JSON file...")
	err = json.NewEncoder(outputFile).Encode(output)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	log.Printf("✓ Output successfully written to %s", outputFile.Name())
	log.Println("=== CRC ERAM Map Processor Complete ===")
}
