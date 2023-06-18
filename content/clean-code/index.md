+++
title = "How I Code Clean Code"
#date = "2023-03-20T19:18:19Z"
#author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
#cover = "/photo.jpg"
#tags = ["sandro-lain"]
keywords = ["sandro", "lain", "curriculum"]
description = "About"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

## Name and structure definition

In a data structure it is good that each value belongs to the property of an object, therefore it is better to avoid:

- values whose meaning depends on the position in an array
- values as keys for a map or dictionary

A pattern for defining the name of the variables can be a composition in sequence of:

- Adjective
- Subject
- Type

Example:

```js
const promotedUsersArr = [];
const enabledViewBool = false;
const badBoysNum = NaN;
```

The names of the methods should instead start with verbs (eg. `get`, `set`, `write`, `send`, …).

## Coding

```js
/**
 * TODO::
 * • .editorconfig
 * • .gitignore
*/
```

## Linting

```js
/**
 * TODO:
 * • Eslint
 * • Autoformat
 */
```

## Testing

```js
// TODO
```
