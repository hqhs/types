Types: template for optional types suitable for any project
===

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Types: template for optional types suitable for any project](#types-template-for-optional-types-suitable-for-any-project)
    - [Motivation](#motivation)
    - [Usage](#usage)
    - [How to add your own methods for optional types](#how-to-add-your-own-methods-for-optional-types)
    - [How does it work](#how-does-it-work)

<!-- markdown-toc end -->


## Motivation

Sometimes input for your code could have optional optional values. It's
definetly possible to build entire codebase using default zero-values for each
type. I believe google does that, but couldn't find proof link, kind of guessing
based on how optional fields implemented in gRPC and Go, and since golang enforces you to
write code in such way, it does exacly [what](https://golang.org/pkg/encoding/json/#Marshal) then encoding/decoding data to
various formats. But different companies has different
approaches and not everyone has 20 years of expirience of building systems that
scales well, and in that scenario default zero value could break old
compatability. There're two major ways to do deal with what:

- Using pointers:

``` go
type Data struct {
    id int `json:"id"`
    name *string `json:"name"`
    description *string `josn:"description"`
}
```

I'm have comprehensive explanation why it's terrible here (TODO blog post link).

TL;DR: It's classique example of [billion dollar
mistake](https://en.wikipedia.org/w/index.php?title=Tony_Hoare&action=edit&section=3&editintro=Template:BLP_editintro),
which had broke thousands of production services. One single missed check from
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

``` sh
cd myprojectname/pkg # assuming you follow golang [standard project layout](https://github.com/golang-standards/project-layout)
git clone https://github.com/hqhs/types.git && cd types 
rm -rf .git
IMPORT="$(golang_import_path_for_project)/pkg/types" make replace # replace import path 
make generate
# now you're ready to use types in your project!
```

``` go
package main

import (
    "bytes"
    "fmt"
    "encoding/json"

    "example.com/project/name/pkg/types" // let golang_import_path_for_project be "example.com/project/name"
) 

type Data struct {
    id types.Int
    name types.String
}

func main() {
    d := Data{types.OfInt(3)}
    b := &bytes.Buffer{}
    err := json.NewEncoder(b).Encode(d)
    if err != nil {
        panic(err)
    }
    fmt.Println("encoded struct: %s", b.String())
}
```

## How to add your own methods for optional types


## How does it work




