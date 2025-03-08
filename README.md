# go-markdown

go-markdown is a library to parse markdown into HTML

## add to your project

```
go get github.com/dearing/go-markdown
```

## install cli (todo)

```
go install github.com/dearing/go-markdown/cmd/
```

## markdown support

- [ ] paragraph
- [x] heading 1
- [x] heading 2
- [x] heading 3
- [x] heading 4
- [x] heading 5
- [x] heading 6
- [x] horizontal rule
- [ ] inline code span
- [ ] block preformatted
- [ ] unsorted list and list items
- [ ] sorted list and list items

## development notes

- not attempting to be 100% compliant with [commonmark](https://commonmark.org/) and others, just enough for my content needs
- Rob Pike has a [good talk](https://www.youtube.com/watch?v=HxaD_trXwRE) about using a state machine lexical scanning in relation to templates that this project takes inspiration from. 
- the root of this project will remain a library and the `cmd/` folder will contain a CLI tool