package samad


type Config struct {
	Username      string
	Password      string
	GetTokenUrl   string
	GetProgramUrl string
	ReserveUrl    string
}

// func GetConf() (*config.Config, error) {

// 	conf, err := config.New()

// 	if err != nil {
// 		return nil, err
// 	}

// 	conf.Username = config.GetEnv("username", "")
// 	conf.Password = config.GetEnv("password", "")
// 	conf.GetProgramUrl = config.GetEnv("GetProgramUrl", "")
// 	conf.GetTokenUrl = config.GetEnv("GetTokenUrl", "")
// 	conf.ReserveUrl = config.GetEnv("ReserveUrl", "")

// 	return conf, nil
// }
