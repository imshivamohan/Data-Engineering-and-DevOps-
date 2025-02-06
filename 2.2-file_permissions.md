# Day 2: File Permissions and Ownership

## Objective:
Understand file permissions, ownership, and how to manage them in a Linux environment.

---

## 1. Introduction to File Permissions
- Every file and directory in Linux has an associated set of permissions that control who can read, write, and execute it.
- Permissions are assigned to three categories:
  - **Owner (User)**: The creator of the file.
  - **Group**: A set of users who can access the file.
  - **Others**: All other users.

---

## 2. Viewing File Permissions
Use the `ls -l` command to view file permissions:
```bash
ls -l
```
Example output:
```
-rw-r--r-- 1 alice developers 1234 Feb 5 12:34 file.txt
```
Breakdown:
- `-rw-r--r--` → File permissions.
- `alice` → Owner.
- `developers` → Group.
- `1234` → File size.
- `Feb 5 12:34` → Last modified date.

**Understanding the Permission Structure:**
```
-rwxr-xr--
```
Each set of three characters represents **User, Group, and Others**:
- `r` → Read (4)
- `w` → Write (2)
- `x` → Execute (1)

---

## 3. Modifying File Permissions
Permissions can be changed using the `chmod` command.

### Using Numeric Mode:
Each permission type has a numerical value:
- Read (`r`) = 4
- Write (`w`) = 2
- Execute (`x`) = 1

Example:
```bash
chmod 755 script.sh
```
Breakdown:
- `7 (rwx)` → Owner
- `5 (r-x)` → Group
- `5 (r-x)` → Others

### Using Symbolic Mode:
```bash
chmod u+x file.sh   # Give execute permission to the owner
chmod g-w file.sh   # Remove write permission from the group
chmod o+r file.sh   # Give read permission to others
```

---

## 4. Changing File Ownership
The `chown` command is used to change file ownership.
```bash
chown new_user file.txt
chown new_user:new_group file.txt
```
Example:
```bash
sudo chown bob:bob myfile.txt
```

---

## 5. Special Permissions
### SUID (Set User ID)
Allows a file to be executed with the permissions of its owner:
```bash
chmod u+s file
```

### SGID (Set Group ID)
Ensures files created in a directory inherit group ownership:
```bash
chmod g+s directory
```

### Sticky Bit
Prevents users from deleting files they don’t own in shared directories:
```bash
chmod +t directory
```
Example:
```bash
ls -ld /tmp
```
Expected output:
```
drwxrwxrwt 10 root root 4096 Feb 5 /tmp
```

---

## 6. Practical Exercises
1. Create a new file and check its permissions:
   ```bash
   touch myfile.txt
   ls -l myfile.txt
   ```
2. Modify its permissions:
   ```bash
   chmod 644 myfile.txt
   ls -l myfile.txt
   ```
3. Change ownership:
   ```bash
   sudo chown bob:bob myfile.txt
   ```
4. Set special permissions and verify:
   ```bash
   chmod +t /shared_directory
   ls -ld /shared_directory
   ```

---

## 7. Summary & Best Practices
- Always use the least amount of permissions necessary for security.
- Use `chmod` to adjust access rights appropriately.
- Use `chown` to manage file ownership.
- Understand special permissions to control user access in shared environments.

**End of Day 2 Lesson.**

