@echo off

start "" "C:/Program Files (x86)/Mozilla Firefox/firefox.exe" "http://localhost:9000"
start C:\Users\Crystal\go\src\github.com\nekohor\gomon\app\app.exe %cd% 1

rem for /f "tokens=3,4" %%a in ('"reg query HKEY_CLASSES_ROOT\http\shell\open\command"') do (set SoftWareRoot=%%a %%b)
rem for /f "tokens=3,4" %%a in ('"reg query HKEY_CLASSES_ROOT\http\shell\open\command"') do (set SoftWareRoot=%%a %%b)
rem start "" % SoftWareRoot % http://localhost:8080

rem cd "C:\Program Files (x86)\Mozilla Firefox"
rem start firefox.exe http://localhost:8080

rem echo import os;os.system('"C:/Program Files (x86)/Mozilla Firefox/firefox.exe" http://localhost:9000') > openFireFox.py
rem echo import os;os.system("C:/Users/Crystal/go/src/github.com/nekohor/gomonitor/app/app.exe {}".format(os.getcwd())) >> openFireFox.py
rem python openFireFox.py