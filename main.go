package main

import (
	flags "Summer_Project_VCS/pkg/flagDeclare"
	copyier "Summer_Project_VCS/pkg/copyFileHelper"
	logger "Summer_Project_VCS/pkg/logger"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)
func ittrateOverDir(path string, d fs.DirEntry, backUpDir string, sourceDir string, encryption bool) {
	if d.IsDir() {
		logger.LogWrier("DIRECTORY DETECTED " + d.Name() + "\n", backUpDir)
		return // skip it
	} else {
		trimmed := strings.TrimPrefix(path, sourceDir)
		copyier.CopyFile(path, backUpDir+trimmed, encryption)
		fmt.Print(".")
		logger.LogWrier("Copying file " + d.Name() + "\n", backUpDir)
	}
}
func main() {
	// declare flags
	backupDir, sourceDir, encryption := flags.DeclareFlags()
	fmt.Print("STARTING\n")
	logger.LogWrier("STARTING\n", backupDir)
	if *encryption {
		logger.LogWrier("ENCRYPTION: ON\n", backupDir)
	}
	// ittrate over the dir
	err := filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Return the error if encountered during traversal
		}
		ittrateOverDir(path, d, backupDir, sourceDir, *encryption)
		return nil
	})
	if err != nil {
		logger.LogWrier("!!!!ERROR!!!!\nimpossible to walk directories: " + err.Error() + "\n\n", backupDir)
		log.Fatalf("!!!!ERROR!!!!\n\nimpossible to walk directories: %s", err)
	}
	print("\nCOMPLETED\n")
	logger.LogWrier("COMPLETED\n\n", backupDir)
}
