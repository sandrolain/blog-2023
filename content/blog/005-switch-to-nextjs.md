+++
title = "Switching from mdx-go to next.js"
date = "2021-05-15T11:14:04Z"
author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
cover = "/cover/sarah-kilian-52jRtc2S_VE-unsplash.jpg"
tags = []
keywords = []
description = "Little drama for my little blog"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

*Photo by [Sarah Kilian](https://unsplash.com/@rojekilian?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText) on [Unsplash](https://unsplash.com/s/photos/broken?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText)*

When I decided to start writing this blog, I wanted to write it using **[MDX](https://mdxjs.com/)** as the post format because I am a fan of Markdown (certainly less laborious than writing HTML directly) and because it would allow me to add dynamic elements written in react.js.

To quickly set up the blog I looked around what were the (not many) platforms ready to create static sites using MDX.
Among all the possible ones I found **[mdx-go](https://jxnblk.github.io/mdx-go/)**, a semi-ready solution for writing sites using MDX.

It was easy to set up the first version of the site and with minimal effort, I had everything I needed.
Over time, however, I noticed that not everything was perfect: the resolution of the paths not always correct, the impossibility of making a dynamic listing, in addition to the fact that the framework seems not to be updated for a while.

Recently GitHub via **dependabot** intervened with a security update PR on a dependency, which *"broke"* mdx-go. The pages were available but the site appeared without style and basic structure.

Thanks to the experience gained with my other blog ([Arché](https://arche.sandrolain.com/)), I decided to migrate to **[next.js](https://nextjs.org/)**, which has official plugins for MDX.
In just over an hour I was able to convert from the old to the new framework, and now I can also take advantage of the next.js integration on [Netlify](https://www.netlify.com/), where this site is hosted.

**Lesson learned**: Sometimes it is better to do a little more work and use a better known and supported platform than to have things ready and then find yourself limited or with many problems.
