package pmodels

// EnvDoc Variables de documentascion
type EnvDoc struct {
	Enable  bool
	URL     string
	Port    string
	Token   string
	Timeout int
}

// EnvGeneral Variables generales
type EnvGeneral struct {
	Produccion     bool
	Panic          bool
	ContextTime    int
	ClaveSecreta   string
	TiempoToken    int
	DBASTRIC       string
	Notificaciones bool
	DisableToken   bool
	DisableConsola bool
	TimeoutClient  int
	TimeoutPingDB  int
	ServerIP       string
}

// EnvAPI Variables Api rest
type EnvAPI struct {
	Enable  bool
	Port    string
	Timeout int
}

// EnvWs Variables Wensocket
type EnvWs struct {
	Enable  bool
	Port    string
	Timeout int
}

// EnvMongo Variables de mongo
type EnvMongo struct {
	Enable bool
	StrCon string
}

// EnvMysql Variables de Mysql
type EnvMysql struct {
	Enable             bool
	User               string
	Pass               string
	Host               string
	Db                 string
	Port               int
	Logger             bool
	SetConnMaxLifetime int
	SetMaxOpenConns    int
	SetMaxIdleConns    int
	SetMaxTimeOutCon   int
}

// EnvMails Varables mails
type EnvMails struct {
	Enable  bool
	From    string
	Smtpstr string
	Port    int
	Address string
	Pass    string
}
