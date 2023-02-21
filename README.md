# Luadoc project

Lua reference documentation in more readable format.

## Tools

Convert manual to markdown:

```
pandoc -s ./doc/manual.html -o docs/manual.md
```


Split manual by headers:

```
cd docs

csplit manual.md -f 'ch' --suffix-format='%02d.md' "/^## /" "{*}" -s
```
