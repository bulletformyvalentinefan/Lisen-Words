# Lisen Words

CLI para buscar y reproducir canciones desde Deezer directamente en la terminal.

## Instalación

### Opcion 1 — Descargar binario

1. Ve a [Releases](https://github.com/bulletformyvalentinefan/Lisen-Words/releases)
2. Descarga `lisen.exe`
3. Colocalo en una carpeta incluida en tu `PATH` (ej: `%USERPROFILE%\go\bin\`)
4. `lisen --help`

### Opcion 2 — Clonar + compilar (requiere Go)

```cmd
git clone https://github.com/bulletformyvalentinefan/lisen-words.git
cd lisen-words
install.bat
lisen --help
```

## Uso

```
lisen register -u <usuario> -p <password>       Crear cuenta
lisen login -u <usuario> -p <password>          Iniciar sesion
lisen search <cancion>                          Buscar canciones en Deezer
lisen play <cancion>                            Reproducir preview de una cancion
lisen --help                                    Ayuda general
```

## Ejemplo

```cmd
lisen register  bullet  123
> Usuario 'bullet' registrado exitosamente!

lisen login  bullet  123
> Bienvenido bullet! Inicio de sesion exitoso (ID: 1).

lisen play "bullet for my valentine tears don't fall"
> Reproduciendo: Tears Don't Fall - Bullet For My Valentine
```
