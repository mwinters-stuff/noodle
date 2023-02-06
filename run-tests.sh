#!/bin/bash

go test -covermode=count -coverprofile coverage.out -v ./noodle/options ./noodle/configure_server ./noodle/api_handlers ./noodle/database ./noodle/heimdall ./noodle/ldap_handler 2>&1 | tee /tmp/test.out | go-junit-report > report.xml
cat /tmp/test.out
go tool cover -html coverage.out -o cover.html
gcov2lcov -infile=coverage.out -outfile=coverage.lcov

# pushd noodlewebclient
# flutter test --coverage --reporter json | tee test.json
