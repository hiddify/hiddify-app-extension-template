#!/usr/bin/env bash
while getopts a:n:u:d: flag
do
    case "${flag}" in
        a) author=${OPTARG};;
        n) name=${OPTARG};;
        u) urlname=${OPTARG};;
        d) description=${OPTARG};;
    esac
done

echo "Author: $author";
echo "Project Name: $name";
echo "Project URL name: $urlname";
echo "Description: $description";

echo "Renaming project..."

original_author="author_name"
original_name="hiddify-app-extension-template"
original_urlname="project_urlname"
original_description="project_description"

orginal_example_name="ExampleExtension"

newclass=$(echo "${urlname}" | awk -F'[-_]' '{ for(i=1; i<=NF; i++) { printf "%s", toupper(substr($i, 1, 1)) substr($i, 2) } }')

# for filename in $(find . -name "*.*") 
for filename in $(git ls-files) 
do
    sed -i "s/$original_author/$author/g" $filename
    sed -i "s/$original_name/$name/g" $filename
    sed -i "s/$original_urlname/$urlname/g" $filename
    sed -i "s/$original_description/$description/g" $filename
    sed -i "s/$orginal_example_name/$newclass/g" $filename
    
    echo "Renamed $filename"
done

