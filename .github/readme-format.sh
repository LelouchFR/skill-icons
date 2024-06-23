head -96 README.md > README.tmp && mv README.tmp README.md

icons=$(ls assets --ignore="*-light.svg" --ignore="*-dark.svg")
icon_list=($(ls assets --ignore="*-light.svg" --ignore="*-dark.svg"))

function get_maximum_length() {
    items=$1
    maximum=0
    for item in $items
    do
        if [[ ${#item} -gt $maximum ]]; then
            maximum=${#item}
        fi
    done
    echo $maximum
}

icons_counter=0
declare -a img_tags
for icon in $icons
do
    icons_counter=$((icons_counter + 1))
    img_tags+=("<img src=\"./assets/$icon\" width=\"48\">")
done

table_headers=("" "")
columns=$(($icons_counter / 100 + 1))
for column in $(seq 1 $columns)
do
    table_headers[0]+="| Icon ID | Icon "
    table_headers[1]+="| :-----------------: | :--------------: "
done
table_headers[0]+="|"
table_headers[1]+="|"

echo ${table_headers[0]} >> README.md
echo ${table_headers[1]} >> README.md

max_icon_id_length=$(get_maximum_length "$icons")
max_img_tag_length=$(get_maximum_length "${img_tags[*]}")

count_id=0
declare -a icon_table
for column in $(seq 1 $columns)
do
    for row in $(seq 1 100)
    do
        count_id=$(($count_id + 1))
        if [[ $count_id -ge $icons_counter ]]; then
            continue
        fi
        icon_id=$(echo ${icon_list[$(($count_id))]} | sed 's/\.[^.]*$//; s/-auto//g')
        img_tag="<img src=\"./assets/${icon_list[$(($count_id))]}\" width=\"48\">"

        padding_icon_id=$(( (max_icon_id_length - ${#icon_id}) / 2 ))
        padding_img_tag=$(( (max_img_tag_length - ${#img_tag}) / 2 ))

        padded_icon_id=$(printf "%*s%s%*s" $padding_icon_id "" "\`$icon_id\`" $padding_icon_id "")
        padded_img_tag=$(printf "%*s%s%*s" $padding_img_tag "" "$img_tag" $padding_img_tag "")


        if [[ $((max_icon_id_length % 2)) -ne 0 ]]; then
            padded_icon_id="$padded_icon_id "
        fi

        if [[ $((max_img_tag_length % 2)) -ne 0 ]]; then
            padded_img_tag="$padded_img_tag "
        fi

        icon_table[$row]+="|$padded_icon_id|$padded_img_tag"
    done
done

for column in $(seq 1 100)
do
    echo "${icon_table[$column]}|" >> README.md
done

echo "" >> README.md
echo "## 💖 Support the Project" >> README.md
echo "Thank you so much already for using my projects!" >> README.md
echo "To support the project directly, feel free to open issues for icon suggestions, or contribute with a pull request!" >> README.md
