### Working ###
The code works as follows:
1. firstly I have initialised a few flags (described in HowTo_readme)
2. I have used the .env file to save the source, backup directory paths, also the encryption key
3. then we iterate over the source directory using filepath.WalkDir and pass 
function ittrateOverDir which skips the directories and sends files to the copyFile function
4. The copyFile function opens the source file, checks if we want to do encryption or not and finally saves 
it at the destination.
5. Encryption: I have used ASE encrypt.
6. decrypt: you can decrypt the encrypted file using -D flag and providing the key prompt.

### Current Issues ###
1. if we delete a file in the source Directory, it doesn't get deleted in the backup directory
as I am using copy right now.
2. also currently it is iterating over all the files and copy-pasting them, 
it is not checking whether the file is edited or not.

### Regret ###
1. my code got untidy and scatted at the end when i tried to apply the decrypt and log section.
2. want to refactor it.
