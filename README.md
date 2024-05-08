<p align="center"><img align="center" width="280" src="./.github/text-logo.svg#gh-dark-mode-only"/></p>
<p align="center"><img align="center" width="280" src="./.github/text-logo-light.svg#gh-light-mode-only"/></p>
<h3 align="center">Showcase your skills on your GitHub or resumé with ease!</h3>
<hr>

# Docs

- [Example](#example)
- [Specifying Icons](#specifying-icons)
- [Themed Icons](#themed-icons)
- [Icons Per Line](#icons-per-line)
- [Centering Icons](#centering-icons)
- [Icons List](#icons-list)

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

|      Icon ID        |                          Icon                           |
| :-----------------: | :-----------------------------------------------------: |
|       `ableton`     |      <img src="./assets/ableton-dark.svg" width="48">    |
|       `acrobat`     |        <img src="./assets/acrobat.svg" width="48">       |
|    `activitypub`    |   <img src="./assets/activitypub-dark.svg" width="48">   |
|       `actix`       |      <img src="./assets/actix-dark.svg" width="48">      |
|       `adonis`      |        <img src="./assets/adonis.svg" width="48">        |
|         `ar`        |         <img src="./assets/aero.svg" width="48">         |
|         `ae`        |     <img src="./assets/aftereffects.svg" width="48">     |
|      `aiscript`     |    <img src="./assets/aiscript-dark.svg" width="48">     |
|      `alpinejs`     |    <img src="./assets/alpinejs-dark.svg" width="48">     |
|      `anaconda`     |    <img src="./assets/anaconda-dark.svg" width="48">     |
|   `androidstudio`   |  <img src="./assets/androidstudio-dark.svg" width="48">  |
|      `angular`      |     <img src="./assets/angular-dark.svg" width="48">     |
|         `an`        |        <img src="./assets/animate.svg" width="48">       |
|      `ansible`      |       <img src="./assets/ansible.svg" width="48">        |
|        `anss`       |       <img src="./assets/anss-dark.svg" width="48">      |
|       `apollo`      |        <img src="./assets/apollo.svg" width="48">        |
|       `apple`       |      <img src="./assets/apple-dark.svg" width="48">      |
|      `appwrite`     |       <img src="./assets/appwrite.svg" width="48">       |
|        `arch`       |      <img src="./assets/arch-dark.svg" width="48">       |
|      `arduino`      |       <img src="./assets/arduino.svg" width="48">        |
|        `asm`        |       <img src="./assets/assembly.svg" width="48">       |
|       `astro`       |        <img src="./assets/astro.svg" width="48">         |
|        `atom`       |         <img src="./assets/atom.svg" width="48">         |
|         `au`        |       <img src="./assets/audition.svg" width="48">       |
|      `autocad`      |     <img src="./assets/autocad-dark.svg" width="48">     |
|        `aws`        |       <img src="./assets/aws-dark.svg" width="48">       |
|        `azul`       |         <img src="./assets/azul.svg" width="48">         |
|       `azure`       |      <img src="./assets/azure-dark.svg" width="48">      |
|       `babel`       |        <img src="./assets/babel.svg" width="48">         |
|        `bash`       |      <img src="./assets/bash-dark.svg" width="48">       |
|         `be`        |        <img src="./assets/behance.svg" width="48">       |
|        `bevy`       |      <img src="./assets/bevy-dark.svg" width="48">       |
|     `bitbucket`     |    <img src="./assets/bitbucket-dark.svg" width="48">    |
|      `blender`      |     <img src="./assets/blender-dark.svg" width="48">     |
|     `bootstrap`     |      <img src="./assets/bootstrap.svg" width="48">       |
|       `brave`       |       <img src="./assets/brave-dark.svg" width="48">     |
|         `br`        |        <img src="./assets/bridge.svg" width="48">        |
|        `bsd`        |       <img src="./assets/bsd-dark.svg" width="48">       |
|        `bun`        |       <img src="./assets/bun-dark.svg" width="48">       |
|         `c`         |          <img src="./assets/c.svg" width="48">           |
|         `ca`        |       <img src="./assets/capture.svg" width="48">        |
|         `cs`        |          <img src="./assets/cs.svg" width="48">          |
|        `cpp`        |         <img src="./assets/cpp.svg" width="48">          |
|         `cc`        |     <img src="./assets/creativecloud.svg" width="48">    |
|      `crystal`      |     <img src="./assets/crystal-dark.svg" width="48">     |
|     `cassandra`     |    <img src="./assets/cassandra-dark.svg" width="48">    |
|         `ch`        |  <img src="./assets/characteranimator.svg" width="48">   |
|       `chrome`      |      <img src="./assets/chrome-dark.svg" width="48">     |
|      `chromium`     |      <img src="./assets/chromium-dark.svg" width="48">   |
|       `clion`       |      <img src="./assets/clion-dark.svg" width="48">      |
|      `clojure`      |     <img src="./assets/clojure-dark.svg" width="48">     |
|     `cloudflare`    |   <img src="./assets/cloudflare-dark.svg" width="48">    |
|       `cmake`       |      <img src="./assets/cmake-dark.svg" width="48">      |
|      `codepen`      |     <img src="./assets/codepen-dark.svg" width="48">     |
|    `coffeescript`   |  <img src="./assets/coffeescript-dark.svg" width="48">   |
|        `css`        |         <img src="./assets/css.svg" width="48">          |
|      `cypress`      |     <img src="./assets/cypress-dark.svg" width="48">     |
|         `d3`        |       <img src="./assets/d3-dark.svg" width="48">        |
|        `dart`       |      <img src="./assets/dart-dark.svg" width="48">       |
|       `debian`      |     <img src="./assets/debian-dark.svg" width="48">      |
|      `defold`       |     <img src="./assets/defold-dark.svg" width="48">      |
|        `deno`       |      <img src="./assets/deno-dark.svg" width="48">       |
|       `devto`       |      <img src="./assets/devto-dark.svg" width="48">      |
|         `dn`        |       <img src="./assets/dimension.svg" width="48">      |
|      `discord`      |       <img src="./assets/discord.svg" width="48">        |
|        `bots`       |     <img src="./assets/discordbots.svg" width="48">      |
|     `discordjs`     |    <img src="./assets/discordjs-dark.svg" width="48">    |
|       `django`      |        <img src="./assets/django.svg" width="48">        |
|       `docker`      |        <img src="./assets/docker.svg" width="48">        |
|      `docksal`      |     <img src="./assets/docksal-dark.svg" width="48">     |
|       `dotnet`      |        <img src="./assets/dotnet.svg" width="48">        |
|         `dw`        |     <img src="./assets/dreamweaver.svg" width="48">      |
|       `drupal`      |      <img src="./assets/drupal-dark.svg" width="48">     |
|      `dynamodb`     |    <img src="./assets/dynamodb-dark.svg" width="48">     |
|      `eclipse`      |     <img src="./assets/eclipse-dark.svg" width="48">     |
|        `edge`       |       <img src="./assets/edge-dark.svg" width="48">      |
|   `elasticsearch`   |  <img src="./assets/elasticsearch-dark.svg" width="48">  |
|      `electron`     |       <img src="./assets/electron.svg" width="48">       |
|       `elixir`      |     <img src="./assets/elixir-dark.svg" width="48">      |
|       `elysia`      |     <img src="./assets/elysia-dark.svg" width="48">      |
|       `emacs`       |        <img src="./assets/emacs.svg" width="48">         |
|       `ember`       |        <img src="./assets/ember.svg" width="48">         |
|      `emotion`      |     <img src="./assets/emotion-dark.svg" width="48">     |
|       `excel`       |     <img src="./assets/excel-dark.svg" width="48">       |
|      `express`      |    <img src="./assets/expressjs-dark.svg" width="48">    |
|      `fastapi`      |       <img src="./assets/fastapi.svg" width="48">        |
|     `fediverse`     |    <img src="./assets/fediverse-dark.svg" width="48">    |
|       `figma`       |      <img src="./assets/figma-dark.svg" width="48">      |
|      `firebase`     |    <img src="./assets/firebase-dark.svg" width="48">     |
|      `firefox`      |      <img src="./assets/firefox-dark.svg" width="48">    |
|       `flask`       |      <img src="./assets/flask-dark.svg" width="48">      |
|      `flutter`      |     <img src="./assets/flutter-dark.svg" width="48">     |
|       `fonts`       |        <img src="./assets/fonts.svg" width="48">         |
|       `forth`       |        <img src="./assets/forth.svg" width="48">         |
|      `fortran`      |       <img src="./assets/fortran.svg" width="48">        |
|       `fresco`      |        <img src="./assets/fresco.svg" width="48">        |
|        `fuse`       |        <img src="./assets/fuse.svg" width="48">          |
|  `gamemakerstudio`  |   <img src="./assets/gamemakerstudio.svg" width="48">    |
|       `gatsby`      |        <img src="./assets/gatsby.svg" width="48">        |
|        `gcp`        |       <img src="./assets/gcp-dark.svg" width="48">       |
|        `git`        |         <img src="./assets/git-dark.svg" width="48">     |
|       `github`      |     <img src="./assets/github-dark.svg" width="48">      |
|   `githubactions`   |  <img src="./assets/githubactions-dark.svg" width="48">  |
|       `gitlab`      |     <img src="./assets/gitlab-dark.svg" width="48">      |
|       `gmail`       |      <img src="./assets/gmail-dark.svg" width="48">      |
|       `gleam`       |      <img src="./assets/gleam-dark.svg" width="48">      |
|      `gherkin`      |     <img src="./assets/gherkin-dark.svg" width="48">     |
|         `go`        |        <img src="./assets/golang.svg" width="48">        |
|  `googleanalytics`  | <img src="./assets/googleanalytics-dark.svg" width="48"> |
| `googleappsscript`  |<img src="./assets/googleappsscript-dark.svg" width="48"> |
|       `gradle`      |     <img src="./assets/gradle-dark.svg" width="48">      |
|       `godot`       |      <img src="./assets/godot-dark.svg" width="48">      |
|      `grafana`      |     <img src="./assets/grafana-dark.svg" width="48">     |
|      `graphql`      |     <img src="./assets/graphql-dark.svg" width="48">     |
|        `gtk`        |       <img src="./assets/gtk-dark.svg" width="48">       |
|        `gulp`       |         <img src="./assets/gulp.svg" width="48">         |
|      `haskell`      |     <img src="./assets/haskell-dark.svg" width="48">     |
|        `haxe`       |      <img src="./assets/haxe-dark.svg" width="48">       |
|     `haxeflixel`    |   <img src="./assets/haxeflixel-dark.svg" width="48">    |
|       `helix`       |       <img src="./assets/helix-dark.svg" width="48">     |
|       `heroku`      |        <img src="./assets/heroku.svg" width="48">        |
|     `hibernate`     |    <img src="./assets/hibernate-dark.svg" width="48">    |
|        `html`       |         <img src="./assets/html.svg" width="48">         |
|        `htmx`       |      <img src="./assets/htmx-dark.svg" width="48">       |
|      `hydrogen`     |      <img src="./assets/hydrogen-dark.svg" width="48">   |
|        `idea`       |      <img src="./assets/idea-dark.svg" width="48">       |
|         `ai`        |     <img src="./assets/illustrator.svg" width="48">      |
|         `ic`        |        <img src="./assets/incopy.svg" width="48">        |
|         `id`        |       <img src="./assets/indesign.svg" width="48">       |
|     `instagram`     |      <img src="./assets/instagram.svg" width="48">       |
|        `ipfs`       |      <img src="./assets/ipfs-dark.svg" width="48">       |
|        `java`       |      <img src="./assets/java-dark.svg" width="48">       |
|         `js`        |      <img src="./assets/javascript.svg" width="48">      |
|      `jenkins`      |     <img src="./assets/jenkins-dark.svg" width="48">     |
|        `jest`       |         <img src="./assets/jest.svg" width="48">         |
|       `jquery`      |        <img src="./assets/jquery.svg" width="48">        |
|       `julia`       |        <img src="./assets/julia-dark.svg" width="48">    |
|      `kakoune`      |       <img src="./assets/kakoune-dark.svg" width="48">   |
|       `kafka`       |        <img src="./assets/kafka.svg" width="48">         |
|        `kali`       |      <img src="./assets/kali-dark.svg" width="48">       |
|       `kotlin`      |     <img src="./assets/kotlin-dark.svg" width="48">      |
|        `ktor`       |      <img src="./assets/ktor-dark.svg" width="48">       |
|     `kubernetes`    |      <img src="./assets/kubernetes.svg" width="48">      |
|      `laravel`      |     <img src="./assets/laravel-dark.svg" width="48">     |
|       `latex`       |      <img src="./assets/latex-dark.svg" width="48">      |
|      `leetcode`     |    <img src="./assets/leetcode-dark.svg" width="48">     |
|        `less`       |      <img src="./assets/less-dark.svg" width="48">       |
|         `lr`        |      <img src="./assets/lightroom.svg" width="48">       |
|         `lrc`       |  <img src="./assets/lightroomclassic.svg" width="48">    |
|      `linkedin`     |       <img src="./assets/linkedin.svg" width="48">       |
|       `linux`       |      <img src="./assets/linux-dark.svg" width="48">      |
|        `lit`        |       <img src="./assets/lit-dark.svg" width="48">       |
|       `litmus`      |       <img src="./assets/litmus-dark.svg" width="48">    |
|       `looker`      |        <img src="./assets/looker-dark.svg" width="48">   |
|        `lua`        |       <img src="./assets/lua-dark.svg" width="48">       |
|         `md`        |    <img src="./assets/markdown-dark.svg" width="48">     |
|      `mastodon`     |    <img src="./assets/mastodon-dark.svg" width="48">     |
|     `materialui`    |   <img src="./assets/materialui-dark.svg" width="48">    |
|       `matlab`      |     <img src="./assets/matlab-dark.svg" width="48">      |
|     `matplotlib`    |    <img src="./assets/matplotlib-dark.svg" width="48">   |
|       `maven`       |      <img src="./assets/maven-dark.svg" width="48">      |
|         `me`        |     <img src="./assets/mediaencoder.svg" width="48">     |
|      `million`      |    <img src="./assets/millionjs-dark.svg" width="48">    |
|        `mint`       |      <img src="./assets/mint-dark.svg" width="48">       |
|      `misskey`      |     <img src="./assets/misskey-dark.svg" width="48">     |
|        `mjml`       |        <img src="./assets/mjml-dark.svg" width="48">     |
|        `ml5`        |     <img src="./assets/ml5-dark.svg" width="48">         |
|      `mongodb`      |       <img src="./assets/mongodb.svg" width="48">        |
|       `mysql`       |      <img src="./assets/mysql-dark.svg" width="48">      |
|       `neovim`      |     <img src="./assets/neovim-dark.svg" width="48">      |
|       `nestjs`      |     <img src="./assets/nestjs-dark.svg" width="48">      |
|      `netlify`      |     <img src="./assets/netlify-dark.svg" width="48">     |
|       `nextjs`      |     <img src="./assets/nextjs-dark.svg" width="48">      |
|       `nginx`       |        <img src="./assets/nginx.svg" width="48">         |
|        `nim`        |       <img src="./assets/nim-dark.svg" width="48">       |
|        `nix`        |       <img src="./assets/nix-dark.svg" width="48">       |
|       `nodejs`      |     <img src="./assets/nodejs-dark.svg" width="48">      |
|       `notion`      |     <img src="./assets/notion-dark.svg" width="48">      |
|        `npm`        |       <img src="./assets/npm-dark.svg" width="48">       |
|        `numpy`      |       <img src="./assets/numpy-dark.svg" width="48">     |
|       `nuxtjs`      |     <img src="./assets/nuxtjs-dark.svg" width="48">      |
|      `obsidian`     |    <img src="./assets/obsidian-dark.svg" width="48">     |
|       `ocaml`       |        <img src="./assets/ocaml.svg" width="48">         |
|       `octave`      |     <img src="./assets/octave-dark.svg" width="48">      |
|      `onedrive`     |     <img src="./assets/onedrive-dark.svg" width="48">    |
|       `onenote`     |     <img src="./assets/onenote-dark.svg" width="48">     |
|       `opencv`      |     <img src="./assets/opencv-dark.svg" width="48">      |
|     `openshift`     |      <img src="./assets/openshift.svg" width="48">       |
|     `openstack`     |    <img src="./assets/openstack-dark.svg" width="48">    |
|       `opera`       |       <img src="./assets/opera-dark.svg" width="48">     |
|       `outlook`     |       <img src="./assets/outlook-dark.svg" width="48">     |
|        `p5js`       |         <img src="./assets/p5js.svg" width="48">         |
|       `pandas`      |      <img src="./assets/pandas-dark.svg" width="48">     |
|        `pbi`        |        <img src="./assets/pbi-dark.svg" width="48">      |
|        `perl`       |         <img src="./assets/perl.svg" width="48">         |
|         `ps`        |      <img src="./assets/photoshop.svg" width="48">       |
|        `psc`        |  <img src="./assets/photoshopclassic.svg" width="48">    |
|        `psx`        |  <img src="./assets/photoshopexpress.svg" width="48">    |
|        `php`        |       <img src="./assets/php-dark.svg" width="48">       |
|      `phpstorm`     |    <img src="./assets/phpstorm-dark.svg" width="48">     |
|       `pinia`       |      <img src="./assets/pinia-dark.svg" width="48">      |
|        `pkl`        |       <img src="./assets/pkl-dark.svg" width="48">       |
|       `plan9`       |      <img src="./assets/plan9-dark.svg" width="48">      |
|    `planetscale`    |   <img src="./assets/planetscale-dark.svg" width="48">   |
|        `pnpm`       |      <img src="./assets/pnpm-dark.svg" width="48">       |
|     `pocketbase`    |   <img src="./assets/pocketbase-dark.svg" width="48">    |
|         `pf`        |       <img src="./assets/portfolio.svg" width="48">      |
|      `postgres`     |   <img src="./assets/postgresql-dark.svg" width="48">    |
|      `postman`      |       <img src="./assets/postman.svg" width="48">        |
|     `powerpoint`    |   <img src="./assets/powerpoint-dark.svg" width="48">    |
|     `powershell`    |   <img src="./assets/powershell-dark.svg" width="48">    |
|       `preact`      |       <img src="./assets/preact-dark.svg" width="48">    |
|         `pl`        |       <img src="./assets/prelude.svg" width="48">        |
|         `pr`        |       <img src="./assets/premiere.svg" width="48">       |
|         `ru`        |     <img src="./assets/premiererush.svg" width="48">     |
|       `prisma`      |        <img src="./assets/prisma.svg" width="48">        |
|     `processing`    |   <img src="./assets/processing-dark.svg" width="48">    |
|     `prometheus`    |      <img src="./assets/prometheus.svg" width="48">      |
|        `pug`        |       <img src="./assets/pug-dark.svg" width="48">       |
|      `pupeteer`     |     <img src="./assets/pupeteer-dark.svg" width="48">    |
|      `pycharm`      |     <img src="./assets/pycharm-dark.svg" width="48">     |
|         `py`        |     <img src="./assets/python-dark.svg" width="48">      |
|      `pytorch`      |     <img src="./assets/pytorch-dark.svg" width="48">     |
|         `qt`        |       <img src="./assets/qt-dark.svg" width="48">        |
|         `r`         |        <img src="./assets/r-dark.svg" width="48">        |
|      `rabbitmq`     |    <img src="./assets/rabbitmq-dark.svg" width="48">     |
|       `rails`       |        <img src="./assets/rails.svg" width="48">         |
|    `raspberrypi`    |   <img src="./assets/raspberrypi-dark.svg" width="48">   |
|       `react`       |      <img src="./assets/react-dark.svg" width="48">      |
|     `reactivex`     |    <img src="./assets/reactivex-dark.svg" width="48">    |
|       `redhat`      |     <img src="./assets/redhat-dark.svg" width="48">      |
|       `redis`       |      <img src="./assets/redis-dark.svg" width="48">      |
|       `redux`       |        <img src="./assets/redux.svg" width="48">         |
|       `regex`       |      <img src="./assets/regex-dark.svg" width="48">      |
|       `remix`       |      <img src="./assets/remix-dark.svg" width="48">      |
|       `replit`      |     <img src="./assets/replit-dark.svg" width="48">      |
|       `rider`       |      <img src="./assets/rider-dark.svg" width="48">      |
|    `robloxstudio`   |     <img src="./assets/robloxstudio.svg" width="48">     |
|       `rocket`      |        <img src="./assets/rocket.svg" width="48">        |
|      `rollupjs`     |    <img src="./assets/rollupjs-dark.svg" width="48">     |
|        `ros`        |       <img src="./assets/ros-dark.svg" width="48">       |
|        `ruby`       |         <img src="./assets/ruby.svg" width="48">         |
|        `rust`       |         <img src="./assets/rust-dark.svg" width="48">    |
|      `safari`       |       <img src="./assets/safari-dark.svg" width="48">    |
|        `sass`       |         <img src="./assets/sass.svg" width="48">         |
|       `spring`      |     <img src="./assets/spring-dark.svg" width="48">      |
|       `sqlite`      |        <img src="./assets/sqlite.svg" width="48">        |
|      `sqlserver`    |      <img src="./assets/sqlserver-dark.svg" width="48">  |
|   `stackoverflow`   |  <img src="./assets/stackoverflow-dark.svg" width="48">  |
|       `stock`       |         <img src="./assets/stock.svg" width="48">        |
|  `styledcomponents` |   <img src="./assets/styledcomponents.svg" width="48">   |
|      `sublime`      |     <img src="./assets/sublime-dark.svg" width="48">     |
|      `supabase`     |    <img src="./assets/supabase-dark.svg" width="48">     |
|     `surrealdb`     |   <img src="./assets/surrealdb-dark.svg" width="48">     |
|       `scala`       |      <img src="./assets/scala-dark.svg" width="48">      |
|        `scipy`      |       <img src="./assets/scipy-dark.svg" width="48">     |
|      `scratch`      |       <img src="./assets/scratch.svg" width="48">        |
|      `sklearn`      |   <img src="./assets/scikitlearn-dark.svg" width="48">   |
|      `selenium`     |       <img src="./assets/selenium.svg" width="48">       |
|       `sentry`      |        <img src="./assets/sentry.svg" width="48">        |
|     `sequelize`     |    <img src="./assets/sequelize-dark.svg" width="48">    |
|    `sharepoint`     |    <img src="./assets/sharepoint-dark.svg" width="48">   |
|      `shopify`      |       <img src="./assets/shopify-dark.svg" width="48">   |
|      `sketchup`     |    <img src="./assets/sketchup-dark.svg" width="48">     |
|      `solidity`     |       <img src="./assets/solidity.svg" width="48">       |
|      `solidjs`      |     <img src="./assets/solidjs-dark.svg" width="48">     |
|       `spark`       |         <img src="./assets/spark.svg" width="48">        |
|       `svelte`      |        <img src="./assets/svelte.svg" width="48">        |
|        `svg`        |       <img src="./assets/svg-dark.svg" width="48">       |
|       `swift`       |        <img src="./assets/swift.svg" width="48">         |
|      `symfony`      |     <img src="./assets/symfony-dark.svg" width="48">     |
|      `tableau`      |      <img src="./assets/tableau-dark.svg" width="48">    |
|      `tailwind`     |   <img src="./assets/tailwindcss-dark.svg" width="48">   |
|       `tauri`       |      <img src="./assets/tauri-dark.svg" width="48">      |
|       `teams`       |      <img src="./assets/teams-dark.svg" width="48">      |
|     `tensorflow`    |   <img src="./assets/tensorflow-dark.svg" width="48">    |
|     `terraform`     |    <img src="./assets/terraform-dark.svg" width="48">    |
|      `threejs`      |     <img src="./assets/threejs-dark.svg" width="48">     |
|        `tor`        |       <img src="./assets/tor-dark.svg" width="48">       |
|      `twitter`      |       <img src="./assets/twitter.svg" width="48">        |
|         `ts`        |      <img src="./assets/typescript.svg" width="48">      |
|       `ubuntu`      |     <img src="./assets/ubuntu-dark.svg" width="48">      |
|       `unity`       |      <img src="./assets/unity-dark.svg" width="48">      |
|       `unreal`      |     <img src="./assets/unrealengine.svg" width="48">     |
|         `v`         |        <img src="./assets/v-dark.svg" width="48">        |
|        `vala`       |         <img src="./assets/vala.svg" width="48">         |
|         `vb`        |   <img src="./assets/visualbasic-dark.svg" width="48">   |
|       `vercel`      |     <img src="./assets/vercel-dark.svg" width="48">      |
|        `vim`        |       <img src="./assets/vim-dark.svg" width="48">       |
|    `visualstudio`   |  <img src="./assets/visualstudio-dark.svg" width="48">   |
|        `vite`       |      <img src="./assets/vite-dark.svg" width="48">       |
|       `vitest`      |     <img src="./assets/vitest-dark.svg" width="48">      |
|       `vscode`      |     <img src="./assets/vscode-dark.svg" width="48">      |
|      `vscodium`     |    <img src="./assets/vscodium-dark.svg" width="48">     |
|        `vue`        |      <img src="./assets/vuejs-dark.svg" width="48">      |
|      `vuetify`      |     <img src="./assets/vuetify-dark.svg" width="48">     |
|        `wasm`       |     <img src="./assets/webassembly.svg" width="48">      |
|      `webflow`      |       <img src="./assets/webflow.svg" width="48">        |
|      `webpack`      |     <img src="./assets/webpack-dark.svg" width="48">     |
|      `webstorm`     |    <img src="./assets/webstorm-dark.svg" width="48">     |
|      `windicss`     |    <img src="./assets/windicss-dark.svg" width="48">     |
|      `windows`      |     <img src="./assets/windows-dark.svg" width="48">     |
|        `word`       |       <img src="./assets/word-dark.svg" width="48">      |
|     `wordpress`     |      <img src="./assets/wordpress.svg" width="48">       |
|      `workers`      |     <img src="./assets/workers-dark.svg" width="48">     |
|       `xcode`       |       <img src="./assets/xcode-dark.svg" width="48">     |
|         `xd`        |          <img src="./assets/xd.svg" width="48">          |
|       `yammer`      |     <img src="./assets/yammer-dark.svg" width="48">      |
|        `yarn`       |      <img src="./assets/yarn-dark.svg" width="48">       |
|        `yew`        |       <img src="./assets/yew-dark.svg" width="48">       |
|        `zed`        |       <img src="./assets/zed-dark.svg" width="48">       |
|        `zig`        |       <img src="./assets/zig-dark.svg" width="48">       |

---

## 💖 Support the Project

Thank you so much already for using my projects!

To support the project directly, feel free to open issues for icon suggestions, or contribute with a pull request!
