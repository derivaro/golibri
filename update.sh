
#!/bin/sh
##gitpush.sh
if [ "$1" = "" ]; then
echo "no message"
else
git add *
git add .
git commit -m "update $1"
git tag $1
git push origin $1

git push origin head
go list -m github.com/derivaro/golibri@$1
echo "done"
fi



#./.register $1