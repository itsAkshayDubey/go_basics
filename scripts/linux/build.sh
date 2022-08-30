dir=`pwd | sed 's#.*/##'`
if [[ $dir == 'scripts' ]]
then
        cd ../..
fi

./scripts/test.sh
go build main.go deck.go
