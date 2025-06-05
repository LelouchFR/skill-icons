<p align="center"><img align="center" width="280" src="./.github/text-logo.svg#gh-dark-mode-only"/></p>
<p align="center"><img align="center" width="280" src="./.github/text-logo-light.svg#gh-light-mode-only"/></p>
<h3 align="center">Présentez vos compétences sur votre GitHub ou CV en toute simplicité!</h3>
<hr>

# Docs

- [Docs](#docs)
- [Example](#example)
- [Spécifier Les Icônes](#specifier-les-icones)
- [Utiliser Les Thèmes d'Icônes](#utiliser-les-themes-d'icones)
- [Icônes Par Ligne](#icon-par-ligne)
- [Titre Des Icônes](#titre-des-icones)
- [Centrer Des Icônes](#centrer-des-icones)
- [Liste Des Icônes](#liste-des-icones)
- [💖 Soutenir Le Projet](#-soutenir-le-projet)

# Example

<p align="center"><img align="center" src="./.github/example-dark.png#gh-dark-mode-only"/></p>
<p align="center"><img align="center" src="./.github/example-light.png#gh-light-mode-only"/></p>

# Spécifier Les Icônes

Copiez et collez le bloc de code ci-dessous dans votre README pour ajouter des icônes !

Remplacez le `?i=js,html,css,wasm` par une liste de vos compétences séparées par des virgules "," ! Vous trouverez la liste complète des icônes [ici] (#liste-des-icones).

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=js,html,css,wasm)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=js,html,css,wasm)

# Utiliser Les Thèmes d'Icônes

Certaines icônes ont un arrière-plan sombre et un arrière-plan clair. Vous pouvez spécifier le thème que vous souhaitez en tant que paramètre de l'url.

Cette option est facultative. Le thème par défaut est spécifique au système.

Changez le `&theme=light` en `dark` ou `light`. Le thème est la couleur d'arrière-plan, donc le thème light a un arrière-plan d'icône blanc, et dark a un arrière-plan noir.

**Example de thème clair:**

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=java,kotlin,nodejs,figma&theme=light)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=java,kotlin,nodejs,figma&theme=light)

# Icônes Par Ligne

Vous pouvez spécifier le nombre d'icônes que vous souhaitez par ligne ! C'est un argument facultatif, et la valeur par défaut est 15.

Remplacez `&perline=3` par un nombre compris entre 1 et 50.

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=aws,gcp,azure,react,vue,flutter&perline=3)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=aws,gcp,azure,react,vue,flutter&perline=3)

# Titre Des Icônes

Vous pouvez avoir la possibilité d'ajouter le nom des icônes que vous mettez pour aider ceux qui ne savent pas ce que c'est en utilisant `&titles`.

La valeur de `titles` est un booléen, elle doit donc être `true` ou `false`, la valeur par défaut est `false`

```md
![My Skills](https://go-skill-icons.vercel.app/api/icons?i=rust,surrealdb,actix,yew&titles=true)
```

![My Skills](https://go-skill-icons.vercel.app/api/icons?i=rust,surrealdb,actix,yew&titles=true)

# Centrer Des Icônes

Vous voulez centrer les icônes dans votre README ? Les SVG sont automatiquement redimensionnés, vous pouvez donc le faire de la même manière que vous centreriez une image.

```html
<p align="center">
  <a href="https://go-skill-icons.vercel.app/">
    <img
      src="https://go-skill-icons.vercel.app/api/icons?i=git,kubernetes,docker,c,vim"
    />
  </a>
</p>
```

<p align="center">
  <a href="https://go-skill-icons.vercel.app/">
    <img src="https://go-skill-icons.vercel.app/api/icons?i=git,kubernetes,docker,c,vim" />
  </a>
</p>

# Liste Des Icônes

# 💖 Soutenir Le Projet

Merci beaucoup d'avoir utilisé mes projets !

Pour soutenir le projet directement, n'hésitez pas à ouvrir des issues pour des suggestions d'icônes, ou à contribuer avec une Pull Request!
