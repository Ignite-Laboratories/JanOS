@echo off
echo Synchronizing JanOS
echo Ignite Laboratories

echo [janos]
if exist .git (
    git pull
) else (
    git clone "https://github.com/ignite-laboratories/JanOS"
)

:: Function-like structure for synchronizing repositories
:Synchronize
if exist "%~1\.git" (
    echo [%~1]
    pushd "%~1"
    git pull
    popd
) else (
    echo [%~1]
    git clone "https://github.com/ignite-laboratories/%~1"
)
goto :eof

:: Call the 'Synchronize' subroutine for multiple repositories
call :Synchronize core
call :Synchronize fugue
call :Synchronize glitter
call :Synchronize host
call :Synchronize life
call :Synchronize spark
call :Synchronize support
call :Synchronize tiny