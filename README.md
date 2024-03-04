# blockchain-quiz-2

# Go-Block-Manager

This is a simple Go project for managing blocks, with functions to display all blocks, create a new block, and modify an existing block.

## Project Structure

The project has the following structure:

- `main.go`: Main source file containing the implementation of block functions.
- `README.md`: Documentation for the project.
- `go.mod`: Go module file for dependency management.

## Functions

### 1. DisplayAllBlocks

Prints all the blocks in the database.

### 2. NewBlock

Creates a new block and adds it to the database.

### 3. ModifyBlock

Updates the data of an existing block in the database.

## Branches

The project has two branches:

- `main`: The main branch.
- `dev`: The development branch.

## Commands Used

### GitHub Initialization

```bash
# Create a new directory for your project
mkdir your_project_name
cd your_project_name

# Create the main.go file with the source code

# Initialize the Go module
go mod init your_module_name


# Initialize a new Git repository
git init

# Add the files to the repository
git add .

# Commit the changes
git commit -m "Initial commit"

# Create a new branch (Dev)
git branch dev

# Switch to the Dev branch
git checkout dev

# Make changes and commit in the Dev branch

# Switch back to the main branch
git checkout main

# Merge the changes from Dev to main
git merge dev

# Push the changes to GitHub
git remote add origin your_github_repo_url
git push -u origin main


# Run the program
go run main.go



Please replace the placeholders with your actual project information before using this README.md file. After creating the README, you can commit and push it to your GitHub repository.
