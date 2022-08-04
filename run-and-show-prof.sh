#!/bin/sh
go run -tags ebitensinglethread main.go
go tool pprof -png profile.prof > profile.png
xdg-open profile.png
