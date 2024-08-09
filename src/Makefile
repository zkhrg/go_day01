READDB_TARET=./bin/readDB
COMPAREDB_TARGET=./bin/compareDB
COMPAREFS_TARGET=./bin/compareFS

all: $(READDB_TARET)

$(READDB_TARET):
	go build -o $(READDB_TARET) cmd/cli_read_db/main.go 

clean:
	rm -f $(READDB_TARET)