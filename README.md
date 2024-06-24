<p align="center"><img align="center" width="280" src="./.github/text-logo.svg#gh-dark-mode-only"/></p>
<p align="center"><img align="center" width="280" src="./.github/text-logo-light.svg#gh-light-mode-only"/></p>
<h3 align="center">Showcase your skills on your GitHub or resumé with ease!</h3>
<hr>

# Docs

- [Docs](#docs)
- [Example](#example)
- [Specifying Icons](#specifying-icons)
- [Themed Icons](#themed-icons)
- [Icons Per Line](#icons-per-line)
- [Get Icons Names](#get-icons-names)
- [Centering Icons](#centering-icons)
- [Icons List](#icons-list)
    - [💖 Support the Project](#-support-the-project)

# Example

<p align="center"><img align="center" src="./.github/example-dark.png#gh-dark-mode-only"/></p>
<p align="center"><img align="center" src="./.github/example-light.png#gh-light-mode-only"/></p>

# Specifying Icons

Copy and paste the code block below into your readme to add the skills icon element!

Change the `?i=js,html,css` to a list of your skills separated by ","s! You can find a full list of icons [here](#icons-list).

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=js,html,css,wasm)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=js,html,css,wasm)

# Themed Icons

Some icons have a dark and light themed background. You can specify which theme you want as a url parameter.

This is optional. The default theme is dark.

Change the `&theme=light` to either `dark` or `light`. The theme is the background color, so light theme has a white icon background, and dark has a black-ish.

**Light Theme Example:**

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=java,kotlin,nodejs,figma&theme=light)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=java,kotlin,nodejs&theme=light)

# Icons Per Line

You can specify how many icons you would like per line! It's an optional argument, and the default is 15.

Change the `&perline=3` to any number between 1 and 50.

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=aws,gcp,azure,react,vue,flutter&perline=3)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=aws,gcp,azure,react,vue,flutter&perline=3)

# Get Icons Names

You can get the possiblity to add the name of the icons you put to help others that doesnt know what they are by using `&titles`.

The value of `titles` is a boolean, so it should be `true` or `false`, default is `false`

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=rust,surrealdb,actix,yew&titles=true)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=rust,surrealdb,actix,yew&titles=true)

# Centering Icons

Want to center the icons in your readme? The SVGs are automatically resized, so you can do it the same way you'd normally center an image.

```html
<p align="center">
  <a href="https://go-skill-icons.vercel.app/">
    <img src="https://go-skill-icons.vercel.app/api/icons?i=git,kubernetes,docker,c,vim" />
  </a>
</p>
```

<p align="center">
  <a href="https://go-skill-icons.vercel.app/">
    <img src="https://go-skill-icons.vercel.app/api/icons?i=git,kubernetes,docker,c,vim" />
  </a>
</p>

# Icons List

Here's a list of all the icons currently supported. Feel free to open an issue to suggest icons to add!

|      Icon ID        |                           Icon                           |
| :-----------------: | :------------------------------------------------------: |
|       `ableton`     |      <img src="./assets/ableton-auto.svg" width="48">    |
|       `acrobat`     |        <img src="./assets/acrobat.svg" width="48">       |
|    `activitypub`    |   <img src="./assets/activitypub-auto.svg" width="48">   |
|       `actix`       |      <img src="./assets/actix-auto.svg" width="48">      |
|       `adonis`      |        <img src="./assets/adonis.svg" width="48">        |
|         `ar`        |         <img src="./assets/aero.svg" width="48">         |
|         `ae`        |     <img src="./assets/aftereffects.svg" width="48">     |
|      `aiscript`     |    <img src="./assets/aiscript-auto.svg" width="48">     |
|       `alchemy`     |    <img src="./assets/alchemy-auto.svg" width="48">      |
|      `alpinejs`     |    <img src="./assets/alpinejs-auto.svg" width="48">     |
|      `anaconda`     |    <img src="./assets/anaconda-auto.svg" width="48">     |
|   `androidstudio`   |  <img src="./assets/androidstudio-auto.svg" width="48">  |
|      `angular`      |     <img src="./assets/angular-auto.svg" width="48">     |
|         `an`        |        <img src="./assets/animate.svg" width="48">       |
|      `ansible`      |       <img src="./assets/ansible.svg" width="48">        |
|        `anss`       |       <img src="./assets/anss-auto.svg" width="48">      |
|       `apollo`      |        <img src="./assets/apollo.svg" width="48">        |
|      `appcode`      |    <img src="./assets/appcode-auto.svg" width="48">      |
|       `apple`       |      <img src="./assets/apple-auto.svg" width="48">      |
|      `appwrite`     |       <img src="./assets/appwrite.svg" width="48">       |
|        `arc`        |    <img src="./assets/arcbrowser-auto.svg" width="48">   |
|        `arch`       |      <img src="./assets/arch-auto.svg" width="48">       |
|      `arduino`      |       <img src="./assets/arduino.svg" width="48">        |
|       `argocd`      |     <img src="./assets/argocd-auto.svg" width="48">      |
|        `asm`        |       <img src="./assets/assembly.svg" width="48">       |
|       `astro`       |        <img src="./assets/astro.svg" width="48">         |
|        `atom`       |         <img src="./assets/atom.svg" width="48">         |
|         `au`        |       <img src="./assets/audition.svg" width="48">       |
|      `autocad`      |     <img src="./assets/autocad-auto.svg" width="48">     |
|       `aqua`        |       <img src="./assets/aqua-auto.svg" width="48">      |
|        `aws`        |       <img src="./assets/aws-auto.svg" width="48">       |
|        `azul`       |         <img src="./assets/azul.svg" width="48">         |
|       `azure`       |      <img src="./assets/azure-auto.svg" width="48">      |
|       `babel`       |        <img src="./assets/babel.svg" width="48">         |
|        `bash`       |      <img src="./assets/bash-auto.svg" width="48">       |
|     `beeceptor`     |   <img src="./assets/beeceptor-auto.svg" width="48">     |
|         `be`        |        <img src="./assets/behance.svg" width="48">       |
|        `bevy`       |      <img src="./assets/bevy-auto.svg" width="48">       |
|     `bitbucket`     |    <img src="./assets/bitbucket-auto.svg" width="48">    |
|       `blazor`      |     <img src="./assets/blazor-auto.svg" width="48">      |
|      `blender`      |     <img src="./assets/blender-auto.svg" width="48">     |
|     `bootstrap`     |      <img src="./assets/bootstrap.svg" width="48">       |
|       `brave`       |       <img src="./assets/brave-auto.svg" width="48">     |
|       `breeze`      |       <img src="./assets/breeze.svg" width="48">         |
|         `br`        |        <img src="./assets/bridge.svg" width="48">        |
|        `bsd`        |       <img src="./assets/bsd-auto.svg" width="48">       |
|       `bulma`       |      <img src="./assets/bulma-auto.svg" width="48">      |
|        `bun`        |       <img src="./assets/bun-auto.svg" width="48">       |
|         `c`         |          <img src="./assets/c.svg" width="48">           |
|         `ca`        |       <img src="./assets/capture.svg" width="48">        |
|      `cashier`      |       <img src="./assets/cashier.svg" width="48">        |
|       `canva`       |      <img src="./assets/canva-auto.svg" width="48">      |
|     `capacitor`     |    <img src="./assets/capacitor-auto.svg" width="48">    |
|         `cs`        |          <img src="./assets/cs.svg" width="48">          |
|        `cpp`        |         <img src="./assets/cpp.svg" width="48">          |
|         `cc`        |     <img src="./assets/creativecloud.svg" width="48">    |
|      `crystal`      |     <img src="./assets/crystal-auto.svg" width="48">     |
|     `cassandra`     |    <img src="./assets/cassandra-auto.svg" width="48">    |
|         `ch`        |  <img src="./assets/characteranimator.svg" width="48">   |
|      `chatgpt`      |  <img src="./assets/chatgpt-auto.svg" width="48">        |
|       `chrome`      |      <img src="./assets/chrome-auto.svg" width="48">     |
|      `chromium`     |      <img src="./assets/chromium-auto.svg" width="48">   |
|      `circleci`     |      <img src="./assets/circleci-auto.svg" width="48">   |
|       `clion`       |      <img src="./assets/clion-auto.svg" width="48">      |
|      `clojure`      |     <img src="./assets/clojure-auto.svg" width="48">     |
|     `cloudflare`    |   <img src="./assets/cloudflare-auto.svg" width="48">    |
|       `cmake`       |      <img src="./assets/cmake-auto.svg" width="48">      |
|      `codeberg`     |     <img src="./assets/codeberg-auto.svg" width="48">    |
|     `codeigniter`   |   <img src="./assets/codeigniter-auto.svg" width="48">   |
|      `codepen`      |     <img src="./assets/codepen-auto.svg" width="48">     |
|    `coffeescript`   |  <img src="./assets/coffeescript-auto.svg" width="48">   |
|        `css`        |         <img src="./assets/css.svg" width="48">          |
|      `cypress`      |     <img src="./assets/cypress-auto.svg" width="48">     |
|         `d`         |         <img src="./assets/d.svg" width="48">            |
|         `d3`        |       <img src="./assets/d3-auto.svg" width="48">        |
|     `dailydev`      |     <img src="./assets/dailydev-auto.svg" width="48">    |
|     `databricks`    |     <img src="./assets/databricks-auto.svg" width="48">  |
|     `datadog`       |      <img src="./assets/datadog.svg" width="48">         |
|     `datagrip`      |   <img src="./assets/datagrip-auto.svg" width="48">      |
|     `dataspell`     |   <img src="./assets/dataspell-auto.svg" width="48">     |
|      `dbeaver`      |    <img src="./assets/dbeaver-auto.svg" width="48">      |
|       `dart`        |      <img src="./assets/dart-auto.svg" width="48">       |
|       `debian`      |        <img src="./assets/debian.svg" width="48">        |
|      `defold`       |     <img src="./assets/defold-auto.svg" width="48">      |
|        `deno`       |      <img src="./assets/deno-auto.svg" width="48">       |
|      `desmos`       |        <img src="./assets/desmos.svg" width="48">        |
|       `devto`       |      <img src="./assets/devto-auto.svg" width="48">      |
|         `dn`        |       <img src="./assets/dimension.svg" width="48">      |
|      `directus`     |       <img src="./assets/directus.svg" width="48">       |
|      `discord`      |       <img src="./assets/discord.svg" width="48">        |
|        `bots`       |     <img src="./assets/discordbots.svg" width="48">      |
|    `digitalocean`   |  <img src="./assets/digitalocean-light.svg" width="48">  |
|     `discordjs`     |    <img src="./assets/discordjs-auto.svg" width="48">    |
|       `django`      |        <img src="./assets/django.svg" width="48">        |
|       `docker`      |        <img src="./assets/docker.svg" width="48">        |
|      `docksal`      |     <img src="./assets/docksal-auto.svg" width="48">     |
|       `dotnet`      |        <img src="./assets/dotnet.svg" width="48">        |
|         `dw`        |     <img src="./assets/dreamweaver.svg" width="48">      |
|       `drupal`      |      <img src="./assets/drupal-auto.svg" width="48">     |
|    `duckduckgo`     |     <img src="./assets/duckduckgo.svg" width="48">       |
|       `dusk`        |      <img src="./assets/dusk.svg" width="48">            |
|      `dynamodb`     |    <img src="./assets/dynamodb-auto.svg" width="48">     |
|       `echo`        |    <img src="./assets/echo.svg" width="48">              |
|      `eclipse`      |     <img src="./assets/eclipse-auto.svg" width="48">     |
|        `edge`       |       <img src="./assets/edge-auto.svg" width="48">      |
|   `elasticsearch`   |  <img src="./assets/elasticsearch-auto.svg" width="48">  |
|      `electron`     |       <img src="./assets/electron.svg" width="48">       |
|       `elixir`      |     <img src="./assets/elixir-auto.svg" width="48">      |
|       `elysia`      |     <img src="./assets/elysia-auto.svg" width="48">      |
|       `emacs`       |        <img src="./assets/emacs.svg" width="48">         |
|       `ember`       |        <img src="./assets/ember.svg" width="48">         |
|      `emotion`      |     <img src="./assets/emotion-auto.svg" width="48">     |
|      `envoyer`      |     <img src="./assets/envoyer.svg" width="48">          |
|       `excel`       |     <img src="./assets/excel-auto.svg" width="48">       |
|      `express`      |    <img src="./assets/expressjs-auto.svg" width="48">    |
|     `facebook`      |       <img src="./assets/facebook.svg" width="48">       |
|       `fastai`      |      <img src="./assets/fastai-auto.svg" width="48">     |
|      `fastapi`      |       <img src="./assets/fastapi.svg" width="48">        |
|     `fediverse`     |    <img src="./assets/fediverse-auto.svg" width="48">    |
|       `figma`       |      <img src="./assets/figma-auto.svg" width="48">      |
|      `filament`     |       <img src="./assets/filament.svg" width="48">       |
|      `firebase`     |    <img src="./assets/firebase-auto.svg" width="48">     |
|      `firefox`      |      <img src="./assets/firefox-auto.svg" width="48">    |
|     `flameshot`     |      <img src="./assets/flameshot.svg" width="48">       |
|       `flask`       |      <img src="./assets/flask-auto.svg" width="48">      |
|       `fleet`       |       <img src="./assets/fleet-auto.svg" width="48">     |
|      `flutter`      |     <img src="./assets/flutter-auto.svg" width="48">     |
|       `fonts`       |        <img src="./assets/fonts.svg" width="48">         |
|       `forge`       |        <img src="./assets/forge.svg" width="48">         |
|       `forth`       |        <img src="./assets/forth.svg" width="48">         |
|      `fortran`      |       <img src="./assets/fortran.svg" width="48">        |
|       `fresco`      |        <img src="./assets/fresco.svg" width="48">        |
|        `fuse`       |        <img src="./assets/fuse.svg" width="48">          |
|  `gamemakerstudio`  |   <img src="./assets/gamemakerstudio.svg" width="48">    |
|      `ganache`      |     <img src="./assets/ganache-auto.svg" width="48">     |
|       `gatsby`      |        <img src="./assets/gatsby.svg" width="48">        |
|        `gcp`        |       <img src="./assets/gcp-auto.svg" width="48">       |
|       `gemini`      |      <img src="./assets/gemini-auto.svg" width="48">     |
|        `gimp`       |        <img src="./assets/gimp-auto.svg" width="48">     |
|        `git`        |         <img src="./assets/git-auto.svg" width="48">     |
|       `github`      |     <img src="./assets/github-auto.svg" width="48">      |
|   `githubactions`   |  <img src="./assets/githubactions-auto.svg" width="48">  |
|   `githubcopilot`   |  <img src="./assets/githubcopilot-auto.svg" width="48">  |
|   `gitkraken`       |  <img src="./assets/gitkraken-auto.svg" width="48">      |
|       `gitlab`      |     <img src="./assets/gitlab-auto.svg" width="48">      |
|       `gleam`       |      <img src="./assets/gleam-auto.svg" width="48">      |
|       `gmail`       |      <img src="./assets/gmail-auto.svg" width="48">      |
|       `gnome`       |      <img src="./assets/gnome-auto.svg" width="48">      |
|      `gherkin`      |     <img src="./assets/gherkin-auto.svg" width="48">     |
|         `go`        |        <img src="./assets/golang.svg" width="48">        |
|      `goland`       |     <img src="./assets/goland-auto.svg" width="48">      |
|  `googleanalytics`  | <img src="./assets/googleanalytics-auto.svg" width="48"> |
| `googleappsscript`  |<img src="./assets/googleappsscript-auto.svg" width="48"> |
|       `gradle`      |     <img src="./assets/gradle-auto.svg" width="48">      |
|       `godot`       |      <img src="./assets/godot-auto.svg" width="48">      |
|      `grafana`      |     <img src="./assets/grafana-auto.svg" width="48">     |
|       `grails`      |        <img src="./assets/grails.svg" width="48">        |
|      `graphql`      |     <img src="./assets/graphql-auto.svg" width="48">     |
|       `grunt`       |      <img src="./assets/grunt-auto.svg" width="48">      |
|        `gsap`       |       <img src="./assets/gsap-auto.svg" width="48">      |
|        `gtk`        |       <img src="./assets/gtk-auto.svg" width="48">       |
|        `gulp`       |         <img src="./assets/gulp.svg" width="48">         |
|      `hardhat`      |     <img src="./assets/hardhat-auto.svg" width="48">     |
|      `haskell`      |     <img src="./assets/haskell-auto.svg" width="48">     |
|        `haxe`       |      <img src="./assets/haxe-auto.svg" width="48">       |
|     `haxeflixel`    |   <img src="./assets/haxeflixel-auto.svg" width="48">    |
|       `helix`       |       <img src="./assets/helix-auto.svg" width="48">     |
|       `helm`        |       <img src="./assets/helm-auto.svg" width="48">      |
|       `herd`        |       <img src="./assets/herd.svg" width="48">           |
|       `heroku`      |        <img src="./assets/heroku.svg" width="48">        |
|     `hibernate`     |    <img src="./assets/hibernate-auto.svg" width="48">    |
|         `hc`        |         <img src="./assets/holyc.svg" width="48">        |
|        `hono`       |       <img src="./assets/hono-auto.svg" width="48">      |
|        `horizon`    |       <img src="./assets/horizon.svg" width="48">        |
|        `html`       |         <img src="./assets/html.svg" width="48">         |
|        `htmx`       |      <img src="./assets/htmx-auto.svg" width="48">       |
|        `htop`       |       <img src="./assets/htop-auto.svg" width="48">      |
|      `hydrogen`     |      <img src="./assets/hydrogen-auto.svg" width="48">   |
|      `hyprland`     |      <img src="./assets/hyprland-auto.svg" width="48">   |
|         `i3`        |       <img src="./assets/i3-auto.svg" width="48">        |
|        `idea`       |      <img src="./assets/idea-auto.svg" width="48">       |
|       `ignite`      |      <img src="./assets/ignite-auto.svg" width="48">     |
|         `ai`        |     <img src="./assets/illustrator.svg" width="48">      |
|         `ic`        |        <img src="./assets/incopy.svg" width="48">        |
|         `id`        |       <img src="./assets/indesign.svg" width="48">       |
|      `inertia`      |       <img src="./assets/inertia.svg" width="48">        |
|       `infura`      |        <img src="./assets/infura.svg" width="48">        |
|      `insomnia`     |       <img src="./assets/insomnia.svg" width="48">       |
|     `instagram`     |      <img src="./assets/instagram.svg" width="48">       |
|        `ipfs`       |      <img src="./assets/ipfs-auto.svg" width="48">       |
|        `java`       |      <img src="./assets/java-auto.svg" width="48">       |
|         `js`        |      <img src="./assets/javascript.svg" width="48">      |
|      `jenkins`      |     <img src="./assets/jenkins-auto.svg" width="48">     |
|        `jest`       |         <img src="./assets/jest.svg" width="48">         |
|   `jetpackcompose`  |  <img src="./assets/jetpackcompose-auto.svg" width="48"> |
|     `jetstream`     |  <img src="./assets/jetstream.svg" width="48">           |
|        `jira`       |      <img src="./assets/jira-auto.svg" width="48">       |
|       `joomla`      |      <img src="./assets/joomla-auto.svg" width="48">     |
|       `jquery`      |        <img src="./assets/jquery.svg" width="48">        |
|       `julia`       |        <img src="./assets/julia-auto.svg" width="48">    |
|      `kakoune`      |       <img src="./assets/kakoune-auto.svg" width="48">   |
|       `kafka`       |        <img src="./assets/kafka.svg" width="48">         |
|       `kaggle`      |     <img src="./assets/kaggle-auto.svg" width="48">      |
|        `kali`       |      <img src="./assets/kali-auto.svg" width="48">       |
|        `kde`        |       <img src="./assets/kde-auto.svg" width="48">       |
|      `keycloak`     |       <img src="./assets/keycloak.svg" width="48">       |
|       `kotlin`      |     <img src="./assets/kotlin-auto.svg" width="48">      |
|        `ktor`       |      <img src="./assets/ktor-auto.svg" width="48">       |
|     `kubernetes`    |      <img src="./assets/kubernetes.svg" width="48">      |
|     `langchain`     |    <img src="./assets/langchain-auto.svg" width="48">    |
|      `laravel`      |     <img src="./assets/laravel-auto.svg" width="48">     |
|   `laravelspark`    |  <img src="./assets/laravelspark-auto.svg" width="48">   |
|       `latex`       |      <img src="./assets/latex-auto.svg" width="48">      |
|      `leetcode`     |    <img src="./assets/leetcode-auto.svg" width="48">     |
|        `less`       |      <img src="./assets/less-auto.svg" width="48">       |
|     `lightning`     |    <img src="./assets/lightning-auto.svg" width="48">    |
|         `lr`        |      <img src="./assets/lightroom.svg" width="48">       |
|         `lrc`       |  <img src="./assets/lightroomclassic.svg" width="48">    |
|      `linkedin`     |       <img src="./assets/linkedin.svg" width="48">       |
|       `linux`       |      <img src="./assets/linux-auto.svg" width="48">      |
|        `lit`        |       <img src="./assets/lit-auto.svg" width="48">       |
|       `litmus`      |       <img src="./assets/litmus-auto.svg" width="48">    |
|      `livewire`     |      <img src="./assets/livewire-auto.svg" width="48">   |
|       `looker`      |        <img src="./assets/looker-auto.svg" width="48">   |
|        `lua`        |       <img src="./assets/lua-auto.svg" width="48">       |
|     `lucidchart`    |   <img src="./assets/lucidchart-auto.svg" width="48">    |
|         `md`        |    <img src="./assets/markdown-auto.svg" width="48">     |
|      `manjaro`      |       <img src="./assets/manjaro.svg" width="48">        |
|      `mastodon`     |    <img src="./assets/mastodon-auto.svg" width="48">     |
|     `materialui`    |   <img src="./assets/materialui-auto.svg" width="48">    |
|       `matlab`      |     <img src="./assets/matlab-auto.svg" width="48">      |
|     `matplotlib`    |    <img src="./assets/matplotlib-auto.svg" width="48">   |
|       `maven`       |      <img src="./assets/maven-auto.svg" width="48">      |
|         `me`        |     <img src="./assets/mediaencoder.svg" width="48">     |
|       `mermaid`     |       <img src="./assets/mermaid.svg" width="48">        |
|      `meteorjs`     |    <img src="./assets/meteorjs-auto.svg" width="48">     |
|  `microsoftcopilot` | <img src="./assets/microsoftcopilot-auto.svg" width="48">|
|      `million`      |    <img src="./assets/millionjs-auto.svg" width="48">    |
|        `mint`       |      <img src="./assets/mint-auto.svg" width="48">       |
|        `miro`       |         <img src="./assets/miro.svg" width="48">         |
|      `misskey`      |     <img src="./assets/misskey-auto.svg" width="48">     |
|        `mjml`       |        <img src="./assets/mjml-auto.svg" width="48">     |
|        `ml5`        |     <img src="./assets/ml5-auto.svg" width="48">         |
|        `mojo`       |     <img src="./assets/mojo-auto.svg" width="48">        |
|      `mongodb`      |       <img src="./assets/mongodb.svg" width="48">        |
|       `mysql`       |      <img src="./assets/mysql-auto.svg" width="48">      |
|       `neovim`      |     <img src="./assets/neovim-auto.svg" width="48">      |
|       `nestjs`      |     <img src="./assets/nestjs-auto.svg" width="48">      |
|      `netlify`      |     <img src="./assets/netlify-auto.svg" width="48">     |
|       `nextjs`      |     <img src="./assets/nextjs-auto.svg" width="48">      |
|       `nginx`       |        <img src="./assets/nginx.svg" width="48">         |
|       `ngrok`       |        <img src="./assets/ngrok.svg" width="48">         |
|        `nim`        |       <img src="./assets/nim-auto.svg" width="48">       |
|        `nix`        |      <img src="./assets/nixos-auto.svg" width="48">      |
|       `nodejs`      |     <img src="./assets/nodejs-auto.svg" width="48">      |
|       `notion`      |     <img src="./assets/notion-auto.svg" width="48">      |
|       `nova`        |     <img src="./assets/nova.svg" width="48">             |
|        `npm`        |       <img src="./assets/npm-auto.svg" width="48">       |
|        `numpy`      |       <img src="./assets/numpy-auto.svg" width="48">     |
|       `nuxtjs`      |     <img src="./assets/nuxtjs-auto.svg" width="48">      |
|      `obsidian`     |    <img src="./assets/obsidian-auto.svg" width="48">     |
|       `ocaml`       |        <img src="./assets/ocaml.svg" width="48">         |
|       `octane`      |        <img src="./assets/octane.svg" width="48">        |
|       `octave`      |     <img src="./assets/octave-auto.svg" width="48">      |
|       `ollama`      |     <img src="./assets/ollama-auto.svg" width="48">      |
|      `onedrive`     |     <img src="./assets/onedrive-auto.svg" width="48">    |
|       `onenote`     |     <img src="./assets/onenote-auto.svg" width="48">     |
|       `opencv`      |     <img src="./assets/opencv-auto.svg" width="48">      |
|     `openshift`     |      <img src="./assets/openshift.svg" width="48">       |
|     `openstack`     |    <img src="./assets/openstack-auto.svg" width="48">    |
|    `openzeppelin`   |   <img src="./assets/openzeppelin-auto.svg" width="48">  |
|       `opera`       |       <img src="./assets/opera-auto.svg" width="48">     |
|       `oracle`      |       <img src="./assets/oracle-auto.svg" width="48">    |
|      `outlook`      |       <img src="./assets/outlook-auto.svg" width="48">   |
|        `p5js`       |         <img src="./assets/p5js.svg" width="48">         |
|        `pail`       |         <img src="./assets/pail-auto.svg" width="48">    |
|       `pandas`      |      <img src="./assets/pandas-auto.svg" width="48">     |
|     `papertrail`    |       <img src="./assets/papertrail.svg" width="48">     |
|       `payload`     |      <img src="./assets/payload-auto.svg" width="48">    |
|        `pbi`        |        <img src="./assets/pbi-auto.svg" width="48">      |
|       `pennant`     |      <img src="./assets/pennant.svg" width="48">         |
|        `perl`       |         <img src="./assets/perl.svg" width="48">         |
|         `ps`        |      <img src="./assets/photoshop.svg" width="48">       |
|        `psc`        |  <img src="./assets/photoshopclassic.svg" width="48">    |
|        `psx`        |  <img src="./assets/photoshopexpress.svg" width="48">    |
|        `php`        |       <img src="./assets/php-auto.svg" width="48">       |
|      `phpstorm`     |    <img src="./assets/phpstorm-auto.svg" width="48">     |
|      `picocss`      |    <img src="./assets/picocss-auto.svg" width="48">      |
|      `pinecone`     |    <img src="./assets/pinecone-auto.svg" width="48">     |
|       `pinia`       |      <img src="./assets/pinia-auto.svg" width="48">      |
|       `pint`        |      <img src="./assets/pint.svg" width="48">            |
|        `pkl`        |       <img src="./assets/pkl-auto.svg" width="48">       |
|       `plan9`       |      <img src="./assets/plan9-auto.svg" width="48">      |
|    `planetscale`    |   <img src="./assets/planetscale-auto.svg" width="48">   |
|     `playwright`    |     <img src="./assets/playwright-auto.svg" width="48">  |
|        `pnpm`       |      <img src="./assets/pnpm-auto.svg" width="48">       |
|        `pop`        |         <img src="./assets/popos.svg" width="48">        |
|     `pocketbase`    |   <img src="./assets/pocketbase-auto.svg" width="48">    |
|         `pf`        |       <img src="./assets/portfolio.svg" width="48">      |
|      `postgres`     |   <img src="./assets/postgresql-auto.svg" width="48">    |
|      `postman`      |       <img src="./assets/postman.svg" width="48">        |
|     `powerpoint`    |   <img src="./assets/powerpoint-auto.svg" width="48">    |
|     `powershell`    |   <img src="./assets/powershell-auto.svg" width="48">    |
|       `preact`      |       <img src="./assets/preact-auto.svg" width="48">    |
|         `pl`        |       <img src="./assets/prelude.svg" width="48">        |
|         `pr`        |       <img src="./assets/premiere.svg" width="48">       |
|         `ru`        |     <img src="./assets/premiererush.svg" width="48">     |
|       `prisma`      |        <img src="./assets/prisma.svg" width="48">        |
|     `processing`    |   <img src="./assets/processing-auto.svg" width="48">    |
|     `prometheus`    |      <img src="./assets/prometheus.svg" width="48">      |
|     `prompts`       |      <img src="./assets/prompts.svg" width="48">         |
|       `proton`      |      <img src="./assets/proton-auto.svg" width="48">     |
|      `proxmox`      |      <img src="./assets/proxmox-auto.svg" width="48">    |
|        `pug`        |       <img src="./assets/pug-auto.svg" width="48">       |
|      `pulse`        |       <img src="./assets/pulse-auto.svg" width="48">     |
|      `puppeteer`    |    <img src="./assets/puppeteer-auto.svg" width="48">    |
|      `pycharm`      |     <img src="./assets/pycharm-auto.svg" width="48">     |
|         `py`        |     <img src="./assets/python-auto.svg" width="48">      |
|      `pytorch`      |     <img src="./assets/pytorch-auto.svg" width="48">     |
|       `pyxel`       |       <img src="./assets/pyxel-auto.svg" width="48">     |
|      `qodana`       |      <img src="./assets/qodana-auto.svg" width="48">     |
|         `qt`        |       <img src="./assets/qt-auto.svg" width="48">        |
|         `r`         |        <img src="./assets/r-auto.svg" width="48">        |
|      `rabbitmq`     |    <img src="./assets/rabbitmq-auto.svg" width="48">     |
|       `rails`       |        <img src="./assets/rails.svg" width="48">         |
|    `raspberrypi`    |   <img src="./assets/raspberrypi-auto.svg" width="48">   |
|       `react`       |      <img src="./assets/react-auto.svg" width="48">      |
|     `reactivex`     |    <img src="./assets/reactivex-auto.svg" width="48">    |
|    `reactquery`     |    <img src="./assets/reactquery-auto.svg" width="48">   |
|      `recoil`       |         <img src="./assets/recoil.svg" width="48">       |
|       `redhat`      |     <img src="./assets/redhat-auto.svg" width="48">      |
|       `redis`       |      <img src="./assets/redis-auto.svg" width="48">      |
|      `redshift`     |   <img src="./assets/redshift-auto.svg" width="48">      |
|       `redux`       |        <img src="./assets/redux.svg" width="48">         |
|       `regex`       |      <img src="./assets/regex-auto.svg" width="48">      |
|       `remix`       |      <img src="./assets/remix-auto.svg" width="48">      |
|       `replit`      |     <img src="./assets/replit-auto.svg" width="48">      |
|     `resharper`     |   <img src="./assets/resharper-auto.svg" width="48">     |
|     `reverb`        |   <img src="./assets/reverb.svg" width="48">             |
|       `rider`       |      <img src="./assets/rider-auto.svg" width="48">      |
|    `robloxstudio`   |     <img src="./assets/robloxstudio.svg" width="48">     |
|       `rocket`      |        <img src="./assets/rocket.svg" width="48">        |
|      `rollupjs`     |    <img src="./assets/rollupjs-auto.svg" width="48">     |
|        `ros`        |       <img src="./assets/ros-auto.svg" width="48">       |
|      `rubocop`      |     <img src="./assets/rubocop-auto.svg" width="48">     |
|        `ruby`       |         <img src="./assets/ruby.svg" width="48">         |
|     `rubymine`      |    <img src="./assets/rubymine-auto.svg" width="48">     |
|        `rust`       |         <img src="./assets/rust-auto.svg" width="48">    |
|     `rustrover`     |     <img src="./assets/rustrover-auto.svg" width="48">   |
|      `safari`       |       <img src="./assets/safari-auto.svg" width="48">    |
|      `sail`         |       <img src="./assets/sail.svg" width="48">           |
|     `sanctum`       |       <img src="./assets/sanctum.svg" width="48">        |
|        `sass`       |         <img src="./assets/sass.svg" width="48">         |
|       `spring`      |     <img src="./assets/spring-auto.svg" width="48">      |
|       `sqlite`      |        <img src="./assets/sqlite.svg" width="48">        |
|      `sqlserver`    |      <img src="./assets/sqlserver-auto.svg" width="48">  |
|   `stackoverflow`   |  <img src="./assets/stackoverflow-auto.svg" width="48">  |
|       `stock`       |         <img src="./assets/stock.svg" width="48">        |
|     `storyblok`     |     <img src="./assets/storyblok-auto.svg" width="48">   |
|     `storybook`     |     <img src="./assets/storybook-auto.svg" width="48">   |
|      `strapi`       |         <img src="./assets/strapi.svg" width="48">       |
|     `streamlit`     |     <img src="./assets/streamlit-auto.svg" width="48">   |
|  `styledcomponents` |   <img src="./assets/styledcomponents.svg" width="48">   |
|       `stylus`      |      <img src="./assets/stylus-auto.svg" width="48">     |
|      `sublime`      |     <img src="./assets/sublime-auto.svg" width="48">     |
|      `supabase`     |    <img src="./assets/supabase-auto.svg" width="48">     |
|     `surrealdb`     |   <img src="./assets/surrealdb-auto.svg" width="48">     |
|       `scala`       |      <img src="./assets/scala-auto.svg" width="48">      |
|        `scipy`      |       <img src="./assets/scipy-auto.svg" width="48">     |
|      `scout`        |       <img src="./assets/scout.svg" width="48">          |
|      `scratch`      |       <img src="./assets/scratch.svg" width="48">        |
|      `seaborn`      |     <img src="./assets/seaborn-auto.svg" width="48">     |
|      `sklearn`      |   <img src="./assets/scikitlearn-auto.svg" width="48">   |
|      `selenium`     |       <img src="./assets/selenium.svg" width="48">       |
|       `sentry`      |        <img src="./assets/sentry.svg" width="48">        |
|     `sequelize`     |    <img src="./assets/sequelize-auto.svg" width="48">    |
|    `sharepoint`     |    <img src="./assets/sharepoint-auto.svg" width="48">   |
|      `shopify`      |       <img src="./assets/shopify-auto.svg" width="48">   |
|     `skeletonui`    |    <img src="./assets/skeletonui-auto.svg" width="48">   |
|      `sketchup`     |    <img src="./assets/sketchup-auto.svg" width="48">     |
|       `slack`       |    <img src="./assets/slack-auto.svg" width="48">        |
|      `snowflake`    |    <img src="./assets/snowflake-auto.svg" width="48">    |
|        `snyk`       |     <img src="./assets/snyk-auto.svg" width="48">        |
|     `socialite`     |     <img src="./assets/socialite.svg" width="48">        |
|      `solidity`     |       <img src="./assets/solidity.svg" width="48">       |
|      `solidjs`      |     <img src="./assets/solidjs-auto.svg" width="48">     |
|       `spark`       |         <img src="./assets/spark.svg" width="48">        |
|       `svelte`      |        <img src="./assets/svelte.svg" width="48">        |
|        `svg`        |       <img src="./assets/svg-auto.svg" width="48">       |
|        `svn`        |         <img src="./assets/svn.svg" width="48">          |
|      `swagger`      |       <img src="./assets/swagger-auto.svg" width="48">   |
|       `swift`       |        <img src="./assets/swift.svg" width="48">         |
|      `symfony`      |     <img src="./assets/symfony-auto.svg" width="48">     |
|      `tableau`      |      <img src="./assets/tableau-auto.svg" width="48">    |
|      `tailwind`     |   <img src="./assets/tailwindcss-auto.svg" width="48">   |
|     `tallyprime`    |      <img src="./assets/tallyprime.svg" width="48">      |
|       `tauri`       |      <img src="./assets/tauri-auto.svg" width="48">      |
|       `teams`       |      <img src="./assets/teams-auto.svg" width="48">      |
|      `telescope`    |      <img src="./assets/telescope.svg" width="48">       |
|     `tensorflow`    |   <img src="./assets/tensorflow-auto.svg" width="48">    |
|     `terraform`     |    <img src="./assets/terraform-auto.svg" width="48">    |
|   `testinglibrary`  |  <img src="./assets/testinglibrary-auto.svg" width="48"> |
|      `threejs`      |     <img src="./assets/threejs-auto.svg" width="48">     |
|        `tidb`       |      <img src="./assets/tidb-auto.svg" width="48">       |
|        `tor`        |       <img src="./assets/tor-auto.svg" width="48">       |
|       `trpc`        |       <img src="./assets/trpc.svg" width="48">           |
|      `truffle`      |     <img src="./assets/truffle-auto.svg" width="48">     |
|         `ts`        |      <img src="./assets/typescript.svg" width="48">      |
|       `ubuntu`      |        <img src="./assets/ubuntu.svg" width="48">        |
|       `unity`       |      <img src="./assets/unity-auto.svg" width="48">      |
|       `unreal`      |     <img src="./assets/unrealengine.svg" width="48">     |
|         `v`         |        <img src="./assets/v-auto.svg" width="48">        |
|        `vala`       |         <img src="./assets/vala.svg" width="48">         |
|        `vapor`      |         <img src="./assets/vapor.svg" width="48">        |
|         `vb`        |   <img src="./assets/visualbasic-auto.svg" width="48">   |
|       `vercel`      |     <img src="./assets/vercel-auto.svg" width="48">      |
|        `vim`        |       <img src="./assets/vim-auto.svg" width="48">       |
|    `visualstudio`   |  <img src="./assets/visualstudio-auto.svg" width="48">   |
|        `vite`       |      <img src="./assets/vite-auto.svg" width="48">       |
|       `vitest`      |     <img src="./assets/vitest-auto.svg" width="48">      |
|       `vscode`      |     <img src="./assets/vscode-auto.svg" width="48">      |
|      `vscodium`     |    <img src="./assets/vscodium-auto.svg" width="48">     |
|        `vue`        |      <img src="./assets/vuejs-auto.svg" width="48">      |
|      `vuetify`      |     <img src="./assets/vuetify-auto.svg" width="48">     |
|       `vyper`       |      <img src="./assets/vyper-auto.svg" width="48">      |
|        `wasm`       |     <img src="./assets/webassembly.svg" width="48">      |
|      `webflow`      |       <img src="./assets/webflow.svg" width="48">        |
|      `webpack`      |     <img src="./assets/webpack-auto.svg" width="48">     |
|      `webstorm`     |    <img src="./assets/webstorm-auto.svg" width="48">     |
|      `windicss`     |    <img src="./assets/windicss-auto.svg" width="48">     |
|      `windows`      |     <img src="./assets/windows-auto.svg" width="48">     |
|        `word`       |       <img src="./assets/word-auto.svg" width="48">      |
|     `wordpress`     |      <img src="./assets/wordpress.svg" width="48">       |
|      `workers`      |     <img src="./assets/workers-auto.svg" width="48">     |
|         `x`         |         <img src="./assets/x-auto.svg" width="48">       |
|       `xcode`       |       <img src="./assets/xcode-auto.svg" width="48">     |
|         `xd`        |          <img src="./assets/xd.svg" width="48">          |
|        `yaml`       |     <img src="./assets/yaml-auto.svg" width="48">        |
|       `yammer`      |     <img src="./assets/yammer-auto.svg" width="48">      |
|        `yarn`       |      <img src="./assets/yarn-auto.svg" width="48">       |
|        `yew`        |       <img src="./assets/yew-auto.svg" width="48">       |
|      `youtube`      |        <img src="./assets/youtube.svg" width="48">       |
|        `zed`        |       <img src="./assets/zed-auto.svg" width="48">       |
|        `zig`        |       <img src="./assets/zig-auto.svg" width="48">       |

---

## 💖 Support the Project

Thank you so much already for using my projects!

To support the project directly, feel free to open issues for icon suggestions, or contribute with a pull request!
