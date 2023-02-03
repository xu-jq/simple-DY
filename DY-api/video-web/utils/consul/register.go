/*
 * @Date: 2023-01-27 16:24:44
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-03 10:48:29
 * @FilePath: /simple-DY/DY-api/video-web/utils/consul/register.go
 * @Description: Consul
 */
package consul

import (
	"strconv"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

type Registry struct {
	Host string
	Port string
}

type RegistryClient interface {
	Register(address string, port string, name string, tag []string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port string) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry) Register(address string, port string, name string, tag []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = r.Host + ":" + r.Port

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://" + address + ":" + port + "/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	portInt, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		panic(err)
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = uuid.NewV4().String()
	registration.Port = int(portInt)
	registration.Tags = tag
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = r.Host + ":" + r.Port

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}
