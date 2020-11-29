#!/bin/bash
# refs:
# https://developer.twitter.com/en/docs/authentication/oauth-2-0/bearer-tokens
# jq -r : https://www.qoosky.io/techs/1ee07c140f

set -u

# https://developer.twitter.com/en/apps/17280063
API_KEY=$1
API_SECRET_KEY=$2

BEARER_TOKEN=`curl -u "$API_KEY:$API_SECRET_KEY" \
  --data 'grant_type=client_credentials' \
  'https://api.twitter.com/oauth2/token' | jq -r .access_token`

curl https://api.twitter.com/1.1/statuses/user_timeline.json?user_id=1012936956353761280 \
    -H "Authorization: Bearer ${BEARER_TOKEN}"