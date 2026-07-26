# Lisen Words

CLI para buscar y reproducir canciones desde Deezer directamente en la terminal.

## Instalación

### Opción 1 — Descargar binario

1. Ve a [Releases](https://github.com/bulletformyvalentinefan/Lisen-Words/releases)
2. Descarga `lisen.exe`
3. Colócalo en una carpeta incluida en tu `PATH` (ej: `%USERPROFILE%\go\bin\`)
4. `lisen --help`

### Opción 2 — Clonar + compilar (requiere Go)

```cmd
git clone https://github.com/bulletformyvalentinefan/lisen-words.git
cd lisen-words
install.bat     → compila y lo deja en $GOPATH/bin
lisen --help
```

## Uso

```
lisen register              Crear una cuenta
lisen login                 Iniciar sesión
lisen search <canción>      Buscar canciones en Deezer
lisen play <canción>        Reproducir preview de una canción
lisen --help                Ayuda general
```

## Ejemplo

```cmd
lisen register
> Usuario: bullet
> Contraseña: ****

lisen login
> Usuario: bullet
> Contraseña: ****
> ✓ Sesión iniciada

lisen play "bullet for my valentine tears don't fall"
> Reproduciendo: Tears Don't Fall - Bullet For My Valentine
```
