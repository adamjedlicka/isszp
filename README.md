# isszp

Instalace:
````
go get gitlab.fit.cvut.cz/isszp/isszp
cd ~/go/src/gitlab.fit.cvut.cz/isszp/isszp
go run main.go install
````

Spuštění:
````
cd ~/go/src/gitlab.fit.cvut.cz/isszp/isszp
go run main.go
````

Pro zprovoznění UTF-8 znaků stačí upravit soubor /etc/my.cnf, konkrétně do něj vložit:
````
[client]
default-character-set = utf8mb4

[mysql]
default-character-set = utf8mb4

[mysqld]
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
````