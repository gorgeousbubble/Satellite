:: Satellite Project - Clean(Golang)
:: Copyright(C) 2019~2020, Team Gorgeous Bubble, All Rights Reserved.
@echo off

:: Get the module dir name
set dirName =
goto getDirName

:: Get the absolute dir path
set absPath =
goto getAbsPath

:: Get dir name function
:getDirName
    set "lj=%~p0"
    set "lj=%lj:\= %"
    for %%a in (%lj%) do set wjj=%%a
    set dirName=%wjj%

:: Get absolute dir path
:getAbsPath
    set absPath=%~dp0

:: Clean Golang App
:clean_go_app:
    echo clean go app:
    echo.
    echo.

    :: clean build folder
    echo clean build folder:
    rd /s /q %absPath%build
    echo.
    echo.

    :: clean log folder
    echo clean log folder:
    rd /s /q %absPath%log
    echo.
    echo.

    :: clean finished
    echo clean success.
    echo.
    echo.
