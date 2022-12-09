+++
title = "CSS variable as a responsive selector"
date = "2020-07-05T18:00:00Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/pankaj-patel-6JVlSdgMacE-unsplash.jpg"
tags = ["css variables"]
keywords = ["css variables"]
description = ""
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

Photo by <a href="https://unsplash.com/@pankajpatel?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Pankaj Patel</a> on <a href="https://unsplash.com/s/photos/css?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Unsplash</a>

# CSS variable as a responsive selector

A few days ago, while doing some tests of the use of CSS variables, I found a solution to apply dynamic styles through the use of only **CSS variables**, similar to what is possible with some CSS frameworks that provide **different values** ​​of the same class (margin, padding, border, etc ...).

With this trick, it is possible to **reduce the redundancy** of the selectors and cancel their limits since the possible values ​​are not only those foreseen by the but **virtually infinite**.

The solution also provides that the style is **applied only to elements that explicitly have the CSS variable** in the style attribute.
To be precise, look for the elements that have indicated a CSS variable that ends with the one in question (note the less hyphen at the beginning into the example below).
This will allow us to add prefixes, which we will see later.

```css
[style*="-margin:"] { /*...*/ }
```

The variable we will use will not be the same as the selector, but one that ends differently ( ex. "-val" suffix ).

```css
[style*="-margin:"] {
  margin: var(--margin-val);
}
```

A second selector will search for the exact CSS variable and copy its value into the variable used by the applied properties.

```css
[style*="--margin:"] {
  --margin-val: var(--margin);
}
```

This is to allow you to **create variations** of the value that depend on the prefix we said earlier.
An example is a possibility of having different prefixes for different `@media`, allowing us to make the **CSS variable responsive**.

```css
@media screen and (max-width: 640px) {
  [style*="--phone-margin:"] {
    --margin-val: var(--phone-margin);
  }
}
```


## Pros
This technique allows you with a few lines to have the same functionality as multiple specific classes for each value.
In the example code a margin value is applied, but thanks to the CSS variables and the `calc()` function we can define complex and responsive styles, such as the offset and blur values ​​of a shadow.

## Cons
The selector for substring value of an attribute has **lower performance** than that of a CSS class, and the use of CSS variables and the `calc()` function can also be a performance overhead.

## Final code

Example CSS source:
```css
[style*="-margin:"] {
  margin: var(--margin-val);
}
[style*="--margin:"] {
  --margin-val: var(--margin);
}
@media screen and (max-width: 640px) {
  [style*="--phone-margin:"] {
    --margin-val: var(--phone-margin);
  }
}
@media screen and (max-width: 1024px) {
  [style*="--tablet-margin:"] {
    --margin-val: var(--tablet-margin);
  }
}
```

Example HTML source:
```html
<div style="
  --margin: 100px;
  --tablet-margin: 50px;
  --phone-margin: 25px;
  border: 1px solid #000000;
  ">Hello! Resize the viewport to switch between styles</div>
```

<Link href="/001-example.html"><a>Live Example</a></Link>
