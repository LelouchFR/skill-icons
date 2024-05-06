package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
	"sp":                "spark",
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
}

var (
	app *gin.Engine
)

const (
	ICONS_PER_LINE = 15
	ONE_ICON       = 48
	SCALE          = float64(ONE_ICON) / float64(300-44)
)

func generateSvg(iconNames []string, perLine int) string {
	iconSvgList := make([]string, len(iconNames))

	for i, name := range iconNames {
		iconSvgList[i] = icons[name]
	}

	length := int(math.Min(float64(perLine*300), float64(len(iconNames)*300))) - 44
	height := int(math.Ceil(float64(len(iconSvgList))/float64(perLine)))*300 - 44
	scaledHeight := int(float64(height) * SCALE)
	scaledWidth := int(float64(length) * SCALE)

	svg := fmt.Sprintf(`
	<svg width="%d" height="%d" viewBox="0 0 %d %d" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1">
	`, scaledWidth, scaledHeight, length, height)

	for index, i := range iconSvgList {
		x := (index % perLine) * 300
		y := (index / perLine) * 300
		svg += fmt.Sprintf(`
		<g transform="translate(%d, %d)">
			%s
		</g>
		`, x, y, i)
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

func iconRoute(r *gin.RouterGroup) {
	r.GET("/icons", func(ctx *gin.Context) {

		ctx.Request.ParseForm()
		iconParam := ctx.Request.Form.Get("i")

		theme := ctx.Request.Form.Get("theme")
		if theme == "" {
			theme = "dark"
		}

		perLineStr := ctx.Request.Form.Get("perline")
		if perLineStr == "" {
			perLineStr = "15"
		}

		if iconParam == "" {
			ctx.String(http.StatusBadRequest, "You didn't specify any icons!")
			return
		}

		if theme != "dark" && theme != "light" && theme != "" {
			ctx.String(http.StatusBadRequest, "Theme must be either 'light' or 'dark'")
			return
		}

		perLine, err := strconv.Atoi(perLineStr)
		if err != nil || perLine < -1 || perLine > 50 {
			ctx.String(http.StatusBadRequest, "Icons per line must be a number between 1 and 50")
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
			ctx.String(http.StatusBadRequest, "You didn't format the icons param correctly!")
			return
		}

		svg := generateSvg(iconNames, perLine)

		ctx.Header("Content-Type", "image/svg+xml")
		ctx.String(http.StatusOK, svg)
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")

	iconRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(strings.NewReader(iconsJSON))
	if err := decoder.Decode(&icons); err != nil {
		panic(err)
	}

	// Populate iconNameList and themedIcons
	for key := range icons {
		iconNameList = append(iconNameList, strings.Split(key, "-")[0])
		if strings.Contains(key, "-light") || strings.Contains(key, "-dark") {
			themedIcons = append(themedIcons, strings.Split(key, "-")[0])
		}
	}

	app.ServeHTTP(w, r)
}
