Command json2yaml converts JSON documents to YAML

If you have Go installed, install this command with:

    go install github.com/artyom/json2yaml@latest

If command is invoked as `yaml2json`, it converts in the opposite direction — from YAML to JSON.
If your Go binaries are installed under the default $GOPATH/bin, then you can create a symlink like this:

    ln -s -v json2yaml $(go env GOPATH)/bin/yaml2json
