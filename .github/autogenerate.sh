#!/bin/bash

echo "|      Icon ID        |                           Icon                           |" > generated_table.md
echo "| :-----------------: | :------------------------------------------------------: |" >> generated_table.md

for icon_path in assets/*.svg; do
    icon_filename=$(basename "$icon_path")
    icon_id=$(echo "$icon_filename" | sed -E 's/(-dark|-light|-auto)?\.svg$//')

    if [[ "$icon_filename" =~ -light\.svg$ || "$icon_filename" =~ -dark\.svg$ ]]; then
        continue
    fi

    if [ -f "assets/${icon_id}-auto.svg" ]; then
        printf "| \`%s\` | <img src=\"./assets/%s\" width=\"48\"> |\n" \
               "$icon_id" "$icon_filename" >> generated_table.md
    elif [ -f "assets/${icon_id}.svg" ]; then
        printf "| \`%s\` | <img src=\"./assets/%s\" width=\"48\"> |\n" "$icon_id" "$icon_filename" >> generated_table.md
    fi
done
