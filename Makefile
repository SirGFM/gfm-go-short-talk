
all: out/go-security-i-guess.pdf

.PHONY: images all

# Recursive wildcard from https://stackoverflow.com/a/18258352
#
# To find every .c file in src:
#   FILES := $(call rwildcard, ,src/ *.c)
# To find all the .c and .h files in src:
#   FILES := $(call rwildcard, src/, *.c *.h)
rwildcard=$(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

DEPS := $(call rwildcard, slides/, *.tex)
DEPS := $(DEPS) $(call rwildcard, imgs/, *.png)

# Regenerate the slides if any file changes
#
# Since pdflatex creates the output from the main input file
# (e.g., main.tex -> main.pdf), replace .pdf suffix in the output
# by .tex in the input, also renaming the input directory.
#
# As dumb as this seems, the toc only appears if pdflatex
# is run twice, and (afaik) there's no way to rename the file generated
# by pdflatex...
out/go-security-i-guess.pdf: slides/main.tex $(DEPS) | out/.mkdir
	pdflatex -halt-on-error -output-directory=out $<
	pdflatex -halt-on-error -output-directory=out $<
	mv $(<:slides/%.tex=out/%.pdf) $@

images: out/img/.timestamp

out/img/.timestamp: out/go-security-i-guess.pdf | out/img/.mkdir
	pdftoppm -png -r 210 -scale-to 1280 -H 720 $< $(@D)/tmp
	bash -c 'i=0; for f in `ls -1 $(@D)/tmp*`; do mv $${f} $(@D)/img-$${i}.png; let "i = i + 1"; done'
	@ touch out/img/.timestamp
	@ echo "Images generated at $(@D)"

# Automatically creates a directory.
%.mkdir:
	mkdir -p $(@D)
	touch $@

clean:
	rm -rf out/
