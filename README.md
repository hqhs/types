Types -- template for optional types suitable for any project
===

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Types -- template for optional types suitable for any project](#types----template-for-optional-types-suitable-for-any-project)
    - [Motivation](#motivation)
    - [Usage](#usage)
    - [Examples](#examples)
    - [Code generation](#code-generation)

<!-- markdown-toc end -->


## Motivation

Sometimes input for your code could have optional optional values. It's
definetly possible to build entire codebase using default zero-values for each
type. Google does that, (TODO proof link?) and since golang enforces you to
write code in google's way, it does exacly what then encoding/decoding data to
various formats  (TODO proof links). But different companies has different
approaches and not everyone has 20 years of expirience of building systems that
scales well, and in that scenario default zero value could break old
compatability. There're two major ways to do deal with what:

- Using pointers:

``` go
type Data struct {
    id int
    name *string
    description *string
}
```

I'm have comprehensive explanation why it's terrible here (TODO blog post link).
TL;DR: It's classique example of [billion dollar
mistake](https://en.wikipedia.org/w/index.php?title=Tony_Hoare&action=edit&section=3&editintro=Template:BLP_editintro),
which breaks thousands of production services. One single missed check from
programmer and your runtime crashed because of NillPointerDereference. 

- Using optional fields:

``` go
type Data struct {
    id int
    name String
    description String
}

type String struct {
    V string
    Set bool
}
```

Much better & simpler approach. But because of golang verboseness it requires
you to write a lot of boilerplate code and could slow speed of work. This repo
provides minimalistic template for code generation which solves optional fields
downsides and easily extandable.

## Usage

NOTE what originally this repo is hard fork of leighmcculloch's
[go-optional](https://github.com/leighmcculloch/go-optional), but I use other
implementation of optional types and sugger different use scenario.

## Examples

## Code generation

