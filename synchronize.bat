@echo off
setlocal enabledelayedexpansion

:: Get the script directory
for %%I in ("%~dp0") do set SCRIPT_DIR=%%~fI

:: Change to the script directory
cd /d "%SCRIPT_DIR%" || exit /b 1

echo Synchronizing JanOS
echo Ignite Laboratories
echo.

:: Function to synchronize a repository
:sync
if exist "%SCRIPT_DIR%\%1\.git" (
    echo [%1]
    cd /d "%SCRIPT_DIR%\%1" || exit /b 1
    git pull
    cd /d "%SCRIPT_DIR%"
) else (
    echo [%1]
    git clone "https://github.com/ignite-laboratories/%1" "%SCRIPT_DIR%\%1"
)
exit /b 0

:: Synchronize JanOS
if exist ".git" (
    echo [janos]
    git pull
) else (
    echo [janos]
    git clone "https://github.com/ignite-laboratories/janos" "%SCRIPT_DIR%\janos"
)

:: Call the sync function for multiple repositories
call :sync arwen
call :sync core
call :sync fugue
call :sync glitter
call :sync host
call :sync hydra
call :sync life
call :sync spark
call :sync support
call :sync tiny

:: End script
endlocal
exit /b 0