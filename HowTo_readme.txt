### Functionality ###
1. at first you need to set up your source and backup directory paths using -cSP, and -cBP, if you don't do so,
the program itself shows you an error and will ask you to do so
2. After the set up when ever you simply go run . it will bydefault copy the 
same directory to the backup_directory.
3. you can use the encryption flag to encrypt the files, -E,
To do that you are required to provide the encryption key which is a 32-char key
like "0123456789abcdef0123456789abcdef"
4. you can any time change the directories and view them my using -Printpaths flag, also the encryption key using -PrintEK.
5. you can decrypt the encrypted file by using -D flag and answaring the required prompts, it will save a copy in your current directory.

### General ###
you can get help any time directly by using go run main.go -h or --help, to know about the flages.
"""
  -D    To decrypt, include -D
  -E    To encrypt and back-it-up, include -E
  -PrintEK
        To print the current set Key, include -PrintEK
  -Printpaths
        To print the current set paths, include -Printpaths
  -cBP string
        To set backUp directory path, include -cBP=
  -cEK string
        To set encryption key, include -cEK=
  -cSP string
        To set source directory path, include -cSP=
"""