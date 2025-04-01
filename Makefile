BUILD = bin
EXE = shell
SRC = .
LDFLAGS = -ldflags="-s -w"

all: windows macos linux 
	echo "done."

windows:
	mkdir -p $(BUILD)
	GOOS=windows go build -o $(BUILD)/$(EXE)_win.exe $(LDFLAGS) $(SRC)

macos:
	mkdir -p $(BUILD)
	GOOS=darwin go build -o $(BUILD)/$(EXE)_macos $(LDFLAGS) $(SRC)

linux:
	mkdir -p $(BUILD)
	GOOS=linux go build -o $(BUILD)/$(EXE)_linux $(LDFLAGS) $(SRC)

clean:
	rm -rf $(BUILD)
