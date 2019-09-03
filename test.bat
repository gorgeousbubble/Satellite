:: Satellite Project - Build(Golang)
:: Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.
@echo off

:: Start Test All
echo Start Test Satellite:
goto test_pack

:: Test pack package
:test_pack
echo test pack:
echo.
echo.

cd pack
go test -v -cover -bench .

:: Test unpack package
:test_unpack
echo test unpack:
echo.
echo.

cd ..
cd unpack
go test -v -cover -bench .

:: Test comp package
:test_comp
echo test comp:
echo.
echo.

cd ..
cd comp
go test -v -cover -bench .
