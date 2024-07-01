.PHONY: code database test/smoke test/race test/integration test/coverage image/build image/latest

code:
	@"$(CURDIR)/scripts/code.sh"

database:
	@"$(CURDIR)/scripts/database.sh"

test/smoke:
	@"$(CURDIR)/scripts/test.sh" smoke

test/race:
	@"$(CURDIR)/scripts/test.sh" race

test/coverage:
	@"$(CURDIR)/scripts/test.sh" coverage
