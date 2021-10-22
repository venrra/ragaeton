package direccion

import (
	"regexp"
)

func comprobar_ip_valida(ip string) (bool, error) {
	expresion_ip := "^[0-255].[0-255].[0-255].[0-255]$"
	return regexp.MatchString(expresion_ip, ip)
}


func comprobar_mask_valida(ip string) (bool, error) {
	expresion_ip := "^(0,255).(0,255).(0,255).(0,255)$"
	return regexp.MatchString(expresion_ip, mask)
}

type Direccion struct {
	ip      string
	mask	string
	address string
}

type Opcion func(*Direccion)

func Ip(ip string) Opcion {
	return func(dir *Direccion){
		if comprobar_ip_valida(ip)
			dir.ip = ip
	}
}

func Mask(mask string) Opcion {
	return func(dir *Direccion){
		if comprobar_mask_valida(mask)
			dir.mask = mask
	}
}

func Address(add string) Opcion {
	return func(dir *Direccion){
		if !comprobar_ip_valida(add)
			dir.add = add
	}
}


func NewUsuario(opciones ...Opcion) *Direccion {
	dir := Direccion{ip: "0.0.0.0", apellido: "0.0.0.0", Address: "N/A"}
	for _, opcion := range opciones {
		opcion(&dir)
	}
	return &dir
}

func (dir Direccion) comprobar_ip_valida() (bool, error) {
	expresion_ip := "^[0-255].[0-255].[0-255].[0-255]$"
	return regexp.MatchString(expresion_ip, dir.ip)
}


func (dir Direccion) GetIp() string{
	return dir.ip
}

func (dir *Direccion) SetIp(ip string){
	(dir):=Ip(ip)
}
