package model
import ("gopkg.in/mgo.v2/bson"
	"time"
)

type Thing struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	DisplayElementTopic string `json:"display_topic"`
	ControlElementTopic string `json:"control_topic"`
	UiElement string `json:"ui_element"`
	MaxValue float32  `json:"max_value"`
	MinValue float32  `json:"min_value"`
	Value string  `json:"value"`
	Unit string  `json:"unit"`
	UpdatedAt time.Time `json:"updated_at"`

}

type View struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Room string `json:"room"`
	Floor byte `json:"floor"`
	ZoneName string `json:"zone_name"`
	UiGroup string `json:"ui_group"`
	Things []Thing `json:"thing"`
}

type MobileUi struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Project bson.ObjectId `json:"project_id"`
	Views []View `json:"view"`
}

type Project struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `json:"name"`
	Domain string `json:"domain"`
	Comments string `json:"comments"`
	GeoLocation GeoLocation `json:"geo_location"`
}

type GeoLocation struct {
	Lat float32 `json:"lat"`
	Long float32 `json:"long"`
}
