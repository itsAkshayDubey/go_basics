export GO111MODULE=auto
dir=`pwd | sed 's#.*/##'`
if [[ $dir == 'scripts' ]]
then
        cd ../..
fi

go run main.go deck.go
