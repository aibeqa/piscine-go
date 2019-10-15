$ curl -0 https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json
echo "$curl"
jq ".[].id"
echo "jq"

