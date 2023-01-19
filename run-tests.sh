#!/bin/bash

go test -covermode=count -coverprofile cover.out -v ./handlers ./noodle/database ./noodle/heimdall ./noodle/ldap_handler
go tool cover -html cover.out -o cover.html