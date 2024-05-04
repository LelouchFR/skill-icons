package api

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var skillAliases = map[string]string{
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

func resolveSkillAlias(alias string) string {
	if skill, ok := skillAliases[alias]; ok {
		return skill
	}
	return alias
}

func GetSkillIcons(ctx *gin.Context) {
	skills := ctx.Query("i")
	theme := ctx.Query("theme")
	perLineQuery := ctx.Query("perline")
	hasTitleQuery := ctx.Query("titles")

	perLine, err := strconv.Atoi(perLineQuery)
	if err != nil {
		perLine = IconsPerLine
	}

	hasTitle, err := strconv.ParseBool(hasTitleQuery)
	if err != nil {
		hasTitle = false
	}

	if theme == "" {
		theme = "dark"
	} else {
		theme = "light"
	}

	skillNames := strings.Split(skills, ",")

	for i, skill := range skillNames {
		skillNames[i] = resolveSkillAlias(skill)
	}

	svg := GenerateSVG(skillNames, theme, hasTitle, perLine)

	ctx.Data(http.StatusOK, "image/svg+xml", []byte(svg))
}

const (
	IconsPerLine = 15
	OneIcon      = 48
	Scale        = 0.1875
)

func GenerateSVG(iconNames []string, theme string, hasTitle bool, perLine int) string {
	var svg strings.Builder

	length := math.Min(float64(perLine*300), float64(len(iconNames)*300)) - 44
	height := math.Ceil(float64(len(iconNames))/float64(perLine))*300 - 44
	scaledWidth := length * Scale
	scaledHeight := height * Scale

	svg.WriteString(fmt.Sprintf(`<svg width="%f" height="%f" viewBox="0 0 %f %f" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1">`, scaledWidth, scaledHeight, length, height))

	for idx, iconName := range iconNames {
		dirPath := "."

		// Read the directory
		files, err := os.ReadDir(dirPath)
		if err != nil {
			fmt.Println("Error reading directory:", err)
		}

		// Iterate over the files
		fmt.Println("Files and folders in", dirPath+":")
		for _, file := range files {
			// Check if it's a file or a directory
			if file.IsDir() {
				fmt.Println("Directory:", file.Name())
			} else {
				fmt.Println("File:", file.Name())
			}
		}

		var title string
		if hasTitle {
			title = fmt.Sprintf("<title>%s</title>", iconName)
		}

		transformX := (idx % perLine) * 300
		transformY := (idx / perLine) * 300

		svg.WriteString(fmt.Sprintf(`
            <g transform="translate(%d, %d)">
                %s
                %s
            </g>
        `, transformX, transformY, title, string("zaza")))
	}

	svg.WriteString(`</svg>`)

	return svg.String()
}

func myRoute(r *gin.RouterGroup) {
	r.GET("/icons", func(ctx *gin.Context) {
		GetSkillIcons(ctx)
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
