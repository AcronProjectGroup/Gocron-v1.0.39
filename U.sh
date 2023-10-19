#!/bin/bash
echo -n "Git push on branch develop "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
git add .
git commit -m " ++ "
git push -u origin develop
echo ""
echo "====================================================================="


git checkout beforeMergeToMain
echo -n "git checkout beforeMergeToMain "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
echo "====================================================================="


git merge develop
echo -n "git merge develop "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
echo "====================================================================="


echo -n "Git is Push for beforeMergeToMain "
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
git add .
git commit -m " ++ "
git push -u origin beforeMergeToMain
echo ""
echo "====================================================================="


git checkout develop
echo -n "git checkout develop"
for i in {1..3}
do
    echo -n ". "
    sleep 1
done
echo ""
echo "====================================================================="



