#!/bin/sh

# File to be generated.
# It will be regenerated if already exists.
filename=deserializers.go

if [ -e $filename ]
then 
    # Deleting file before applying code generation.
    rm $filename 
fi

# Only the first run of the loop keep header.
# All the following remove header until import 
# so the resulting file has a single header.
keep_header_until_imports=true

# Camel Case (first letter upper case)
for Prototype in $@
do
    # Upper Case
    PROTOTYPE=$(echo "${Prototype}" | sed 's/[a-z]/\U&/g')

    # Lower Case
    prototype=$(echo "${Prototype}" | sed 's/[A-Z]/\L&/g')

    if [ $keep_header_until_imports == true ]
    then
        cat deserializers.tmpl.go |
        sed '1,2d' | # removes go:build ignore from template
        sed -e 's/Prototype/'${Prototype}'/g' | 
        sed -e 's/PROTOTYPE/'${PROTOTYPE}'/g' | 
        sed -e 's/prototype/'${prototype}'/g' >> $filename
    else
        cat deserializers.tmpl.go |
        sed '1,10d' | # removes go:build ignore, package name and imports from template
        sed -e 's/Prototype/'${Prototype}'/g' | 
        sed -e 's/PROTOTYPE/'${PROTOTYPE}'/g' | 
        sed -e 's/prototype/'${prototype}'/g' >> $filename
    fi

    keep_header_until_imports=false
done
