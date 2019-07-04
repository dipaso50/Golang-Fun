
# Implementación simple de grep con Go
----------------------------------------
Esta es una implementación libre y con fines didácticos de la utilidad grep de los sistemas *nix.

# Uso
----------------------------------------
Para mostrar ayuda:

```sh
$ ggrep -h
Usage of ggrep:
  -F string
        Folder to search
  -e string
        Regex expression for search
  -f string
        File to search
  -fe string
        Regex expression for the file name for search
``` 

Es posible redirigir la salida de un comando a la entrada de ggrep, de la siguiente forma:
```sh
$comando | ggrep -e=expr
```

# Ejemplos
```sh
$ ggrep -F=/opt -e=error // busca dentro de la carpeta /opt todos los ficheros que contengan la palabra 'error'
$ ggrep -f=/var/log/syslog -e=panic // busca la palabra panic en el fichero /var/log/syslog
$ ggrep -F=/var/log -fe=log -e=something //busca la palaba something en todos los ficheros que contengan en su nombre la palabra log dentro del directorio log
```



# Instalación
----------------------------------------
Para instalarlo en cualquier sistema solamente hace falta añadir la carpeta donde se encuentra el ejecutable a la variable Path.

Golang permite cross-compilation, por tanto es posible compilar este programa desde una plataforma Linux para usarlo en sistemas Windows, por ejemplo para generar un ejecutable para Windows 64bit, bastaría con ejecutar:
```sh
$env GOOS=windows GOARCH=amd64 go build -o ggrep.exe
```

# Licencia
----------------------------------------
MIT 
