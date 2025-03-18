.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ..\..\idl  --type HTTP --service frontend --module https://github.com/somewan820/tiktok/app/frontend --idl ../../idl/frontend/home.proto