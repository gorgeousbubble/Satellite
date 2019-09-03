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

cd pack
go test -v -cover -bench .
echo.
echo.

:: Test unpack package
:test_unpack
echo test unpack:
echo.

cd ..
cd unpack
go test -v -cover -bench .
echo.
echo.

:: Test comp package
:test_comp
echo test comp:
echo.

cd ..
cd comp
go test -v -cover -bench .
echo.
echo.

:: Test decomp package
:test_decomp
echo test decomp:
echo.

cd ..
cd decomp
go test -v -cover -bench .
echo.
echo.

:: Test sorts package
:test_sorts
echo test sorts:
echo.

cd ..
cd sorts
go test -v -cover -bench .
echo.
echo.

:: Test searches package
:test_searches
echo test searches:
echo.

cd ..
cd searches
go test -v -cover -bench .
echo.
echo.

:: Test utils package
:test_utils
echo test utils:
echo.

cd ..
cd utils
go test -v -cover -bench .
echo.
echo.

:: Pasue Test
pause