---
title: Git Branch Operations & Commands
---

# Git Branch Operations & Commands

## 1. **Viewing Branches**
- List all local branches:
  ```sh
  git branch
  ```
- List all remote branches:
  ```sh
  git branch -r
  ```
- List both local and remote branches:
  ```sh
  git branch -a
  ```

## 2. **Creating a Branch**
- Create a new branch:
  ```sh
  git branch <branch-name>
  ```
- Create and switch to the new branch:
  ```sh
  git checkout -b <branch-name>
  ```
- Create a new branch from a specific commit:
  ```sh
  git branch <branch-name> <commit-hash>
  ```

## 3. **Switching Between Branches**
- Switch to another branch:
  ```sh
  git checkout <branch-name>
  ```
- Switch using `git switch` (newer alternative to `checkout`):
  ```sh
  git switch <branch-name>
  ```
- Switch to the previous branch:
  ```sh
  git switch -
  ```

## 4. **Renaming a Branch**
- Rename the current branch:
  ```sh
  git branch -m <new-branch-name>
  ```
- Rename another branch:
  ```sh
  git branch -m <old-branch-name> <new-branch-name>
  ```

## 5. **Deleting a Branch**
- Delete a local branch:
  ```sh
  git branch -d <branch-name>
  ```
- Force delete a local branch:
  ```sh
  git branch -D <branch-name>
  ```
- Delete a remote branch:
  ```sh
  git push origin --delete <branch-name>
  ```

## 6. **Merging Branches**
- Merge another branch into the current branch:
  ```sh
  git merge <branch-name>
  ```
- Abort a conflicted merge:
  ```sh
  git merge --abort
  ```

## 7. **Rebasing Branches**
- Rebase the current branch onto another branch:
  ```sh
  git rebase <branch-name>
  ```
- Abort a rebase:
  ```sh
  git rebase --abort
  ```
- Continue a rebase after resolving conflicts:
  ```sh
  git rebase --continue
  ```

## 8. **Tracking Remote Branches**
- Set up a local branch to track a remote branch:
  ```sh
  git branch --track <branch-name> origin/<branch-name>
  ```
- Set the upstream branch for the current branch:
  ```sh
  git branch --set-upstream-to=origin/<branch-name>
  ```

## 9. **Fetching and Pulling Changes**
- Fetch all branches from remote:
  ```sh
  git fetch --all
  ```
- Pull latest changes for a specific branch:
  ```sh
  git pull origin <branch-name>
  ```

## 10. **Pushing a Branch to Remote**
- Push a new branch to remote:
  ```sh
  git push -u origin <branch-name>
  ```
- Push an existing branch:
  ```sh
  git push origin <branch-name>
  ```

## 11. **Stashing Changes (Before Switching Branches)**
- Save uncommitted changes:
  ```sh
  git stash
  ```
- Apply the latest stash:
  ```sh
  git stash apply
  ```
- List all stashed changes:
  ```sh
  git stash list
  ```
- Drop the latest stash:
  ```sh
  git stash drop
  ```

## 12. **Comparing Branches**
- Show differences between two branches:
  ```sh
  git diff <branch1>..<branch2>
  ```
- Show which commits are in one branch but not another:
  ```sh
  git log <branch1>..<branch2>
  ```

## 13. **Resetting and Reverting**
- Reset branch to a previous commit (hard reset):
  ```sh
  git reset --hard <commit-hash>
  ```
- Undo the last commit while keeping changes:
  ```sh
  git reset --soft HEAD~1
  ```
- Revert a specific commit:
  ```sh
  git revert <commit-hash>
  ```

## 14. **Cleaning Up Stale Branches**
- Remove local branches that no longer exist on remote:
  ```sh
  git fetch -p
  ```
- List branches that are merged into `main`:
  ```sh
  git branch --merged main
  ```
- Delete merged branches:
  ```sh
  git branch --merged | grep -v "\*" | xargs -n 1 git branch -d
  ```

These commands cover most of the essential Git branch operations. Let me know if you need more details!

