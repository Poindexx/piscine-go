curl -s "https://zero.academie.one/assets/superhero/all.json" | jq '.[] | select(.id == '$HERO_ID') | .connections.relatives' | tr -d '"'

