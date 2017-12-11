#!/bin/bash

# upload cross-compiled link checker binaries
gsutil -m cp cross/* gs://$GOOGLE_BUCKET/*
# set uploaded binaries to readable by the public
gsutil -m acl set public-read gs://$GOOGLE_BUCKET/*
# force all objects to become immediately consistent
gsutil -m setmeta -h "Cache-Control:public, max-age=0" gs://$GOOGLE_BUCKET/*
