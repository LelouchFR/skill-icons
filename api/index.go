package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

var icons map[string]string = make(map[string]string)
var iconNameList []string
var themedIcons []string

/*
Say hi to the bob line.

"Why he exists without any icons?""
- Well, the icons are added when actions are run, we make this so the code can keep being clean.

Because of IDEs not coping with the bob line, I decided to put the line down below.
you can run `vercel dev` after running `go run build.go`, happy coding.

-hawl1
*/

var iconsJSON string

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
	"million":           "millionjs",
	"asm":               "assembly",
	"pop":               "popos",
	"nix":               "nixos",
	"hc":                "holyc",
	"yml":               "yaml",
	"twitter":           "x",
	"arc":               "arcbrowser"
}

const (
	ICONS_PER_LINE = 15
	ONE_ICON       = 48
	SCALE          = float64(ONE_ICON) / float64(300-44)
)

func generateSvg(iconNames []string, perLine int, hasTitlesEnabled bool) string {
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

func iconsHandler(ctx *fiber.Ctx) error {
	iconParam := ctx.Query("i")

	theme := ctx.Query("theme")
	if theme == "" {
		theme = "auto"
	}

	perLineStr := ctx.Query("perline")
	if perLineStr == "" {
		perLineStr = "15"
	}

	hasTitles := ctx.Query("titles")
	var hasTitlesEnabled bool
	if hasTitles == "" {
		hasTitlesEnabled = false
	} else {
		hasTitlesEnabled = true
	}

	if iconParam == "" {
		ctx.SendString("You didn't specify any icons!")
		return ctx.SendStatus(400)
	}

	if theme != "dark" && theme != "light" && theme != "auto" && theme != "" {
		ctx.SendString("Theme must be either 'light', 'dark' or 'auto'")
		return ctx.SendStatus(400)
	}

	perLine, err := strconv.Atoi(perLineStr)
	if err != nil || perLine < -1 || perLine > 50 {
		ctx.SendString("Icons per line must be a number between 1 and 50")
		return ctx.SendStatus(400)
	}

	var iconShortNames []string
	if iconParam == "all" {
		iconShortNames = iconNameList
	} else {
		iconShortNames = strings.Split(iconParam, ",")
	}

	iconNames := parseShortNames(iconShortNames, theme)
	if iconNames == nil {
		ctx.SendString("You didn't format the icons param correctly!")
		return ctx.SendStatus(400)
	}

	svg := generateSvg(iconNames, perLine, hasTitlesEnabled)

	ctx.Append("Content-Type", "image/svg+xml")
	ctx.SendString(svg)
	return ctx.SendStatus(200)
}

func handler() http.HandlerFunc {
	app := fiber.New()

	app.Get("/icons", iconsHandler)
	app.Get("/api/icons", iconsHandler)

	return adaptor.FiberApp(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(strings.NewReader(iconsJSON))
	if err := decoder.Decode(&icons); err != nil {
		panic(err)
	}
	r.RequestURI = r.URL.String()

	// Populate iconNameList and themedIcons
	for key := range icons {
		iconNameList = append(iconNameList, strings.Split(key, "-")[0])
		if strings.Contains(key, "-light") || strings.Contains(key, "-dark") || strings.Contains(key, "-auto") {
			themedIcons = append(themedIcons, strings.Split(key, "-")[0])
		}
	}

	handler().ServeHTTP(w, r)
}
