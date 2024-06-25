import os

# Function to get maximum length of items in a list
def get_maximum_length(items):
    return max(len(item) for item in items)

# List all icons in the assets directory
icons = [icon for icon in os.listdir('assets') if not (icon.endswith('-light.svg') or icon.endswith('-dark.svg'))]

# Generate img tags for each icon
img_tags = [f'<img src="./assets/{icon}" width="48">' for icon in icons]

# Prepare table headers
icons_counter = len(icons)
table_headers = ["| Icon ID | Icon "] * 2
columns = (icons_counter // 100) + 1

for i in range(len(table_headers)):
    table_headers[i] += "|"

# Write table headers to README.md
with open('README.md', 'w') as f:
    f.write("\n".join(table_headers) + "\n\n")

# Function to pad strings to center
def pad_to_center(text, width):
    total_padding = width - len(text)
    if total_padding > 0:
        left_padding = total_padding // 2
        right_padding = total_padding - left_padding
        return ' ' * left_padding + text + ' ' * right_padding
    else:
        return text

# Write icon table to README.md
max_icon_id_length = get_maximum_length(icons)
max_img_tag_length = get_maximum_length(img_tags)

count_id = 0
icon_table = [[] for _ in range(100)]

for column in range(columns):
    for row in range(100):
        count_id += 1
        if count_id > icons_counter:
            continue
        
        icon_id = os.path.splitext(icons[count_id - 1])[0].replace('-auto', '')
        img_tag = img_tags[count_id - 1]

        padded_icon_id = pad_to_center(f"`{icon_id}`", max_icon_id_length)
        padded_img_tag = pad_to_center(img_tag, max_img_tag_length)

        icon_table[row].append(f"|{padded_icon_id}|{padded_img_tag}")

# Write icon table rows to README.md
with open('README.md', 'a') as f:
    for row in icon_table:
        f.write("".join(row) + "|\n")

# Write additional content to README.md
with open('README.md', 'a') as f:
    f.write("\n")
    f.write("## 💖 Support the Project\n")
    f.write("Thank you so much already for using my projects!\n")
    f.write("To support the project directly, feel free to open issues for icon suggestions, or contribute with a pull request!\n")
