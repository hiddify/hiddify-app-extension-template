#!/bin/bash
base_gomod=$(curl -sL https://raw.githubusercontent.com/hiddify/hiddify-core/refs/heads/main/go.mod)

sed -i '/^replace /d' go.mod

replace_lines=$(echo "$base_gomod" | grep '^replace ')

if [ -n "$replace_lines" ]; then
    echo "$replace_lines" >> go.mod
    echo "Replace directives appended to go.mod"
else
    echo "No replace directives found in the remote go.mod"
fi

latest_commit_hash=$(git ls-remote https://github.com/hiddify/hiddify-core.git refs/heads/main | awk '{print $1}')
sed -i "s|github.com/hiddify/hiddify-core [^ ]*|github.com/hiddify/hiddify-core $latest_commit_hash|g" go.mod
go mod tidy

mkdir -p extension/html
curl -L -o extension/html/index.html https://raw.githubusercontent.com/hiddify/hiddify-core/refs/heads/main/extension/html/index.html
curl -L -o extension/html/rpc.js https://raw.githubusercontent.com/hiddify/hiddify-core/refs/heads/main/extension/html/rpc.js

