#!/bin/sh
# ref: https://help.github.com/articles/adding-an-existing-project-to-github-using-the-command-line/
#
# Usage example: /bin/sh ./git_push.sh wing328 openapi-petstore-perl "minor update" "gitlab.com"

m-den-i=$1
go-firefly-iii-client=$2
release_note=$3
git_host=$4

if [ "$git_host" = "" ]; then
    git_host="github.com"
    echo "[INFO] No command line input provided. Set \$git_host to $git_host"
fi

if [ "$m-den-i" = "" ]; then
    m-den-i="m-den-i"
    echo "[INFO] No command line input provided. Set \$m-den-i to $m-den-i"
fi

if [ "$go-firefly-iii-client" = "" ]; then
    go-firefly-iii-client="go-firefly-iii-client"
    echo "[INFO] No command line input provided. Set \$go-firefly-iii-client to $go-firefly-iii-client"
fi

if [ "$release_note" = "" ]; then
    release_note="Minor update"
    echo "[INFO] No command line input provided. Set \$release_note to $release_note"
fi

# Initialize the local directory as a Git repository
git init

# Adds the files in the local repository and stages them for commit.
git add .

# Commits the tracked changes and prepares them to be pushed to a remote repository.
git commit -m "$release_note"

# Sets the new remote
git_remote=$(git remote)
if [ "$git_remote" = "" ]; then # git remote not defined

    if [ "$GIT_TOKEN" = "" ]; then
        echo "[INFO] \$GIT_TOKEN (environment variable) is not set. Using the git credential in your environment."
        git remote add origin https://${git_host}/${m-den-i}/${go-firefly-iii-client}.git
    else
        git remote add origin https://${m-den-i}:"${GIT_TOKEN}"@${git_host}/${m-den-i}/${go-firefly-iii-client}.git
    fi

fi

git pull origin master

# Pushes (Forces) the changes in the local repository up to the remote repository
echo "Git pushing to https://${git_host}/${m-den-i}/${go-firefly-iii-client}.git"
git push origin master 2>&1 | grep -v 'To https'
