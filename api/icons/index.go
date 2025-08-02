package icons

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

var icons map[string]string = make(map[string]string)
var iconNameList []string
var themedIcons []string





var shortNames = map[string]string{
	"js":                "javascript",
	"ts":                "typescript",
	"py":                "python",
	"tailwind":          "tailwindcss",
	"vue":               "vuejs",
	"nuxt":              "nuxtjs",
	"go":                "golang",
	"cf":                "cloudflare",
	"wasm":              "webassembly",
	"postgres":          "postgresql",
	"k8s":               "kubernetes",
	"next":              "nextjs",
	"mongo":             "mongodb",
	"md":                "markdown",
	"ps":                "photoshop",
	"ai":                "illustrator",
	"pr":                "premiere",
	"ae":                "aftereffects",
	"scss":              "sass",
	"sc":                "scala",
	"net":               "dotnet",
	"gatsbyjs":          "gatsby",
	"gql":               "graphql",
	"vlang":             "v",
	"amazonwebservices": "aws",
	"bots":              "discordbots",
	"express":           "expressjs",
	"googlecloud":       "gcp",
	"mui":               "materialui",
	"windi":             "windicss",
	"unreal":            "unrealengine",
	"nest":              "nestjs",
	"ktorio":            "ktor",
	"pwsh":              "powershell",
	"au":                "audition",
	"rollup":            "rollupjs",
	"rxjs":              "reactivex",
	"rxjava":            "reactivex",
	"ghactions":         "githubactions",
	"sklearn":           "scikitlearn",
	"ml5":               "ml5js",
	"vb":                "visualbasic",
	"an":                "animate",
	"ca":                "capture",
	"cc":                "creativecloud",
	"ch":                "characteranimator",
	"me":                "mediaencoder",
	"pl":                "prelude",
	"ru":                "premiererush",
	"fs":                "fuse",
	"id":                "indesign",
	"ic":                "incopy",
	"sp":                "adobespark",
	"dw":                "dreamweaver",
	"dn":                "dimension",
	"ar":                "aero",
	"psc":               "photoshopclassic",
	"psx":               "photoshopexpress",
	"lr":                "lightroom",
	"lrc":               "lightroomclassic",
	"fr":                "fresco",
	"pf":                "portfolio",
	"st":                "stock",
	"be":                "behance",
	"br":                "bridge",
	"million":           "millionjs",
	"asm":               "assembly",
	"pop":               "popos",
	"nix":               "nixos",
	"hc":                "holyc",
	"yml":               "yaml",
	"twitter":           "x",
	"arc":               "arcbrowser",
	"hf":                "huggingface",
	"sqla":              "sqlalchemy",
    "notepad++":         "notepadpp",
    "jq":                "jqlang",
}

const (
	ICONS_PER_LINE = 15
	ONE_ICON       = 48
	SCALE          = float64(ONE_ICON) / float64(300-44)
)

func generateSvg(iconNames []string, perLine int, hasTitlesEnabled bool, align string) string {
	iconSvgList := make([]string, len(iconNames))

	for i, name := range iconNames {
		iconSvgList[i] = icons[name]
	}

	length := int(math.Min(float64(perLine*300), float64(len(iconNames)*300))) - 44
	height := int(math.Ceil(float64(len(iconSvgList))/float64(perLine)))*300 - 44
	scaledHeight := int(float64(height) * SCALE)
    scaledWidth := int(float64(length) * SCALE)

    var svg string
    if align == "center" {
        svg = fmt.Sprintf(`
            <svg width="100%%" height="%d" viewBox="0 0 %d %d" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1">
        `, scaledHeight, length, height)
    } else if align == "right" {
        svg = fmt.Sprintf(`
            <svg width="calc(200%% - %dpx)" height="%d" viewBox="0 0 %d %d" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1">
        `, scaledWidth, scaledHeight, length, height)
    } else {
        svg = fmt.Sprintf(`
            <svg width="%d" height="%d" viewBox="0 0 %d %d" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1">
        `, scaledWidth, scaledHeight, length, height)
    }

	for index, i := range iconSvgList {
        var title string
        if hasTitlesEnabled {
            title = fmt.Sprintf("<title>%s</title>", iconNames[index])
        }

		x := (index % perLine) * 300
		y := (index / perLine) * 300
		svg += fmt.Sprintf(`
		<g transform="translate(%d, %d)">
            %s
			%s
		</g>
		`, x, y, title, i)
	}

	svg += "</svg>"
	return svg
}

func parseShortNames(names []string, theme string) []string {
	result := make([]string, len(names))

	for i, name := range names {
		if contains(iconNameList, name) {
			if contains(themedIcons, name) {
				result[i] = name + "-" + theme
			} else {
				result[i] = name
			}
		} else if val, ok := shortNames[name]; ok {
			if contains(themedIcons, val) {
				result[i] = val + "-" + theme
			} else {
				result[i] = val
			}
		}
	}
	return result
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func init(){
	decoder := json.NewDecoder(strings.NewReader(iconsJSON))
	if err := decoder.Decode(&icons); err != nil {
		panic(err)
	}

	// Populate iconNameList and themedIcons
	for key := range icons {
		iconNameList = append(iconNameList, strings.Split(key, "-")[0])
		if strings.Contains(key, "-light") || strings.Contains(key, "-dark") || strings.Contains(key, "-auto") {
			themedIcons = append(themedIcons, strings.Split(key, "-")[0])
		}
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	parse_form_err := r.ParseForm()
	if parse_form_err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
	}

	iconParam := r.Form.Get("i")

	theme := r.Form.Get("theme")
	if theme == "" {
		theme = "auto"
	}

	perLineStr := r.Form.Get("perline")
	if perLineStr == "" {
		perLineStr = "15"
	}

	hasTitles := r.Form.Get("titles")
	var hasTitlesEnabled bool
	if hasTitles == "" {
		hasTitlesEnabled = false
	} else {
		hasTitlesEnabled = true
	}

	align := r.Form.Get("align")
	if align != "left" && align != "right" && align != "center" {
		align = "left"
	}

	if iconParam == "" {
		http.Error(w, "You didn't specificy any icons!", http.StatusBadRequest)
		return
	}

	if theme != "dark" && theme != "light" && theme != "auto" && theme != "" {
		http.Error(w, "Theme must be either 'light', 'dark' or 'auto'", http.StatusBadRequest)
		return
	}

	perLine, err := strconv.Atoi(perLineStr)
	if err != nil || perLine < -1 || perLine > 50 {
		http.Error(w, "Icons per line must be a number between 1 and 50", http.StatusBadRequest)
		return
	}

	var iconShortNames []string
	if iconParam == "all" {
		iconShortNames = iconNameList
	} else {
		iconShortNames = strings.Split(iconParam, ",")
	}

	iconNames := parseShortNames(iconShortNames, theme)
	if iconNames == nil {
		http.Error(w, "You didn't format the icons param correctly!", http.StatusBadRequest)
		return
	}

	svg := generateSvg(iconNames, perLine, hasTitlesEnabled, align)
	
	// Cache header set to 7 days (7 days * 24 hours * 60 minutes * 60 seconds = 604800 seconds)
	cacheHeader := "public, s-maxage=604800"
	w.Header().Set("Cache-Control", cacheHeader)
	w.Header().Set("CDN-Cache-Control", cacheHeader)
	w.Header().Set("Vercel-CDN-Cache-Control", cacheHeader)

	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, svg)
}
