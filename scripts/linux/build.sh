export GO111MODULE=auto
dir=`pwd | sed 's#.*/##'`
if [[ $dir == 'linux' ]]
then
        cd ../..
fi

./scripts/linux/test.sh
go build main.go deck.go
