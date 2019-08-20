:: Satellite Project - Build(Golang)
:: Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.
@echo off

:: Get the module dir name
set dirName = 
goto getDirName

:: Delete image process
taskkill /f /fi "IMAGENAME eq %dirName%.exe" > nul

for /f "skip=3 tokens=4" %%i in ('sc query AeLookupSvc') do set "zt=%%i" 
if /i "%zt%"=="RUNNING" (  echo . ) else (  net start "AeLookupSvc" )

for /f "skip=3 tokens=4" %%i in ('sc query PcaSvc') do set "zt=%%i" 
if /i "%zt%"=="RUNNING" (  echo . ) else (  net start "PcaSvc" )

:: Delete old compile file
if     exist    %dirName%.exe      (  del %dirName%.exe ) > nul

:: Build Golang App
goto build_go_app

:: Get dir name fuction
:getDirName
    set "lj=%~p0"
    set "lj=%lj:\= %"
    for %%a in (%lj%) do set wjj=%%a
    set dirName=%wjj%

:build_go_app
    echo.
    echo build go app:

    :: Build Golang
    go build -o bin/satellite.exe main.go
    echo go build -o bin/satellite.exe main.go
    echo.
    echo.
    echo.

    if     exist    %dirName%.exe       (    ( choice /t 1 /d y /n >nul )   | ( echo  ok, build success. )  & ( start %dirName%.exe )  )      else    (  echo  build failure! | pause  )
    echo.
    echo.