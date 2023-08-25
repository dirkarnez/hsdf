hsdf
====
Hierarchical Structural Data Format. An experiment to see if file system (or **any other external network storages**) can work like a [The HDF5® Library & File Format - The HDF Group](https://www.hdfgroup.org/solutions/hdf5/) by wrapping IO operations to emulate random file access.

### TODOs
- [ ] use go-git
  - ```go
    package main
    
    import (
    	"fmt"
    	"io/ioutil"
    
    	"gopkg.in/src-d/go-git.v4"
    	"gopkg.in/src-d/go-git.v4/plumbing"
    	"gopkg.in/src-d/go-git.v4/storage"
    	"gopkg.in/src-d/go-git.v4/storage/filesystem"
    )
    
    func main() {
    	// Set the path to the physical repository directory on disk
    	repoPath := "/path/to/repository"
    
    	// Initialize a new repository with the physical storage
    	repo, err := git.PlainInit(repoPath, false)
    	if err != nil {
    		fmt.Println("Error initializing repository:", err)
    		return
    	}
    
    	// Get a reference to the virtual file system
    	fs := repo.Worktree.Filesystem
    
    	// Start a transaction
    	tx, err := fs.Begin()
    	if err != nil {
    		fmt.Println("Error starting transaction:", err)
    		return
    	}
    	defer tx.Rollback()
    
    	// Perform file operations within the transaction
    	filePath := "path/to/file.txt"
    	content := []byte("Hello, world!")
    
    	err = tx.WriteFile(filePath, content, 0644)
    	if err != nil {
    		fmt.Println("Error writing file:", err)
    		return
    	}
    
    	// Commit the transaction
    	err = tx.Commit()
    	if err != nil {
    		fmt.Println("Error committing transaction:", err)
    		return
    	}
    
    	// Retrieve the commit hash of the latest commit
    	ref, err := repo.Head()
    	if err != nil {
    		fmt.Println("Error retrieving HEAD reference:", err)
    		return
    	}
    	commitHash := ref.Hash().String()
    
    	fmt.Println("File operation completed successfully")
    	fmt.Println("Latest commit hash:", commitHash)
    }
    ```

### Why
- Simplest way to store data
- Textual and binary data are not mixed
- Source-control friendly
- File locking
- parallel processing
