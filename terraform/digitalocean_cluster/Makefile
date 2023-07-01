all: plan

FLAGS=
TMPDIR=tmp
TMPFILENAME=plan.out
TMPFILE=$(TMPDIR)/$(TMPFILENAME)

tmpdir:
	@mkdir -p $(TMPDIR)

init:
	terraform init

$(TMPFILE): tmpdir plan

plan: init tmpdir
	terraform plan $(FLAGS) -out $(TMPFILE)

apply: init tmpdir $(TMPFILE)
	terraform apply $(FLAGS) $(TMPFILE)

destroy: init
	terraform destroy $(FLAGS)

force-%:
	make $(subst force-,,$@) FLAGS=--auto-approve

.PHONY: all plan apply destroy init
