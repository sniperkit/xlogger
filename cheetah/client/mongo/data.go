package mongo

var (
	// DefaultConfig hols the singleton instace of the package configured parameters
	DefaultConfig *Configuration
)

// Configuration is the mongo datasource parameters
// mongo:
//    url: string
//    database: string
//    username: string
//    password: string
//    pool: int
//    timeout: int
type Configuration struct {
	URL      string `json:"url" mapstructure:"url"`
	Database string `json:"database" mapstructure:"database"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Pool     int    `json:"pool" mapstructure:"pool"`
	Timeout  int    `json:"timeout" mapstructure:"timeout"`
}

// Setup initialyzes the package
func Setup(config Configuration) error {
	DefaultConfig = &config
	return nil
}
