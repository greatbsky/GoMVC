@echo off
cd %~dp0
set GOPATH=%~dp0
set PATH=%PATH%;%GOPATH%\bin
start cmd /K
