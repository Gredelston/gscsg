**GSCSG** ([IPA](https://en.wiktionary.org/wiki/Appendix:English_pronunciation): /gʌs'kʌsgʌ/) is **G**reg's **S**tatic **C**ontent **S**ite **G**enerator.

# Background

Websites are cool! [I have one](https://www.gredelston.com). At the time of this writing, my website is made with [Gatsby](https://www.gatsbyjs.org), a static-site generator for Node.js.

The only problem is, I suck at Node.js and I don't have time to learn someone else's system! Surely, it's easier to [make my own system](https://en.wikipedia.org/wiki/Not_invented_here).

# Goals

* Make it trivial to extend my own website with new content
* Don't have terrible code that I will hate in six weeks/months/years

## Reach goals

* Spin this out into a package that other people can use for their own websites... lol yeah right

# Design

Needs fleshing out, but here's the gist:

* `site.cfg` will contain information like which page is the homepage, and probably other stuff too.
* We'll use [Mustache](https://github.com/cbroglie/mustache) templates to make modular pages/components. Example: blog-post page will probably inherit from the standard "page" template, which will probably contain a "sidebar" template and a "background" template... modularity is cool!
* I'll also define some content-types. Example: blog post, which will contain fields like "title", "short-blurb", "date", "tags", and "content".
* Certain templates will generate webpages based on content-types. Example: each blog post will generate a webpage; and also, each unique tag will generate a webpage which links to all the blog posts.

That's enough detail for now, but I'd love to spin this out into a whole-ass design doc.
