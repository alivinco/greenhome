package model

type AppConfigs struct {
	Auth0Domain string	 `env:"GRH_AUTH0_DOMAIN" envDefault:""`
	AppRootUrl string 	 `env:"GRH_APP_ROOT_URL" envDefault:""`
	AuthClientId string 	`env:"GRH_AUTH_CLIENT_ID" envDefault:""`
	AuthClientSecret string `env:"GRH_AUTH_CLIENT_SECRET" envDefault:""`
	JwtSecret string	 `env:"GRH_AUTH_JWT_SECRET" envDefault:""`
	BindAddress string	 `env:"GRH_BIND_ADDR" envDefault:":5010"`
	VarDirPath string
	MqttBrokerUri string	 `env:"GRH_MQTT_BROKER_URI" envDefault:"tcp://localhost:1883"`
	MqttConnClientId string  `env:"GRH_MQTT_CONN_CLIENT_ID" envDefault:"greenhome"`
	MqttConnUsername string	 `env:"GRH_MQTT_CONN_USERNAME" envDefault:""`
	MqttConnPassword string	 `env:"GRH_MQTT_CONN_PASS" envDefault:""`
	MongoConnUri string	 `env:"GRH_MONGO_CONN_URI" envDefault:"localhost"`
	SessionStoreFile string  `env:"GRH_SESSION_STORE" envDefault:"./sessionsecret.db"`
	LogDir string 		 `env:"GRH_LOG_DIR" envDefault:"./"`
}
