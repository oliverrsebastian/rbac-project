package file

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"rbac-project/authorize"
	"time"
)

const (
	casbinModelPath  = "./config/casbin_model.conf"
	casbinPolicyPath = "./config/rbac_policy.csv"
)

type fileAccessManager struct {
	enforcer casbin.IEnforcer
}

func NewFileAccessManager() authorize.AccessManager {
	fileEnforcer, err := casbin.NewSyncedEnforcer(casbinModelPath, casbinPolicyPath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	
	fileEnforcer.StartAutoLoadPolicy(time.Duration(10) * time.Second)

	return &fileAccessManager{fileEnforcer}
}

func (am *fileAccessManager) Check(subject, resource, permission string) (bool, error) {
	return am.enforcer.Enforce(subject, resource, permission)
}
