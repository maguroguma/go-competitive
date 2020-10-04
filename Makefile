.PHONY: testlib
testlib:
	go test ./lib/*

.PHONY: testlib-fully
testlib-fully:
	go test -v -cover ./lib/*
