@echo off
setlocal enabledelayedexpansion

:: Function-like logic to synchronize repositories
:sync
if exist "%~1\.git" (
    echo Opening directory %~1 and pulling the latest changes...
    pushd "%~1"
    git pull
    popd
) else (
    echo Cloning repository https://github.com/ignite-laboratories/%~1 ...
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