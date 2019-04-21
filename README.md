# GoCourse
git log
git show [ID коммита]

git branch - Показать локальные ветки
git branch -a - Показать локальные и удалённые ветки
git branch love - Создаём локальную ветку "love"
git checkout love - Переключаемся на ветку "love"
echo 'I love you so much!' > Love.txt
git status
git add Love.txt
git status
git commit -m "Start love feature"
git push origin love
git checkout master
git merge love
git push origin

Удаление веток
git branch -d love - Удаляем локальную ветку
git push origin --delete love - Удаляем ветку из репозитория
