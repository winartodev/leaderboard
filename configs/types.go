package configs

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
		Env  string `yaml:"env"`
	} `yaml:"app"`

	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}
