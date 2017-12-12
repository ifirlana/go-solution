package envvar

import(
	"github.com/pkg/errors"
	"github.com/kelseyhightower/envconfig"
	"os"
	"encoding/json"
)

func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "error loading config file")
		}
	}
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from loading env")
}

func LoadFile(path string, config interface{}) error  {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(config);err != nil {
		return errors.Wrap(err, "failed to decoded config file")
	}

	return nil
}