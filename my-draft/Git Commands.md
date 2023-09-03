# Git Commands

### Init

```
git init .
```

### Remote

Add remote

```
git remote add origin https://<token>@github.com/<user>/<repo>.git
```

Remove remote

```
git remote remove origin
```

### Add

```
git add .
```

File exclusion

```
git add . -- ':!.env' ':!config.yml'
```

### Commit

```
git commit -m 'commit message here...'
```

### Push

```
git push origin master
```

Force push (if previous `.git/` removed)

```
git push --force --set-upstream origin master
```

### Pull

```
git pull origin master
```

### Clone

```
git clone https://<token>@github.com/<user>/<repo>.git
```

### Reset

```
git reset --hard
```
