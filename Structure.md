```text
lisen/
├── cmd/                      # Los comandos que escribes en la terminal
│   ├── root.go               # Comando base 'lisen'
│   ├── register.go           # 'lisen register <user> <pass>'
│   ├── login.go              # 'lisen login <user> <pass>'
│   ├── play.go               # 'lisen play <cancion>'
│   ├── search.go             # 'lisen search <busqueda>'
│   └── fav.go                # 'lisen fav add <id>' / 'lisen fav list'
│
├── pkg/                      # Lógica interna de Go
│   ├── db/                   # Manejo de la base de datos (SQLite)
│   │   ├── db.go             # Conexión a la DB local
│   │   ├── user.go           # Crear usuario y validar login
│   │   └── favs.go           # Guardar y listar favoritos
│   │
│   ├── deezer/               # Peticiones HTTP a Deezer
│   │   ├── client.go         # Buscar canciones y obtener URLs de audio
│   │   └── models.go         # Estructuras JSON
│   │
│   └── audio/                # Reproductor de audio por línea de comandos
│       └── player.go         # Descarga el MP3 y suena hasta presionar Ctrl+C
│
├── go.mod                    # Modulo de Go
└── main.go                   # Llama a cmd.Execute()

```