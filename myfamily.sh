curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg id "$HERO_ID" '.[] | select(.id== ($id|tonumber )) | .connections | .relatives' | sed 's/"//g'

