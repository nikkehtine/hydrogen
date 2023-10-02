```mermaid
graph LR
E["prog"] --> S(("stmt*"))

S --> exit(["<i>exit</i>([expr])"])
S --> let(["<i>let</i> ident"])

A["[exit]"] --> B(["exit([expr])"])
C["[expr]"] --> D(["int_lit"])
```
