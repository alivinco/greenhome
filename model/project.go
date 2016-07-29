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
	Value interface{}  `json:"value"`
	Unit string  `json:"unit"`
	UpdatedAt time.Time `json:"updated_at"`
	PropFieldForUi string `json:"prop_field_ui"`

}

type View struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	Room string `json:"room"`
	Floor byte `json:"floor"`
	ZoneName string `json:"zone_name"`
	UiGroup string `json:"ui_group"`
	Things []Thing `json:"thing"`
}

type Project struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Domain string `json:"domain"`
	Comments string `json:"comments"`
	GeoLocation GeoLocation `json:"geo_location"`
	Views []View `json:"view"`
}

type GeoLocation struct {
	Lat float32 `json:"lat"`
	Long float32 `json:"long"`
}

