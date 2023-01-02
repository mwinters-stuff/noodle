#!/bin/bash
go install github.com/vektra/mockery/v2@latest


mockery --all --with-expecter --output internal/mocks/app

# #influxdb
# # ~/go/pkg/mod/github.com/influxdata/influxdb-client-go/v2@v2.10.0
# docker run -v $(pwd):/src -v /home/mathew/src/solar-zero-scrape-golang/internal/mocks/influxdb2:/mnt -w /src vektra/mockery  --all --with-expecter  --output /mnt/

# #project
# docker run -v $(pwd):/src -w /src vektra/mockery  --all --with-expecter  --output internal/mocks/app