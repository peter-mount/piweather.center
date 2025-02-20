BUILD_DIR=../builds

.PHONY: all clean pdf

all: pdf

clean: cleanpart
	@-echo "CLEAN TeX  pdf"; rm -f book.pdf

cleanpart:
	@-echo "CLEAN TeX  work"; rm -f book.[a-oq-su-z]* book.toc

pdf: book.pdf

book.idx book.nlo:
	@-echo "TeX build  idx nlo";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true

book.nls: book.idx book.nlo
	@-echo "TeX nomenc book";\
	makeindex book.idx book.nlo -s nomencl.ist -o book.nls
	@-echo "TeX build  book.tex";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true

book.bcf:
	@-echo "TeX build  bcf";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true

# Bibliography
book.bbl: book.bcf main.bib
	@-echo "TeX biber  book";\
	biber book 2>&1 >/dev/null
	@-echo "TeX build  book.tex";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true

## Glossary
book.glo: book.tex
	@-echo "TeX gloss  book.tex";\
	makeglossaries book

book.pdf: book.tex book.bbl book.glo book.nls
	@-echo "TeX PDF    book.tex";\
	pdflatex -interaction=nonstopmode book.tex >/dev/null 2>&1 || true
