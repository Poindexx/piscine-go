curl -s "https://zero.academie.one/assets/superhero/all.json" | jq '.[] | select(.id == 170) | .name, .powerstats, .appearance.gender'

