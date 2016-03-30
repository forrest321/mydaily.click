package waterdata

type WaterServiceResponse struct {
	DeclaredType    string `json:"declaredType"`
	GlobalScope     bool   `json:"globalScope"`
	Name            string `json:"name"`
	Nil             bool   `json:"nil"`
	Scope           string `json:"scope"`
	TypeSubstituted bool   `json:"typeSubstituted"`
	Value           struct {
		QueryInfo struct {
			CreationTime interface{} `json:"creationTime"`
			Criteria     struct {
				LocationParam string        `json:"locationParam"`
				MethodCalled  interface{}   `json:"methodCalled"`
				Parameter     []interface{} `json:"parameter"`
				TimeParam     interface{}   `json:"timeParam"`
				VariableParam string        `json:"variableParam"`
			} `json:"criteria"`
			Extension interface{} `json:"extension"`
			Note      []struct {
				Href  interface{} `json:"href"`
				Show  interface{} `json:"show"`
				Title string      `json:"title"`
				Type  interface{} `json:"type"`
				Value string      `json:"value"`
			} `json:"note"`
			QueryURL string `json:"queryURL"`
		} `json:"queryInfo"`
		TimeSeries []struct {
			Name       string `json:"name"`
			SourceInfo struct {
				Altname     interface{} `json:"altname"`
				ElevationM  interface{} `json:"elevationM"`
				Extension   interface{} `json:"extension"`
				GeoLocation struct {
					GeogLocation struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
						Srs       string  `json:"srs"`
					} `json:"geogLocation"`
					LocalSiteXY []interface{} `json:"localSiteXY"`
				} `json:"geoLocation"`
				MetadataTime interface{}   `json:"metadataTime"`
				Note         []interface{} `json:"note"`
				Oid          interface{}   `json:"oid"`
				SiteCode     []struct {
					AgencyCode string      `json:"agencyCode"`
					AgencyName interface{} `json:"agencyName"`
					Default    interface{} `json:"default"`
					Network    string      `json:"network"`
					SiteID     interface{} `json:"siteID"`
					Value      string      `json:"value"`
				} `json:"siteCode"`
				SiteName     string `json:"siteName"`
				SiteProperty []struct {
					Name  string      `json:"name"`
					Type  interface{} `json:"type"`
					URI   interface{} `json:"uri"`
					Value string      `json:"value"`
				} `json:"siteProperty"`
				SiteType     []interface{} `json:"siteType"`
				TimeZoneInfo struct {
					DaylightSavingsTimeZone struct {
						ZoneAbbreviation string `json:"zoneAbbreviation"`
						ZoneOffset       string `json:"zoneOffset"`
					} `json:"daylightSavingsTimeZone"`
					DefaultTimeZone struct {
						ZoneAbbreviation string `json:"zoneAbbreviation"`
						ZoneOffset       string `json:"zoneOffset"`
					} `json:"defaultTimeZone"`
					SiteUsesDaylightSavingsTime bool `json:"siteUsesDaylightSavingsTime"`
				} `json:"timeZoneInfo"`
				VerticalDatum interface{} `json:"verticalDatum"`
			} `json:"sourceInfo"`
			Values []struct {
				CensorCode []interface{} `json:"censorCode"`
				Method     []struct {
					MethodCode        interface{} `json:"methodCode"`
					MethodDescription string      `json:"methodDescription"`
					MethodID          int         `json:"methodID"`
					MethodLink        interface{} `json:"methodLink"`
				} `json:"method"`
				Offset    []interface{} `json:"offset"`
				Qualifier []struct {
					Default              interface{} `json:"default"`
					Network              string      `json:"network"`
					QualifierCode        string      `json:"qualifierCode"`
					QualifierDescription string      `json:"qualifierDescription"`
					QualifierID          int         `json:"qualifierID"`
					Vocabulary           string      `json:"vocabulary"`
				} `json:"qualifier"`
				QualityControlLevel []interface{} `json:"qualityControlLevel"`
				Sample              []interface{} `json:"sample"`
				Source              []interface{} `json:"source"`
				Units               interface{}   `json:"units"`
				Value               []struct {
					AccuracyStdDev          interface{} `json:"accuracyStdDev"`
					CensorCode              interface{} `json:"censorCode"`
					CodedVocabulary         interface{} `json:"codedVocabulary"`
					CodedVocabularyTerm     interface{} `json:"codedVocabularyTerm"`
					DateTime                string      `json:"dateTime"`
					DateTimeAccuracyCd      interface{} `json:"dateTimeAccuracyCd"`
					DateTimeUTC             interface{} `json:"dateTimeUTC"`
					LabSampleCode           interface{} `json:"labSampleCode"`
					MetadataTime            interface{} `json:"metadataTime"`
					MethodCode              interface{} `json:"methodCode"`
					MethodID                interface{} `json:"methodID"`
					OffsetTypeCode          interface{} `json:"offsetTypeCode"`
					OffsetTypeID            interface{} `json:"offsetTypeID"`
					OffsetValue             interface{} `json:"offsetValue"`
					Oid                     interface{} `json:"oid"`
					Qualifiers              []string    `json:"qualifiers"`
					QualityControlLevelCode interface{} `json:"qualityControlLevelCode"`
					SampleID                interface{} `json:"sampleID"`
					SourceCode              interface{} `json:"sourceCode"`
					SourceID                interface{} `json:"sourceID"`
					TimeOffset              interface{} `json:"timeOffset"`
					Value                   string      `json:"value"`
				} `json:"value"`
			} `json:"values"`
			Variable struct {
				Categories      interface{}   `json:"categories"`
				DataType        interface{}   `json:"dataType"`
				Extension       interface{}   `json:"extension"`
				GeneralCategory interface{}   `json:"generalCategory"`
				MetadataTime    interface{}   `json:"metadataTime"`
				NoDataValue     int           `json:"noDataValue"`
				Note            []interface{} `json:"note"`
				Oid             string        `json:"oid"`
				Options         struct {
					Option []struct {
						Name       string      `json:"name"`
						OptionCode string      `json:"optionCode"`
						OptionID   interface{} `json:"optionID"`
						Value      interface{} `json:"value"`
					} `json:"option"`
				} `json:"options"`
				Related      interface{} `json:"related"`
				SampleMedium interface{} `json:"sampleMedium"`
				Speciation   interface{} `json:"speciation"`
				TimeScale    interface{} `json:"timeScale"`
				Unit         struct {
					UnitAbbreviation string      `json:"unitAbbreviation"`
					UnitCode         interface{} `json:"unitCode"`
					UnitDescription  interface{} `json:"unitDescription"`
					UnitID           interface{} `json:"unitID"`
					UnitName         interface{} `json:"unitName"`
					UnitType         interface{} `json:"unitType"`
				} `json:"unit"`
				ValueType    string `json:"valueType"`
				VariableCode []struct {
					Default    bool   `json:"default"`
					Network    string `json:"network"`
					Value      string `json:"value"`
					VariableID int    `json:"variableID"`
					Vocabulary string `json:"vocabulary"`
				} `json:"variableCode"`
				VariableDescription string        `json:"variableDescription"`
				VariableName        string        `json:"variableName"`
				VariableProperty    []interface{} `json:"variableProperty"`
			} `json:"variable"`
		} `json:"timeSeries"`
	} `json:"value"`
}
