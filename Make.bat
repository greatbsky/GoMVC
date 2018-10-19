@echo off
cd %~dp0
if %1 == "init" goto init
if %1 == "install" goto install

:init
set GOPATH=%~dp0
set PATH=%PATH%;%GOPATH%\bin
GOTO:EOF

:install
echo -------------------------------------- begin to make install...
::go get github.com/chainlibs/gobtclib
::go get -u go.uber.org/zap
::go get -u github.com/gobasis/log
::mkdir %GOPATH%\src\github.com\chainlibs
::rmdir /s/q %GOPATH%\src\github.com\chainlibs\gobtclib
::mklink /D %GOPATH%\src\github.com\chainlibs\gobtclib %cd%
echo -------------------------------------- finished successfully!
pause
GOTO:EOF