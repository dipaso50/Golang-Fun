
#Implementación simple de grep con Go
----------------------------------------
Esta es una implementación libre y con fines puramente didácticos.



##Uso
----------------------------------------
Para mostrar ayuda:

```sh
$ ggrep -h
Usage of ggrep:
  -F string
    	Folder search.
  -e string
    	Expresion to search.
  -f string
    	File search.
``` 

Es posible redirigir la salida de un comando a la entrada de ggrep, de la siguiente forma:
```sh
$comando | ggrep -e=expr
```

Es posible buscar dentro de un directorio o directamente en un fichero, por ejemplo:
buscar en directorio y todos sus subdirectorios: 'ggrep -F=/opt -e=loquesea'
buscar en fichero: 'ggrep -f=/opt/aFile.txt -e=loquesea'


#Instalación
----------------------------------------
Para instalarlo en cualquier sistema solamente hace falta añadir la carpeta donde se encuentra el ejecutable a la variable Path.

Golang permite cross-compilation, por tanto es posible compilar este programa desde una plataforma linux para usarlo en windows, por ejemplo si se quiere generar un ejecutable para Windows 64bit, se usaría el siguiente comando:
```sh
$env GOOS=windows GOARCH=amd64 go build -o ggrep.exe
```

#Licencia
----------------------------------------i
Copyright 2019 Diego Sotomayor

Se concede permiso por la presente, libre de cargos, a cualquier persona que obtenga una copia de este software y de los archivos de documentación asociados (el "Software"), a utilizar el Software sin restricción,
 incluyendo sin limitación los derechos a usar, copiar, modificar, fusionar, publicar, distribuir, sublicenciar, y/o vender copias del Software, y a permitir a las personas a las que se les proporcione el Software a hacer lo mismo, sujeto a las siguientes condiciones:

El aviso de copyright anterior y este aviso de permiso se incluirán en todas las copias o partes sustanciales del Software.

EL SOFTWARE SE PROPORCIONA "COMO ESTÁ", SIN GARANTÍA DE NINGÚN TIPO, EXPRESA O IMPLÍCITA, INCLUYENDO PERO NO LIMITADO A GARANTÍAS DE COMERCIALIZACIÓN, IDONEIDAD PARA UN PROPÓSITO PARTICULAR E INCUMPLIMIENTO.
 EN NINGÚN CASO LOS AUTORES O PROPIETARIOS DE LOS DERECHOS DE AUTOR SERÁN RESPONSABLES DE NINGUNA RECLAMACIÓN, DAÑOS U OTRAS RESPONSABILIDADES, YA SEA EN UNA ACCIÓN DE CONTRATO, AGRAVIO O CUALQUIER OTRO MOTIVO, DERIVADAS DE, FUERA DE O EN CONEXIÓN CON EL SOFTWARE O SU USO U OTRO TIPO DE ACCIONES EN EL SOFTWARE. 
