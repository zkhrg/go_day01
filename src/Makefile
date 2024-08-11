READDB_TARET=./bin/readDB
COMPAREDB_TARGET=./bin/compareDB
COMPAREFS_TARGET=./bin/compareFS

all: clean $(READDB_TARET) $(COMPAREDB_TARGET) $(COMPAREFS_TARGET)

$(READDB_TARET):
	go build -o $(READDB_TARET) cmd/cli_read_db/main.go

$(COMPAREDB_TARGET):
	go build -o $(COMPAREDB_TARGET) cmd/cli_compare_db/main.go

$(COMPAREFS_TARGET):
	go build -o $(COMPAREFS_TARGET) cmd/cli_compare_fs/main.go

test_cmp:
#	@./bin/compareDB --old ../datasets/recipes001.json --new ../datasets/recipes001.xml
#	@./bin/compareDB --old ../datasets/recipes002.json --new ../datasets/recipes002.xml
	@./bin/compareDB --old ../datasets/recipes002.xml --new ../datasets/recipes002.json

test_fs: 
	@./bin/compareFS --old ../datasets/snapshot3.txt --new ../datasets/snapshot4.txt

clean:
	@echo "files removed successfully: \n$(READDB_TARET)\n$(COMPAREDB_TARGET)\n$(COMPAREFS_TARGET)"
	@rm -f $(READDB_TARET) $(COMPAREDB_TARGET) $(COMPAREFS_TARGET)