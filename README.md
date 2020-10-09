# Ragaeton

Proyecto Infraestructura Virtual

## :memo: Descripción

Lo que se pretende con este micro-servicio es dar facilidades para analizar y evaluar el estado de sitios Webs o sitios accesibles desde *Internet*, analizando así, el estado de los puertos o evaluando posibles amenazas o vulnerabilidades que el sitio en cuestión pueda sufrir.

## Desarrollo del proyecto

El objetivo del proyecto es crear una API REST, que nos de información de direcciones de internet. Se obtendrá una pieza de software que estará lista para desplegarse en la nube, haciendo uso de técnicas DevOps y desarrollo ágil.

### ¿Por que ragaeton?

Como ya se hado el objetivo es desarrollar una api que nos permita conocer el estado de *sitios de internet*, para poder de forma rápida, automatizada y fácil, hacer un pequeño análisis de cualquier sitio, por ejemplo si estamos en desarrollo de una web comprobar si esta activa o si tiene alguna vulnerabilidad letal. 

También se pretende conseguir poner en practica conceptos y conocimientos de ciberseguridad, auditorias web y redes en general.  
Mas personalmente, intento enfocar mis estudios y conocimientos en inmensa rama que es la seguridad y ciberseguridad informáticas.

En cuanto al nombre, no tenia muy claro que elegir y opte por algo que no se olvidara fácilmente y llamase la atención. Como versión oficial es que ragaeton tiene mucha marcha.

### Primeros pasos:

Para comenzar a desarrollar el proyecto se tienen que tomar unos pasos iniciales. 

- [Estructura del proyecto](./docs/estructura.md)
- Primer código: [main](https://github.com/venrra/ragaeton/blob/master/cmd/ragaeton/main.go) (todo programa en go tiene que tener un pakage main), [direccion](https://github.com/venrra/ragaeton/blob/master/pkg/direccion/direccion.go) (primera clase o modulo o crear) 

#### :notebook_with_decorative_cover: Primeras tareas

Issue/milestones y  iniciales. A continuación se encuentran las primeras tareas issue y milestones que he creado para seguir la evolución del proyecto de forma ordenada y documentada

- [ ] [HU 1: consultar si un dirección esta activa o inactiva (ping)](https://github.com/venrra/ragaeton/issues/4): Un usuario podrá obtener información de una dirección dada. pertenece a [este milestone](https://github.com/venrra/ragaeton/milestone/1)
- [ ] [HU 2: enviar una dirección](https://github.com/venrra/ragaeton/issues/5): Un usuario podrá enviar una dirección cualquiera, y sera recibida y procesada.pertenece a [este milestone](https://github.com/venrra/ragaeton/milestone/1)
- [x] [Crear fichero IV.yaml](https://github.com/venrra/ragaeton/issues/1): crear y añadir la entidad en el archivo .yamli [iv.yaml](https://github.com/venrra/ragaeton/blob/master/iv.yaml). pertenece a [este milestone](https://github.com/venrra/ragaeton/milestone/1)
- [x] [Crear inicial estructura del proyecto](https://github.com/venrra/ragaeton/issues/2). pertenece a [este milestone](https://github.com/venrra/ragaeton/milestone/1)
- [x] [Documentación parte 1](https://github.com/venrra/ragaeton/issues/3). pertenece a [este milestone](https://github.com/venrra/ragaeton/milestone/1)


#### Primera ejecución: 

```
go run cmd/ragaeton/main.go
```

## Herramientas

Aun no se ha definido todas las herramientas necesarias: 

- Lenguaje: `Go`


## :page_facing_up: Documentación interesante

- [Mini-Guia de Go](https://github.com/venrra/ragaeton/blob/master/docs/guia-go.md)(en desarrollo :hammer:)
- Documentación go: [golang](https://golang.org/doc/)

## Enlaces de interés

- [Explicando herramientas usadas](./docs/herramientas.md)
- [¿Por que usar software libre?](./docs/softwareLibre.md)
- [Configuración de `git`](./docs/inicialGit.md)

## Autor 

- [Manu Hidalgo Carmona](https://github.com/venrra)