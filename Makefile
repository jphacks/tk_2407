.PHONY: gen
gen:
	cd frontend && pnpm run api:build
	cd backend && make gen
