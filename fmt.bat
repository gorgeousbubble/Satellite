:: Satellite Project - Build(Golang)
:: Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.
@echo off

:: Start Format All
echo Start Format Satellite:
goto fmt_pack

:: Format pack package
:fmt_pack
echo fmt pack:
echo.

cd pack
go fmt
echo.
echo.

:: Format unpack package
:fmt_unpack
echo fmt unpack:
echo.

cd ..
cd unpack
go fmt
echo.
echo.

:: Format comp package
:fmt_comp
echo fmt comp:
echo.

cd ..
cd comp
go fmt
echo.
echo.

:: Format decomp package
:fmt_decomp
echo fmt decomp:
echo.

cd ..
cd decomp
go fmt
echo.
echo.

:: Format sorts package
:fmt_sorts
echo fmt sorts:
echo.

cd ..
cd sorts
go fmt
echo.
echo.

:: Format searches package
:fmt_searches
echo fmt searches:
echo.

cd ..
cd searches
go fmt
echo.
echo.

:: Format utils package
:fmt_utils
echo fmt utils:
echo.

cd ..
cd utils
go fmt
echo.
echo.

:: Format cmd package
:fmt_cmd
echo fmt cmd:
echo.

cd ..
cd cmd
go fmt
echo.
echo.

:: Format ds package
:fmt_ds
echo fmt ds:
echo.

cd ..
cd ds
go fmt
echo.
echo.

:: Format exec package
:fmt_exec
echo fmt exec:
echo.

cd ..
cd exec
go fmt
echo.
echo.

:: Format global package
:fmt_global
echo fmt global:
echo.

cd ..
cd global
go fmt
echo.
echo.

:: Format logging package
:fmt_logging
echo fmt logging:
echo.

cd ..
cd logging
go fmt
echo.
echo.

:: Format logs package
:fmt_logs
echo fmt logs:
echo.

cd ..
cd logs
go fmt
echo.
echo.

:: Format nets package
:fmt_nets
echo fmt nets:
echo.

cd ..
cd nets
go fmt
echo.
echo.

:: Pasue Test
cd ..
pause