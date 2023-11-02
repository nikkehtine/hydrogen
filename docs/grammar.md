```mermaid
graph LR

prog --> stmt*
stmt* --> exit(["<i>exit</i>([expr])"])
stmt* --> let(["<i>let</i> ident = [expr]"])

expr["[expr]"]
expr --> int_lit
expr --> ident
```
