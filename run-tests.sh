#!/bin/bash

go test -covermode=count -coverprofile coverage.out -v ./handlers ./noodle/database ./noodle/heimdall ./noodle/ldap_handler 2>&1 | go-junit-report > report.xml
go tool cover -html coverage.out -o cover.html
gcov2lcov -infile=coverage.out -outfile=coverage.lcov