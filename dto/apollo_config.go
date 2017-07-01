package dto

import (
	"encoding/json"
	"github.com/zouyx/agollo/utils/objectutils"
)

type ApolloConfig struct {
	AppId string `json:"appId"`
	Cluster string `json:"cluster"`
	NamespaceName string `json:"namespaceName"`
	Configurations map[string]string `json:"configurations"`
	ReleaseKey string `json:"releaseKey"`
}

func CreateApolloConfigWithJson(str string) (*ApolloConfig,error) {
	apolloConfig:=&ApolloConfig{}
	err:=json.Unmarshal([]byte(str),apolloConfig)
	if objectutils.IsNotNil(err) {
		return nil,err
	}
	return apolloConfig,nil
}