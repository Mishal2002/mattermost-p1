#!/usr/bin/env bash
if [ -z "$FROM" ]
then
  echo "Missing FROM version. Usage: make diff-config FROM=v1.2.3 TO=v1.2.3"
  exit 1
fi

if [ -z "$TO" ]
then
  echo "Missing TO version. Usage: make diff-config FROM=v1.2.3 TO=v1.2.3"
  exit 1
fi

# Returns the config file for a specific release
fetch_config() {
  local url="https://releases.mattermost.com/$1/mattermost-$1-linux-amd64.tar.gz"
  curl -sf "$url" | tar -xzOf - mattermost/config/config.json | jq -S .
}

echo Fetching config files
from_config="$(fetch_config "$FROM")"
if [ -z "$from_config" ]
then
  echo Invalid version "$FROM"
  exit 1
fi

to_config=$(fetch_config "$TO")
if [ -z "$to_config" ]
then
  echo Invalid version "$TO"
  exit 1
fi

echo Comparing config files
diff -y <(echo "$from_config") <(echo "$to_config")

# We ignore exits with 1 since it just means there's a difference, which is fine for us.
diff_exit=$?
if [ $diff_exit -eq 1 ]; then
  exit 0
else
  exit $diff_exit
fi
