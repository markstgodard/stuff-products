#!/bin/bash
source $(pwd)/scripts/cf.cfg

cf push ${APP_NAME} -o markstgodard/stuff-products:v1 --no-start --no-manifest --no-route
cf set-env ${APP_NAME} REVIEWS_PROXY_ADDR "http://localhost:6379/reviews"
cf set-env ${APP_NAME} A8_SERVICE "products:v1"
cf set-env ${APP_NAME} A8_ENDPOINT_PORT "8080"
cf set-env ${APP_NAME} A8_ENDPOINT_TYPE "http"
cf set-env ${APP_NAME} A8_PROXY "true"
cf set-env ${APP_NAME} A8_REGISTER "true"
cf set-env ${APP_NAME} A8_REGISTRY_URL "http://${REGISTRY_NAME}.${ROUTES_DOMAIN}"
cf set-env ${APP_NAME} A8_CONTROLLER_URL "http://${CONTROLLER_NAME}.${ROUTES_DOMAIN}"

cf start ${APP_NAME}

cf access-allow ${STORE_NAME} ${APP_NAME} --port 8080 --protocol tcp
