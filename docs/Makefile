BUILD_DIR=../builds

.PHONY: all clean pdf

all: pdf

clean: cleanpart
	@-echo "CLEAN TeX  pdf"; rm -f book.pdf

cleanpart:
	@-echo "CLEAN TeX  work"; rm -f book.[a-oq-su-z]* book.toc

pdf: book.pdf

book.pdf: cleanpart book.tex
	@-echo "TeX build  book.tex";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true
	@-echo "TeX nomenc book";\
	makeindex book.idx book.nlo -s nomencl.ist -o book.nls
	@-echo "TeX build  book.tex";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true
	@-echo "TeX biber  book";\
	biber book 2>&1 >/dev/null
	@-echo "TeX build  book.tex";\
	pdflatex -interaction=nonstopmode -draftmode book.tex >/dev/null 2>&1 || true
	@-echo "TeX gloss  book.tex";\
	makeglossaries book
	@-echo "TeX PDF    book.tex";\
	pdflatex -interaction=nonstopmode book.tex >/dev/null 2>&1 || true
