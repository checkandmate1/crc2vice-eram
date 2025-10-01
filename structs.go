package main

import (
	"encoding/json"
	"time"
)

type Output map[string]ERAMMap // ARTCC Map Category -> GeoMaps

type ARTCC struct {
	ID            string    `json:"id"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	Facility      struct {
		ID              string `json:"id"`
		Type            string `json:"type"`
		Name            string `json:"name"`
		ChildFacilities []struct {
			ID              string `json:"id"`
			Type            string `json:"type"`
			Name            string `json:"name"`
			ChildFacilities []struct {
				ID                    string        `json:"id"`
				Type                  string        `json:"type"`
				Name                  string        `json:"name"`
				ChildFacilities       []interface{} `json:"childFacilities"`
				TowerCabConfiguration struct {
					VideoMapID                string `json:"videoMapId"`
					DefaultRotation           int    `json:"defaultRotation"`
					DefaultZoomRange          int    `json:"defaultZoomRange"`
					AircraftVisibilityCeiling int    `json:"aircraftVisibilityCeiling"`
					TowerLocation             struct {
						Lat float64 `json:"lat"`
						Lon float64 `json:"lon"`
					} `json:"towerLocation"`
				} `json:"towerCabConfiguration"`
				AsdexConfiguration struct {
					VideoMapID              string `json:"videoMapId"`
					DefaultRotation         int    `json:"defaultRotation"`
					DefaultZoomRange        int    `json:"defaultZoomRange"`
					TargetVisibilityRange   int    `json:"targetVisibilityRange"`
					TargetVisibilityCeiling int    `json:"targetVisibilityCeiling"`
					FixRules                []struct {
						ID            string `json:"id"`
						SearchPattern string `json:"searchPattern"`
						FixID         string `json:"fixId"`
					} `json:"fixRules"`
					UseDestinationIDAsFix bool `json:"useDestinationIdAsFix"`
					RunwayConfigurations  []struct {
						ID                   string        `json:"id"`
						Name                 string        `json:"name"`
						ArrivalRunwayIds     []string      `json:"arrivalRunwayIds"`
						DepartureRunwayIds   []string      `json:"departureRunwayIds"`
						HoldShortRunwayPairs []interface{} `json:"holdShortRunwayPairs"`
					} `json:"runwayConfigurations"`
					Positions []struct {
						ID        string        `json:"id"`
						Name      string        `json:"name"`
						RunwayIds []interface{} `json:"runwayIds"`
					} `json:"positions"`
					DefaultPositionID string `json:"defaultPositionId"`
					TowerLocation     struct {
						Lat float64 `json:"lat"`
						Lon float64 `json:"lon"`
					} `json:"towerLocation"`
				} `json:"asdexConfiguration,omitempty"`
				TdlsConfiguration struct {
					MandatorySid         bool `json:"mandatorySid"`
					MandatoryClimbout    bool `json:"mandatoryClimbout"`
					MandatoryClimbvia    bool `json:"mandatoryClimbvia"`
					MandatoryInitialAlt  bool `json:"mandatoryInitialAlt"`
					MandatoryDepFreq     bool `json:"mandatoryDepFreq"`
					MandatoryExpect      bool `json:"mandatoryExpect"`
					MandatoryContactInfo bool `json:"mandatoryContactInfo"`
					MandatoryLocalInfo   bool `json:"mandatoryLocalInfo"`
					Sids                 []struct {
						Name        string `json:"name"`
						ID          string `json:"id"`
						Transitions []struct {
							Name               string `json:"name"`
							ID                 string `json:"id"`
							FirstRoutePoint    string `json:"firstRoutePoint"`
							DefaultExpect      string `json:"defaultExpect"`
							DefaultDepFreq     string `json:"defaultDepFreq"`
							DefaultContactInfo string `json:"defaultContactInfo"`
							DefaultLocalInfo   string `json:"defaultLocalInfo"`
						} `json:"transitions"`
					} `json:"sids"`
					Climbouts []struct {
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"climbouts"`
					Climbvias []struct {
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"climbvias"`
					InitialAlts []struct {
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"initialAlts"`
					DepFreqs []struct {
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"depFreqs"`
					Expects []struct {
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"expects"`
					ContactInfos []struct {
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"contactInfos"`
					LocalInfos []struct {
						ID    string `json:"id"`
						Value string `json:"value"`
					} `json:"localInfos"`
				} `json:"tdlsConfiguration,omitempty"`
				FlightStripsConfiguration struct {
					StripBays []struct {
						ID            string `json:"id"`
						Name          string `json:"name"`
						NumberOfRacks int    `json:"numberOfRacks"`
					} `json:"stripBays"`
					ExternalBays []struct {
						FacilityID string `json:"facilityId"`
						BayID      string `json:"bayId"`
					} `json:"externalBays"`
					DisplayDestinationAirportIds bool `json:"displayDestinationAirportIds"`
					DisplayBarcodes              bool `json:"displayBarcodes"`
					EnableArrivalStrips          bool `json:"enableArrivalStrips"`
					EnableSeparateArrDepPrinters bool `json:"enableSeparateArrDepPrinters"`
					LockSeparators               bool `json:"lockSeparators"`
				} `json:"flightStripsConfiguration"`
				Positions []struct {
					ID                 string `json:"id"`
					Name               string `json:"name"`
					Starred            bool   `json:"starred"`
					RadioName          string `json:"radioName"`
					Callsign           string `json:"callsign"`
					Frequency          int    `json:"frequency"`
					StarsConfiguration struct {
						Subset   int    `json:"subset"`
						SectorID string `json:"sectorId"`
						AreaID   string `json:"areaId"`
						ColorSet string `json:"colorSet"`
						TCPID    string `json:"tcpId"`
					} `json:"starsConfiguration"`
					TransceiverIds []string `json:"transceiverIds"`
				} `json:"positions"`
				NeighboringFacilityIds []string      `json:"neighboringFacilityIds"`
				NonNasFacilityIds      []interface{} `json:"nonNasFacilityIds"`
			} `json:"childFacilities"`
			StarsConfiguration struct {
				Areas []struct {
					ID               string `json:"id"`
					Name             string `json:"name"`
					VisibilityCenter struct {
						Lat float64 `json:"lat"`
						Lon float64 `json:"lon"`
					} `json:"visibilityCenter"`
					SurveillanceRange       int      `json:"surveillanceRange"`
					UnderlyingAirports      []string `json:"underlyingAirports"`
					SsaAirports             []string `json:"ssaAirports"`
					TowerListConfigurations []struct {
						ID        string `json:"id"`
						AirportID string `json:"airportId"`
						Range     int    `json:"range"`
					} `json:"towerListConfigurations"`
					LdbBeaconCodesInhibited          bool `json:"ldbBeaconCodesInhibited"`
					PdbGroundSpeedInhibited          bool `json:"pdbGroundSpeedInhibited"`
					DisplayRequestedAltInFdb         bool `json:"displayRequestedAltInFdb"`
					UseVfrPositionSymbol             bool `json:"useVfrPositionSymbol"`
					ShowDestinationDepartures        bool `json:"showDestinationDepartures"`
					ShowDestinationSatelliteArrivals bool `json:"showDestinationSatelliteArrivals"`
					ShowDestinationPrimaryArrivals   bool `json:"showDestinationPrimaryArrivals"`
				} `json:"areas"`
				InternalAirports []string `json:"internalAirports"`
				BeaconCodeBanks  []struct {
					ID     string `json:"id"`
					Type   string `json:"type"`
					Subset int    `json:"subset"`
					Start  int    `json:"start"`
					End    int    `json:"end"`
				} `json:"beaconCodeBanks"`
				Rpcs []struct {
					ID                    string `json:"id"`
					Index                 int    `json:"index"`
					AirportID             string `json:"airportId"`
					PositionSymbolTie     string `json:"positionSymbolTie"`
					PositionSymbolStagger string `json:"positionSymbolStagger"`
					MasterRunway          struct {
						RunwayID             string `json:"runwayId"`
						HeadingTolerance     int    `json:"headingTolerance"`
						NearSideHalfWidth    int    `json:"nearSideHalfWidth"`
						FarSideHalfWidth     int    `json:"farSideHalfWidth"`
						NearSideDistance     int    `json:"nearSideDistance"`
						RegionLength         int    `json:"regionLength"`
						TargetReferencePoint struct {
							Lat float64 `json:"lat"`
							Lon float64 `json:"lon"`
						} `json:"targetReferencePoint"`
						TargetReferenceLineHeading   int `json:"targetReferenceLineHeading"`
						TargetReferenceLineLength    int `json:"targetReferenceLineLength"`
						TargetReferencePointAltitude int `json:"targetReferencePointAltitude"`
						ImageReferencePoint          struct {
							Lat float64 `json:"lat"`
							Lon float64 `json:"lon"`
						} `json:"imageReferencePoint"`
						ImageReferenceLineHeading float64       `json:"imageReferenceLineHeading"`
						ImageReferenceLineLength  int           `json:"imageReferenceLineLength"`
						TieModeOffset             float64       `json:"tieModeOffset"`
						DescentPointDistance      float64       `json:"descentPointDistance"`
						DescentPointAltitude      int           `json:"descentPointAltitude"`
						AbovePathTolerance        int           `json:"abovePathTolerance"`
						BelowPathTolerance        int           `json:"belowPathTolerance"`
						DefaultLeaderDirection    string        `json:"defaultLeaderDirection"`
						ScratchpadPatterns        []interface{} `json:"scratchpadPatterns"`
					} `json:"masterRunway"`
					SlaveRunway struct {
						RunwayID             string `json:"runwayId"`
						HeadingTolerance     int    `json:"headingTolerance"`
						NearSideHalfWidth    int    `json:"nearSideHalfWidth"`
						FarSideHalfWidth     int    `json:"farSideHalfWidth"`
						NearSideDistance     int    `json:"nearSideDistance"`
						RegionLength         int    `json:"regionLength"`
						TargetReferencePoint struct {
							Lat float64 `json:"lat"`
							Lon float64 `json:"lon"`
						} `json:"targetReferencePoint"`
						TargetReferenceLineHeading   float64 `json:"targetReferenceLineHeading"`
						TargetReferenceLineLength    int     `json:"targetReferenceLineLength"`
						TargetReferencePointAltitude int     `json:"targetReferencePointAltitude"`
						ImageReferencePoint          struct {
							Lat float64 `json:"lat"`
							Lon float64 `json:"lon"`
						} `json:"imageReferencePoint"`
						ImageReferenceLineHeading int           `json:"imageReferenceLineHeading"`
						ImageReferenceLineLength  int           `json:"imageReferenceLineLength"`
						TieModeOffset             float64       `json:"tieModeOffset"`
						DescentPointDistance      float64       `json:"descentPointDistance"`
						DescentPointAltitude      int           `json:"descentPointAltitude"`
						AbovePathTolerance        int           `json:"abovePathTolerance"`
						BelowPathTolerance        int           `json:"belowPathTolerance"`
						DefaultLeaderDirection    string        `json:"defaultLeaderDirection"`
						ScratchpadPatterns        []interface{} `json:"scratchpadPatterns"`
					} `json:"slaveRunway"`
				} `json:"rpcs"`
				PrimaryScratchpadRules []struct {
					ID            string   `json:"id"`
					AirportIds    []string `json:"airportIds"`
					SearchPattern string   `json:"searchPattern"`
					Template      string   `json:"template"`
					MinAltitude   int      `json:"minAltitude,omitempty"`
					MaxAltitude   int      `json:"maxAltitude,omitempty"`
				} `json:"primaryScratchpadRules"`
				SecondaryScratchpadRules  []interface{} `json:"secondaryScratchpadRules"`
				RnavPatterns              []interface{} `json:"rnavPatterns"`
				Allow4CharacterScratchpad bool          `json:"allow4CharacterScratchpad"`
				StarsHandoffIds           []struct {
					ID            string `json:"id"`
					FacilityID    string `json:"facilityId"`
					HandoffNumber int    `json:"handoffNumber"`
				} `json:"starsHandoffIds"`
				VideoMapIds []string `json:"videoMapIds"`
				MapGroups   []struct {
					ID     string        `json:"id"`
					MapIds []interface{} `json:"mapIds"`
					Tcps   []string      `json:"tcps"`
				} `json:"mapGroups"`
				AtpaVolumes []struct {
					ID              string `json:"id"`
					AirportID       string `json:"airportId"`
					VolumeID        string `json:"volumeId"`
					Name            string `json:"name"`
					RunwayThreshold struct {
						Lat float64 `json:"lat"`
						Lon float64 `json:"lon"`
					} `json:"runwayThreshold"`
					Ceiling                      int  `json:"ceiling"`
					Floor                        int  `json:"floor"`
					MagneticHeading              int  `json:"magneticHeading"`
					MaximumHeadingDeviation      int  `json:"maximumHeadingDeviation"`
					Length                       int  `json:"length"`
					WidthLeft                    int  `json:"widthLeft"`
					WidthRight                   int  `json:"widthRight"`
					TwoPointFiveApproachDistance int  `json:"twoPointFiveApproachDistance"`
					TwoPointFiveApproachEnabled  bool `json:"twoPointFiveApproachEnabled"`
					Scratchpads                  []struct {
						ID               string `json:"id"`
						Entry            string `json:"entry"`
						ScratchPadNumber string `json:"scratchPadNumber"`
						Type             string `json:"type"`
					} `json:"scratchpads"`
					Tcps []struct {
						ID       string `json:"id"`
						TCP      string `json:"tcp"`
						TCPID    string `json:"tcpId"`
						ConeType string `json:"coneType"`
					} `json:"tcps"`
					TCPExclusions    []interface{} `json:"tcpExclusions"`
					ExcludedTCPIds   []interface{} `json:"excludedTcpIds"`
					LeaderDirections []interface{} `json:"leaderDirections"`
				} `json:"atpaVolumes"`
				RecatEnabled           bool          `json:"recatEnabled"`
				Lists                  []interface{} `json:"lists"`
				ConfigurationPlans     []interface{} `json:"configurationPlans"`
				AutomaticConsolidation bool          `json:"automaticConsolidation"`
				Tcps                   []struct {
					Subset   int    `json:"subset"`
					SectorID string `json:"sectorId"`
					ID       string `json:"id"`
				} `json:"tcps"`
			} `json:"starsConfiguration,omitempty"`
			FlightStripsConfiguration struct {
				StripBays []struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					NumberOfRacks int    `json:"numberOfRacks"`
				} `json:"stripBays"`
				ExternalBays                 []interface{} `json:"externalBays"`
				DisplayDestinationAirportIds bool          `json:"displayDestinationAirportIds"`
				DisplayBarcodes              bool          `json:"displayBarcodes"`
				EnableArrivalStrips          bool          `json:"enableArrivalStrips"`
				EnableSeparateArrDepPrinters bool          `json:"enableSeparateArrDepPrinters"`
				LockSeparators               bool          `json:"lockSeparators"`
			} `json:"flightStripsConfiguration"`
			Positions []struct {
				ID                 string `json:"id"`
				Name               string `json:"name"`
				Starred            bool   `json:"starred"`
				RadioName          string `json:"radioName"`
				Callsign           string `json:"callsign"`
				Frequency          int    `json:"frequency"`
				StarsConfiguration struct {
					Subset   int    `json:"subset"`
					SectorID string `json:"sectorId"`
					AreaID   string `json:"areaId"`
					ColorSet string `json:"colorSet"`
					TCPID    string `json:"tcpId"`
				} `json:"starsConfiguration"`
				TransceiverIds []string `json:"transceiverIds"`
			} `json:"positions"`
			NeighboringFacilityIds []string      `json:"neighboringFacilityIds"`
			NonNasFacilityIds      []interface{} `json:"nonNasFacilityIds"`
			TowerCabConfiguration  struct {
				VideoMapID                string `json:"videoMapId"`
				DefaultRotation           int    `json:"defaultRotation"`
				DefaultZoomRange          int    `json:"defaultZoomRange"`
				AircraftVisibilityCeiling int    `json:"aircraftVisibilityCeiling"`
				TowerLocation             struct {
					Lat float64 `json:"lat"`
					Lon float64 `json:"lon"`
				} `json:"towerLocation"`
			} `json:"towerCabConfiguration,omitempty"`
			AsdexConfiguration struct {
				VideoMapID              string `json:"videoMapId"`
				DefaultRotation         int    `json:"defaultRotation"`
				DefaultZoomRange        int    `json:"defaultZoomRange"`
				TargetVisibilityRange   int    `json:"targetVisibilityRange"`
				TargetVisibilityCeiling int    `json:"targetVisibilityCeiling"`
				FixRules                []struct {
					ID            string `json:"id"`
					SearchPattern string `json:"searchPattern"`
					FixID         string `json:"fixId"`
				} `json:"fixRules"`
				UseDestinationIDAsFix bool `json:"useDestinationIdAsFix"`
				RunwayConfigurations  []struct {
					ID                   string        `json:"id"`
					Name                 string        `json:"name"`
					ArrivalRunwayIds     []string      `json:"arrivalRunwayIds"`
					DepartureRunwayIds   []string      `json:"departureRunwayIds"`
					HoldShortRunwayPairs []interface{} `json:"holdShortRunwayPairs"`
				} `json:"runwayConfigurations"`
				Positions []struct {
					ID        string        `json:"id"`
					Name      string        `json:"name"`
					RunwayIds []interface{} `json:"runwayIds"`
				} `json:"positions"`
				DefaultPositionID string `json:"defaultPositionId"`
				TowerLocation     struct {
					Lat float64 `json:"lat"`
					Lon float64 `json:"lon"`
				} `json:"towerLocation"`
			} `json:"asdexConfiguration,omitempty"`
			TdlsConfiguration struct {
				MandatorySid         bool `json:"mandatorySid"`
				MandatoryClimbout    bool `json:"mandatoryClimbout"`
				MandatoryClimbvia    bool `json:"mandatoryClimbvia"`
				MandatoryInitialAlt  bool `json:"mandatoryInitialAlt"`
				MandatoryDepFreq     bool `json:"mandatoryDepFreq"`
				MandatoryExpect      bool `json:"mandatoryExpect"`
				MandatoryContactInfo bool `json:"mandatoryContactInfo"`
				MandatoryLocalInfo   bool `json:"mandatoryLocalInfo"`
				Sids                 []struct {
					Name        string `json:"name"`
					ID          string `json:"id"`
					Transitions []struct {
						Name               string `json:"name"`
						ID                 string `json:"id"`
						DefaultExpect      string `json:"defaultExpect"`
						DefaultInitialAlt  string `json:"defaultInitialAlt"`
						DefaultContactInfo string `json:"defaultContactInfo"`
						DefaultLocalInfo   string `json:"defaultLocalInfo"`
					} `json:"transitions"`
				} `json:"sids"`
				Climbouts []struct {
					ID    string `json:"id"`
					Value string `json:"value"`
				} `json:"climbouts"`
				Climbvias   []interface{} `json:"climbvias"`
				InitialAlts []struct {
					ID    string `json:"id"`
					Value string `json:"value"`
				} `json:"initialAlts"`
				DepFreqs []struct {
					ID    string `json:"id"`
					Value string `json:"value"`
				} `json:"depFreqs"`
				Expects []struct {
					ID    string `json:"id"`
					Value string `json:"value"`
				} `json:"expects"`
				ContactInfos []struct {
					ID    string `json:"id"`
					Value string `json:"value"`
				} `json:"contactInfos"`
				LocalInfos []struct {
					ID    string `json:"id"`
					Value string `json:"value"`
				} `json:"localInfos"`
				DefaultSidID string `json:"defaultSidId"`
			} `json:"tdlsConfiguration,omitempty"`
		} `json:"childFacilities"`
		EramConfiguration struct {
			NasID   string `json:"nasId"`
			GeoMaps []struct {
				ID         string `json:"id"`
				Name       string `json:"name"`
				LabelLine1 string `json:"labelLine1"`
				LabelLine2 string `json:"labelLine2"`
				FilterMenu []struct {
					ID         string `json:"id"`
					LabelLine1 string `json:"labelLine1"`
					LabelLine2 string `json:"labelLine2"`
				} `json:"filterMenu"`
				BcgMenu     []string `json:"bcgMenu"`
				VideoMapIds []string `json:"videoMapIds"`
			} `json:"geoMaps"`
			EmergencyChecklist      []string `json:"emergencyChecklist"`
			PositionReliefChecklist []string `json:"positionReliefChecklist"`
			InternalAirports        []string `json:"internalAirports"`
			BeaconCodeBanks         []struct {
				ID       string `json:"id"`
				Category string `json:"category"`
				Priority string `json:"priority"`
				Subset   int    `json:"subset"`
				Start    int    `json:"start"`
				End      int    `json:"end"`
			} `json:"beaconCodeBanks"`
			NeighboringStarsConfigurations []struct {
				ID                     string `json:"id"`
				FacilityID             string `json:"facilityId"`
				StarsID                string `json:"starsId"`
				SingleCharacterStarsID string `json:"singleCharacterStarsId,omitempty"`
				FieldEFormat           string `json:"fieldEFormat"`
				FieldELetter           string `json:"fieldELetter,omitempty"`
			} `json:"neighboringStarsConfigurations"`
			NeighboringCaatsConfigurations []interface{} `json:"neighboringCaatsConfigurations"`
			CoordinationFixes              []interface{} `json:"coordinationFixes"`
			ReferenceFixes                 []string      `json:"referenceFixes"`
			AsrSites                       []struct {
				ID       string `json:"id"`
				AsrID    string `json:"asrId"`
				Location struct {
					Lat float64 `json:"lat"`
					Lon float64 `json:"lon"`
				} `json:"location"`
				Range   int `json:"range"`
				Ceiling int `json:"ceiling"`
			} `json:"asrSites"`
			ConflictAlertFloor int           `json:"conflictAlertFloor"`
			AirportSingleChars []interface{} `json:"airportSingleChars"`
		} `json:"eramConfiguration"`
		Positions []struct {
			ID                string `json:"id"`
			Name              string `json:"name"`
			Starred           bool   `json:"starred"`
			RadioName         string `json:"radioName"`
			Callsign          string `json:"callsign"`
			Frequency         int    `json:"frequency"`
			EramConfiguration struct {
				SectorID string `json:"sectorId"`
			} `json:"eramConfiguration"`
			TransceiverIds []string `json:"transceiverIds"`
		} `json:"positions"`
		NeighboringFacilityIds []string `json:"neighboringFacilityIds"`
		NonNasFacilityIds      []string `json:"nonNasFacilityIds"`
	} `json:"facility"`
	VisibilityCenters []struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"visibilityCenters"`
	AliasesLastUpdatedAt time.Time `json:"aliasesLastUpdatedAt"`
	VideoMaps            []struct {
		ID                      string    `json:"id"`
		Name                    string    `json:"name"`
		Tags                    []string  `json:"tags"`
		ShortName               string    `json:"shortName,omitempty"`
		SourceFileName          string    `json:"sourceFileName"`
		LastUpdatedAt           time.Time `json:"lastUpdatedAt"`
		StarsBrightnessCategory string    `json:"starsBrightnessCategory"`
		StarsID                 int       `json:"starsId,omitempty"`
		StarsAlwaysVisible      bool      `json:"starsAlwaysVisible"`
		TdmOnly                 bool      `json:"tdmOnly"`
	} `json:"videoMaps"`
	Transceivers []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Location struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"location"`
		HeightMslMeters int `json:"heightMslMeters"`
		HeightAglMeters int `json:"heightAglMeters"`
	} `json:"transceivers"`
	AutoAtcRules []struct {
		ID                string        `json:"id"`
		Status            string        `json:"status"`
		Name              string        `json:"name"`
		PositionID        string        `json:"positionId"`
		PrecursorRules    []interface{} `json:"precursorRules"`
		ExclusionaryRules []interface{} `json:"exclusionaryRules"`
		Criteria          struct {
			RouteSubstrings        []string      `json:"routeSubstrings"`
			ExcludeRouteSubstrings []interface{} `json:"excludeRouteSubstrings"`
			Departures             []interface{} `json:"departures"`
			Destinations           []string      `json:"destinations"`
			ApplicableToJets       bool          `json:"applicableToJets"`
			ApplicableToTurboprops bool          `json:"applicableToTurboprops"`
			ApplicableToProps      bool          `json:"applicableToProps"`
		} `json:"criteria"`
		DescentCrossingRestriction struct {
			CrossingFix        string `json:"crossingFix"`
			CrossingFixName    string `json:"crossingFixName"`
			AltitudeConstraint struct {
				Value           int    `json:"value"`
				TransitionLevel int    `json:"transitionLevel"`
				ConstraintType  string `json:"constraintType"`
				IsLufl          bool   `json:"isLufl"`
			} `json:"altitudeConstraint"`
			AltimeterStation struct {
				StationID   string `json:"stationId"`
				StationName string `json:"stationName"`
			} `json:"altimeterStation"`
		} `json:"descentCrossingRestriction,omitempty"`
		DescentRestriction struct {
			CrossingLine []struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			} `json:"crossingLine"`
			AltitudeConstraint struct {
				Value           int    `json:"value"`
				TransitionLevel int    `json:"transitionLevel"`
				ConstraintType  string `json:"constraintType"`
				IsLufl          bool   `json:"isLufl"`
			} `json:"altitudeConstraint"`
		} `json:"descentRestriction,omitempty"`
	} `json:"autoAtcRules"`
}

type Point2LL [2]float32

type GeoMap struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			IsLineDefaults bool   `json:"isLineDefaults"`
			Bcg            int    `json:"bcg"`
			Filters        []int  `json:"filters"`
			Style          string `json:"style"`
			Thickness      int    `json:"thickness"`
		} `json:"properties"`
	} `json:"features"`
}

type GeoJSON struct {
	Type     string           `json:"type"`
	Features []GeoJSONFeature `json:"features"`
}

type GeoJSONFeature struct {
	Type     string `json:"type"`
	Geometry struct {
		Type        string             `json:"type"`
		Coordinates GeoJSONCoordinates `json:"coordinates"`
	} `json:"geometry"`
}

// We only extract lines (at the moment at least) and so we only worry
// about [][2]float32s for coordinates. (For points, this would be
// a single [2]float32 and for polygons, it would be [][][2]float32...)
type GeoJSONCoordinates []Point2LL

func (c *GeoJSONCoordinates) UnmarshalJSON(d []byte) error {
	*c = nil

	var coords []Point2LL
	if err := json.Unmarshal(d, &coords); err == nil {
		*c = coords
	}
	// Don't report any errors but assume that it's a point, polygon, ...
	return nil
}

///////////////////////////////////////////////////////////////////////////

// Note: this should match STARSMap in stars.go
type ERAMMap struct {
	BcgName string
	LabelLine1 string
	LabelLine2 string
	Name  string
	Lines [][]Point2LL
}

type ERAMMapGroup []ERAMMap

type ERAMMapGroups map[string]ERAMMapGroup