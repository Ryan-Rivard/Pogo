# Pogo
Git repo cli wizard written in Go


git branch
// local branches
git branch -r
// remote branches
git branch -a 
// all branches

git branch -vva


git for-each-ref --sort=committerdate refs/heads/ --format='%(HEAD) %(color:yellow)%(refname:short)%(color:reset) - %(color:red)%(objectname:short)%(color:reset) - %(contents:subject) - %(authorname) (%(color:green)%(committerdate:relative)%(color:reset))'

git for-each-ref --sort=committerdate refs/heads/ --format='%(HEAD) %(refname:short) -  %(contents:subject) -  %(authorname) %(committerdate:relative)'
