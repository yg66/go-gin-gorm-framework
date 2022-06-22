package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yg66/go-gin-gorm-framework/model"
)

func InitConfig() (*model.Config, error) {
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.1.233",
			Port:   8848,
		},
	}

	// Create clientConfig
	clientConfig := constant.ClientConfig{
		// If you need to support multiple namespaces, you can scenario multiple clients with different Namespaceids. If the namespace is public, fill in the blank string.
		NamespaceId:         "839108c6-58d6-421f-b68e-d2f38be1d18f",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// Another way to create a dynamically configured client (recommended)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	// Obtaining Configuration Information
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "investors-locked.yaml",
		Group:  "DEFAULT_GROUP"})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetConfig err: %v", err))
	}
	config := model.Config{
		BasicsConfig:    &model.BasicsConfig{},
		DbConfig:        &model.DbConfig{},
		ReCAPTCHAConfig: &model.ReCAPTCHAConfig{},
		JobConfig:       &model.JobConfig{},
		SmtpConfig:      &model.SmtpConfig{},
	}
	err = json.Unmarshal([]byte(content), &config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Read config failed: %v", err))
	}

	// Listening to the configuration
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "investors-locked.yaml",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		return nil, err
	}
	//time.Sleep(time.Second * 1000)
	return &config, nil
}
