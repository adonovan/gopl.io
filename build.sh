#!/bin/bash

function do_Test(){

    #echo "gopl.io/$1/$2"
    cd ../..
    go get ./$1/$2
    ret_code=$?
    cd -

    if [[ $ret_code -ne 0 ]]; then
        echo "ERROR ($ret_code): $1/$2"
        echo
        Total_Erros=$((Total_Erros + 1))
    fi
    Total_Tests=$((Total_Tests + 1))
}

#-----------   Init   -----------
Total_Tests=0
Total_Erros=0

#cd src/gopl.io
#export GOPATH=$(pwd)

for chapter in *; do
    if [[ -d $chapter ]]; then
        echo "-- Compiling examples of $chapter..."
        cd $chapter

        for topic in *; do
            if [[ -d $topic ]]; then
                echo "  ==> building $topic..."
                cd $topic
                do_Test $chapter $topic
            fi
            cd ..
        done
        cd ..
    fi
done

#----------- The End -----------
echo
if [ $Total_Erros -eq 0 ]; then
    echo '--------------------------------------------------------------------------------'
    echo "All tests passed without error! Performed $Total_Tests tests in $SECONDS seconds."
    echo '--------------------------------------------------------------------------------'
    exit 0
else
    echo '--------------------------------------------------------------------------------'
    echo "$Total_Erros tests FAILED! Performed $Total_Tests tests in $SECONDS seconds."
    echo '--------------------------------------------------------------------------------'
    exit 1
fi

