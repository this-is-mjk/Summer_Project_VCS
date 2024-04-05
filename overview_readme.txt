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
7. in the pkg folder you can find the flagDeclare module this declares the flags at start and a few functions
like decryption and set up of other env variables
8. for setting up the env variables i have used the envHelper module it have SetEnv and GetEnv functions.
9. for encryption and decryption i have the encrypterHelper module
10. for logging the activities to the log file in the formed backup i have used the logger module
11. you can try it over exampledir

### Current Issues ###
1. if we delete a file in the source Directory, it doesn't get deleted in the backup directory
as I am using copy right now.
2. also currently it is iterating over all the files and copy-pasting them, 
it is not checking whether the file is edited or not.
