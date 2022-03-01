PID_FILE = /tmp/go.pid

build:
	go build -o bin/app main.go && clear

kill:
	-kill `pstree -p \`cat $(PID_FILE)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"` 

run:   
	bin/app & echo $$! > $(PID_FILE)

watch: build run
	fswatch -or --event=Updated  ./app ./main.go | \
	xargs -n1 -I {} make reload

reload: kill build run

# .PHONY is used for reserving tasks words
.PHONY: build kill reload watch

# supress echo commands on cli
.SILENT: build kill run watch reload