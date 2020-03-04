:: Satellite Project - Build(Golang)
:: Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.
@echo off

:: Build Golang Mode(Debug/Release)
set "build_target=Release"

:: Get the module dir name
set dirName = 
goto getDirName

:: Get the absolute dir path
set absPath =
goto getAbsPath

:: Delete image process
taskkill /f /fi "IMAGENAME eq %dirName%.exe" > nul

for /f "skip=3 tokens=4" %%i in ('sc query AeLookupSvc') do set "zt=%%i" 
if /i "%zt%"=="RUNNING" ( echo . ) else ( net start "AeLookupSvc" )

for /f "skip=3 tokens=4" %%i in ('sc query PcaSvc') do set "zt=%%i" 
if /i "%zt%"=="RUNNING" ( echo . ) else ( net start "PcaSvc" )

:: Delete old compile file
if exist %dirName%.exe ( del %dirName%.exe ) > nul

:: Build Golang App
goto build_go_app

:: Get dir name function
:getDirName
    set "lj=%~p0"
    set "lj=%lj:\= %"
    for %%a in (%lj%) do set wjj=%%a
    set dirName=%wjj%

:getAbsPath
    set absPath=%~dp0

:build_go_app
    echo build go app:
    echo.

    :: Build Mode
    if "%build_target%"=="Debug" ( goto build_debug ) else ( goto build_release )

:build_debug
    :: Build Golang(With debug information)
    echo [Build Debug]
    echo go build -o bin/satellite.exe main.go
    go build -o bin/satellite.exe main.go
    echo.
    echo.

    :: Build Result
    if exist bin/%dirName%.exe (( choice /t 1 /d y /n >nul ) | ( echo build success. )) else ( echo build failure! | pause )
    echo.
    echo.

    :: Build Pause & Exit
    pause
    exit

:build_release
    :: Build Golang(Without debug information)
    echo [Build Release]
    echo go build -ldflags="-w -s" -o bin/satellite.exe main.go
    go build -ldflags="-w -s" -o bin/satellite.exe main.go
    echo.
    echo.

    :: UPX Shell Pack
    echo upx -9 bin/%dirName%.exe
    %absPath%tools\upx -9 bin/%dirName%.exe
    echo.
    echo.

    :: Build Result
    if exist bin/%dirName%.exe (( choice /t 1 /d y /n >nul ) | ( echo build success. )) else ( echo build failure! | pause )
    echo.
    echo.

    :: Build Pause & Exit
    pause
    exit