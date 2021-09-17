package direccion

import (
	"regexp"
)

type Direccion struct {
	ip      string
	address string
	Ipad
}

func (dir Direccion) comprobar_ip_valida() (bool, error) {
	expresion_ip := "^[0-255].[0-255].[0-255].[0-255]$"
	return regexp.MatchString(expresion_ip, dir.ip)
}
