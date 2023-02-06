#!/bin/bash
npx @openapitools/openapi-generator-cli generate -i ../swagger/noodle_service.yaml -g  typescript-fetch -o src/api/
# npx swagger-typescript-api -p ../swagger/noodle_service.yaml -o ./src/ -n noodleApi.ts
