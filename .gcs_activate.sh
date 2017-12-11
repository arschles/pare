#!/bin/ash

echo $GOOGLE_JSON >> secret.json
CLOUDSDK_PYTHON_SITEPACKAGES=1 gcloud auth activate-service-account $GOOGLE_ACCOUNT \
    --key-file secret.json
