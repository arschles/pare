# upload cross-compiled link checker binaries
gsutil -m cp pare gs://$GOOGLE_BUCKET/pare
# set uploaded binaries to readable by the public
gsutil -m acl set public-read gs://$GOOGLE_BUCKET/pare
# force all objects to become immediately consistent
gsutil -m setmeta -h "Cache-Control:public, max-age=0" gs://$GOOGLE_BUCKET/pare
