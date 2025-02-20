BUILD_DIR=../builds

.PHONY: all clean pdf

all: pdf

clean: cleanpart
	@-echo "CLEAN TeX  pdf"; rm -f book.pdf

cleanpart:
	@-echo "CLEAN TeX  work"; rm -f book.[a-oq-su-z]* book.toc

pdf: book.pdf

book.pdf: cleanpart book.tex
	@-echo "TeX BUILD  book.tex";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true
	@-echo "TeX biber  book";\
	biber book 2>&1 >/dev/null
	@-echo "TeX PDF    book.tex";\
	pdflatex -interaction=nonstopmode book.tex >/dev/null 2>&1 || true
