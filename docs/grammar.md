```mermaid
graph LR

prog --> stmt*
stmt* --> exit(["exit([expr]);"])
stmt* --> let(["let ident = [expr];"])

expr["[expr]"]
expr --> int_lit
expr --> ident
```
