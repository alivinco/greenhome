package json_types

type IotMsgCmdV0 struct {
	Origin struct {
		ID string `json:"@id"`
		Vendor string `json:"vendor"`
		Type string `json:"@type"`
		Location string `json:"location"`
	} `json:"origin"`
	UUID string `json:"uuid"`
	CreationTime int64 `json:"creation_time"`
	Command struct {
		Default struct {
			Value interface{} `json:"value"`
			Unit string `json:"unit"`
		} `json:"default"`
		Subtype string `json:"subtype"`
		Target string `json:"target"`
		Properties map[string]interface{} `json:"properties"`
		Type string `json:"@type"`
	} `json:"command"`
	Spid string `json:"spid"`
	Corid string `json:"corid"`
	Context string `json:"@context"`
	Transport string `json:"transport"`
}

type IotMsgEvtV0 struct {
	Origin struct {
		ID string `json:"@id"`
		Vendor string `json:"vendor"`
		Type string `json:"@type"`
		Location string `json:"location"`
	} `json:"origin"`
	UUID string `json:"uuid"`
	CreationTime int64 `json:"creation_time"`
	Event struct {
		Default struct {
			Value interface{} `json:"value"`
			Unit string `json:"unit"`
		} `json:"default"`
		Subtype string `json:"subtype"`
		Target string `json:"target"`
		Properties map[string]interface{} `json:"properties"`
		Type string `json:"@type"`
	} `json:"event"`
	Spid string `json:"spid"`
	Corid string `json:"corid"`
	Context string `json:"@context"`
	Transport string `json:"transport"`
}

type IotMsgV0 struct {
	Origin struct {
		ID string `json:"@id"`
		Vendor string `json:"vendor"`
		Type string `json:"@type"`
		Location string `json:"location"`
	} `json:"origin"`
	UUID string `json:"uuid"`
	CreationTime int64 `json:"creation_time"`
	Command struct {
		Default struct {
			Value interface{} `json:"value"`
			Unit string `json:"unit"`
		} `json:"default"`
		Subtype string `json:"subtype"`
		Target string `json:"target"`
		Properties map[string]interface{} `json:"properties"`
		Type string `json:"@type"`
	} `json:"command"`
	Event struct {
		Default struct {
			Value interface{} `json:"value"`
			Unit string `json:"unit"`
		} `json:"default"`
		Subtype string `json:"subtype"`
		Target string `json:"target"`
		Properties map[string]interface{} `json:"properties"`
		Type string `json:"@type"`
	} `json:"event"`
	Spid string `json:"spid"`
	Corid string `json:"corid"`
	Context string `json:"@context"`
	Transport string `json:"transport"`
}