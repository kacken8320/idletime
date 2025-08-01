# Idletime Docs

We only compile locally to avoid bloat, you could've also included it in Docker.

## PlantUML & Java installation local on Debian
Run: `sudo apt install plantuml default-jre -y`
Check: `plantuml -version`

## Compile locally to SVG
Compile: `plantuml -tsvg diagrams/*.puml`
Alternative flags: use `-tpng` or `-tpdf` instead of `-tsvg` for PNG or PDF, respectively.
