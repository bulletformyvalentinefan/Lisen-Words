@echo off
go build -o "%USERPROFILE%\go\bin\lisen.exe" .
if %ERRORLEVEL% neq 0 (
    echo Error en la compilacion.
    exit /b %ERRORLEVEL%
)
echo lisen instalado globalmente en %%USERPROFILE%%\go\bin\lisen.exe
