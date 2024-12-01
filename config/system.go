package config

type System struct {
	Addr        string `json:"addr" yaml:"addr"`
	Provisioner string `json:"provisioner" yaml:"provisioner"`
}

//type DataBase struct {
//	Host string `json:"host" yaml:"host"`
//	Port int    `json:"port" yaml:"port"`
//}
