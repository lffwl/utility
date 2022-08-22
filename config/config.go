package config

import (
	"fmt"
	"github.com/lffwl/utility/file"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// LoadConfig 加载配置
func LoadConfig[T any](path string, conf T) error {

	// 验证文件是否存在
	if !file.IsExist(path) {
		return fmt.Errorf("LoadConfig error ： path : %s 不存在", path)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return err
	}
	return nil
}
