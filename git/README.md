<h2>Git <img width="30" height="30" alt="image" src="https://github.com/user-attachments/assets/74aff18d-e9b3-4f7f-922f-5b2db8f2c22f" />
 and Github <img width="30" height="30" alt="image" src="https://github.com/user-attachments/assets/1a452767-ab98-4839-ac97-130e30bdb1fa" />
</h2>

This repository was created primarily for **personal learning purposes** — a place to apply and reinforce my understanding of Git and GitHub.

It contains helpful notes and frequently used Git commands that I’ve found useful during my learning journey.

- git is version control
- like a time machine for code
- saving progress or checkpoints
- track changes
- go back for old versions
- work in teams
- Github : store and share git projects on cloud + collaborate with others
- repo : project folder tracked by git - contains code, docs, history of changes

### Set name and email

```bash
git config --global user.email "neelkumar247me@gmail.com"
```

> configured for global
> 

```bash
git config --global user.name "Neel Kumar"
```
---
<br><br>
### Initialise git repo
```bash
git init
```
---
<br><br>

### Checking status
```bash
git status
```

output :

```bash
On branch main

No commits yet

nothing to commit (create/copy files and use "git add" to track)
```

- .git folder created - hidden folder
- we are on “main” branch

<img height="306" alt="image" src="https://github.com/user-attachments/assets/931717e3-c9b3-4fce-b5bd-2bd5f52fbfcb" />
---
<br><br>

### Checking current branch

```bash
git branch
```

nothing happens!

as there is no commit

> create file readme.md
> 

Making initial commit

```bash
git add .
git commit -m "initial commit"
```

now run

```bash
git branch
```

output :

```bash
* main
```
---
<br><br>
### Git insiders

- .git - tracks everything we do in our repo

```bash
working directory         [untracked]
			 👇
staging area (snapshot)   [staged]
			 👇
		repository            [commited]
```

```bash
git status
```

output :

```bash
On branch main
nothing to commit, working tree clean
```
---
<br><br>
### Creating files

- data.txt
- new_data.txt

using terminal to create them :

```bash
echo "This is some data of some txt file" > data.txt
```
---
<br><br>
### Moving files to staging area

```bash
git add data.txt
git add new_data.txt
```
---
<br><br>
### Committing changes

```bash
git commit -m "create data.txt"
```

> Tip: write commit msgs in simple present tense
> 

To make all untracked files go in staging area

```bash
git add .
```
---
<br><br>
### Logs

- details, info or meta data
- you get more info from logs
- tracking info

```bash
git log
```

output :

```bash
commit 1c5699a673bfd158c6de8e3596fd157b0fd21031 (HEAD -> main)
Author: Neel Kumar <neelkumar247me@gmail.com>
Date:   Sat Oct 4 07:11:44 2025 +0530

    Add myfile.txt and new_data.txt

commit ffabe2b18d95c4a74709451bd0a2c30c31e9b5d6
Author: Neel Kumar <neelkumar247me@gmail.com>
Date:   Sat Oct 4 07:05:59 2025 +0530

    create data.txt file

commit 8c464e57c7a5aa1aad6df3bf0a25748cd4f8efa4
Author: Neel Kumar <neelkumar247me@gmail.com>
Date:   Sat Oct 4 05:52:27 2025 +0530

    initial commit
```

Getting short log

```bash
git log --oneline
```

output :

```bash
1c5699a (HEAD -> main) Add myfile.txt and new_data.txt
ffabe2b create data.txt file
8c464e5 initial commit
```
---
<br><br>
### Usage of .gitignore file
Created few more files
- .env
- api_secrets.txt

These are hidden files

Need to make sure git don’t track them

Solution :

create new file - .gitignore

Add following data in it,

```bash
.env
api_secrets.txt
```

then do following,

```bash
git add .
git commit -m "add .gitignore"
git status
```
---
<br><br>
### Ignoring folders in our repo

Folder inside repo :

```bash
details
		|____ bank.txt
		|____ password.txt
```

add following in .gitignore

```bash
details/
```
---
<br><br>
### Tracking empty folders for later use

Git does not track empty folders

To make them track,

make folder `mails` and add `.gitkeep`

```bash
mails
	|____ .gitkeep
```
---
<br><br>
### Branches

<img height="270" alt="image" src="https://github.com/user-attachments/assets/6618f2ec-7c11-4b13-8a80-5d8e756e41eb" />


Viewing branches

```bash
git branch
```

Creating new branch

- Method 1

```bash
git checkout -b newBranch
```

- Method 2

```bash
git switch -c newBranch
```

Now checking branch:

```bash
git branch
```

output :

```bash
  main
* newBranch
```

switching back to other branch

```bash
git branch main
```

> when creating any branch, it is a good practice to be in main branch and create new branch from there
> 

Merging git branches in main

```bash
git merge newBranch -m "merge newBranch"
```
---
<br><br>
### Rebase
<img height="602" alt="image" src="https://github.com/user-attachments/assets/f70ec9c8-0f56-4645-ad24-8606889a5a78" />


- First in both branch 1, branch 2 - we do some changes and commit them
- Then, go in branch 2

```bash
git rebase main
```

go in main

```bash
git switch main
```

Finally,

```bash
git merge branch2
```

### Using reflog : a better version of log

```bash
git reflog
```

output :

```bash
73139f9 (HEAD -> main) HEAD@{0}: commit: delete file
6adbd88 (newRebaseBranch) HEAD@{1}: merge newRebaseBranch: Fast-forward
2e08a23 HEAD@{2}: checkout: moving from newRebaseBranch to main
6adbd88 (newRebaseBranch) HEAD@{3}: rebase (finish): returning to refs/heads/newRebaseBranch
6adbd88 (newRebaseBranch) HEAD@{4}: rebase (pick): add email in email.txt
2e08a23 HEAD@{5}: rebase (start): checkout main
2101a7f HEAD@{6}: checkout: moving from main to newRebaseBranch
2e08a23 HEAD@{7}: commit: add id of neel
49e6be2 HEAD@{8}: checkout: moving from newRebaseBranch to main
2101a7f HEAD@{9}: commit: add email in email.txt
49e6be2 HEAD@{10}: checkout: moving from main to newRebaseBranch
49e6be2 HEAD@{11}: commit: add account folder
c7575d0 HEAD@{12}: checkout: moving from rebaseFeature to main
9ba39cb (rebaseFeature) HEAD@{13}: rebase (finish): returning to refs/heads/rebaseFeature
9ba39cb (rebaseFeature) HEAD@{14}: rebase (continue): modify with rebaseBranch
c7575d0 HEAD@{15}: rebase (start): checkout main
742c767 HEAD@{16}: checkout: moving from main to rebaseFeature
c7575d0 HEAD@{17}: commit: modfiy with main branch
00d9fd3 HEAD@{18}: checkout: moving from rebaseFeature to main
742c767 HEAD@{19}: commit: modify with rebaseBranch
00d9fd3 HEAD@{20}: checkout: moving from main to rebaseFeature
00d9fd3 HEAD@{21}: merge bugfix: Merge made by the 'ort' strategy.
e7e8b04 HEAD@{22}: commit: bug fixed from main
```

- head0 is current commit
- head 1 is previous commit and so on
---
<br><br>
### Bringing back deleted file 

```bash
git reset --hard HEAD~1
```

But the above one is not recommended way

⇒ we should always use commit id

```bash
git reset --hard 6adbd88
```

> Put id of that head here after which you delete the file
> 
---
<br><br>
### Using git diff to check our changes

```bash
git diff --staged
```

output :

```bash
diff --git a/readme.md b/readme.md
index e69de29..a66477c 100644
--- a/readme.md
+++ b/readme.md
@@ -0,0 +1,2 @@
+This is readme.md file
+It describes the whole repo!
\ No newline at end of file
```

To undo whatever we add in staging area :

```bash
git reset
```
---
<br><br>
### Cherry picking

⇒ When you want to merge a specific commit of a branch to main branch

<img height="208" alt="image" src="https://github.com/user-attachments/assets/40aaea46-97ac-49d9-b1db-98b33b7bf9d5" />

Made 3 commits with cherryBranch

Now cherry pick 1st commit - “update readme.md”

First get commit id :

```bash
git switch cherryBranch
git log --oneline
```

output :

```bash
90ae239 (HEAD -> cherryBranch) update id.txt
cc102f7 add email in email.txt
06621a1 update readme.md
```

here, commit id = 06621a1

now do these commands

```bash
git switch main
git cherry-pick 06621a1
```

output :

```bash
[main 0506483] update readme.md
 Date: Sat Oct 4 21:22:49 2025 +0530
 1 file changed, 3 insertions(+), 1 deletion(-)
```

### Stashing

- you are in some branch and making some changes
- Now you want to switch your branch but you have not saved your changes
- Then there are 2 ways to do so
    1. commit all changes - but if you want to test them later and not now, it is not good option
    2. You push the changes in stash

Stash - a temporary memory to store changes

If we try switching without saving changes, we will get this :

```bash
error: Your local changes to the following files would be overwritten by checkout:
	development.txt
Please commit your changes or stash them before you switch branches.
Aborting
```

stashing changes :

```bash
git stash push -m "push to stash"
```

Getting stash info :

```bash
git stash list
```

output :

```bash
stash@{0}: On dev1: push to stash
```

bring back changes from stash :

```bash
git stash apply
```

Now all changes will be restored 

But things in stash is still there, we need too remove them :

```bash
git stash pop
```
---
<br><br>
### Working with github
Renaming branch :

```bash
git branch -m newBranchName
```

adding remote origin :

```bash
git remote add origin "https://github.com/neelkumar01/learning-git.git"
```

checking origin :

```bash
git remote -v
```

output :

```bash
origin	https://github.com/neelkumar01/learning-git.git (fetch)
origin	https://github.com/neelkumar01/learning-git.git (push)
```

to remove origin :

```bash
git remote remove origin
```

pushing code

```bash
git push origin main
```

cloning repo

```bash
git clone "https://github.com/neelkumar01/Learning-Git.git"
```

now use cd to go in cloned repo

advisable to create new branch to make changes
