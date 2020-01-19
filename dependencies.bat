:: Satellite Project - Get dependencies(Golang)
:: Copyright(C) 2019~2020, Team Gorgeous Bubble, All Rights Reserved.
@echo off

:: Golang get dependencies automatically
:get_dependencies
    echo Golang get dependencies packages
    echo go get -v -t -d ./...
    go get -v -t -d ./...
    echo.
    echo.