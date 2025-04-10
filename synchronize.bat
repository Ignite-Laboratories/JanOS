@echo off
setlocal enabledelayedexpansion

echo "Synchronizing JanOS"
echo "Ignite Laboratories"

echo "- JanOS"
git pull

:: Function-like logic to synchronize repositories
:sync
if exist "%~1\.git" (
    pushd "%~1"
    echo "- %~1"
    git pull
    popd
) else (
    git clone "https://github.com/ignite-laboratories/%~1"
)
goto :eof

:: Synchronize multiple repositories
call :sync core
call :sync fugue
call :sync glitter
call :sync host
call :sync life
call :sync spark
call :sync support
call :sync tiny

:end
echo Synchronization completed!
exit /b