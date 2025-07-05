@echo off
setlocal enabledelayedexpansion

:: Get directory of the script
for %%I in ("%~f0") do set SCRIPT_DIR=%%~dpI
cd /d "%SCRIPT_DIR%"

echo Synchronizing JanOS
echo Ignite Laboratories
echo.

:: Synchronize the "janos" repository
echo [janos]
if exist .git (
    git pull
) else (
    git clone "https://github.com/ignite-laboratories/janos"
)

:: Define the synchronize function
call :synchronize core
call :synchronize glitter
call :synchronize host
call :synchronize hydra
call :synchronize support
call :synchronize tiny

goto :eof

:synchronize
echo.
if exist "%~1\.git" (
    pushd "%~1"
    echo [%~1]
    git pull
    popd
) else (
    echo [%~1]
    git clone "https://github.com/ignite-laboratories/%~1"
)
goto :eof