:: Satellite Project - Format(Golang)
:: Copyright(C) 2019, Team Gorgeous Bubble, All Rights Reserved.
@echo off

:: Start Format All
echo Start Format Satellite:
goto fmt_all

:: Traversing folder format Golang code
:fmt_all
echo fmt all:
for /d %%i in (*) do echo %%i & cd %%i & echo. & if not %%i == .github ( if not %%i == .idea ( if not %%i == .vscode ( if not %%i == 3pp ( if not %%i == bin ( if not %%i == doc ( if not %%i == log ( if not %%i == plugin ( if not %%i == test ( go fmt ) ) ) ) ) ) ) ) ) & echo. & echo. & cd ..

:: Pause format
pause