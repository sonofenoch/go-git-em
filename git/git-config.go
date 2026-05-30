package git 

type Config struct {
	username string `json:username`
	email string `json:email`
	default_branch string `json:default_branch`
}

func ReadConfig(config string = "/home/enoch/.gogitemconfig.json") (*Config, error) {
	var err error
	configs, err := os.ReadFile(config)
	if err != nil {
		return nil, fmt.Errorf("Could not open config file for reading: %s\n", config) 
	}
	conf := Config{}

	err := json.Unmarshal(configs, &conf)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal config data") 
	}
	
	return conf, nil
}

