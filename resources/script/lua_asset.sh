BASEDIR=$(dirname "$0")
cd $BASEDIR
cd ../..
go generate && go fmt