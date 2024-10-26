package settings

type Infrastructure struct {
	Postgres    Postgres    `yaml:"postgres"`
	GoogleCloud GoogleCloud `yaml:"google_cloud"`
}

type GoogleCloud struct {
	ProjectID           string `yaml:"project_id"`
	UseCredentialsFile  bool   `yaml:"use_credentials_file"`
	CredentialsFilePath string `yaml:"credentials_file_path"`
}

type Postgres struct {
	Protocol   string `yaml:"protocol"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	UnixSocket string `yaml:"unix_socket"`
	User       string `yaml:"username"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"database"`
}
