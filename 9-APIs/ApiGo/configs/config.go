package configs

// importando o viper
import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// Configurações
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExperesIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

// Função onde será colocado um arquivo de configuração
func LoadConfig(path string) (*conf, error) {
	// Criando uma variável
	var cfg *conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	// Criando uma instância do JWTAuth e passando o algoritmo de criptografia e a chave secreta
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, err
}
