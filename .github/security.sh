if [[ $1 -eq 0 ]]; then
    has_img=$(grep -l "data:image/png;base64" ./assets/*.svg)

    if [ -z "$has_img" ]; then
        echo "no image found in svg"
    else
        echo "image found in svg, please reformat you svg to have no image"
        echo "$has_img"
        exit 1
    fi
else
    has_js_tag=$(grep -l "<script" ./assets/*.svg)

    if [[ -z "$has_js_tag" ]]; then
        echo "no xss attack found"
    else
        echo "xss attack detected in following file:"
        echo "$has_js_tag"
        exit 1
    fi
fi
