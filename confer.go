package confer


import (
	"os"
	"errors"
	"github.com/BurntSushi/toml"
)



func LoadConfig(paths []string, conf interface{}) (err error){
	if len(os.Args) > 1 {
		paths = append(os.Args[1:2], paths...)
	}
	err = errors.New("No path was passed to find the configuration file")
	for _, path := range paths {
		_, err = toml.DecodeFile(path, conf)
		if err == nil {
			return
		}
	}
	return
}
