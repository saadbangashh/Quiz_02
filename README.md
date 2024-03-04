# Quiz_02

# Simple Blockchain in Go
## Project Overview
This project demonstrates the implementation of a basic blockchain structure using Go (Golang). It includes functionalities to display all blocks within the blockchain, add a new block, and modify existing blocks. The project is structured to practice and showcase the use of Git and GitHub commands, following best practices for version control.

## Key Features
### DisplayAllBlocks: 
Function to display all the blocks in the blockchain.
### NewBlock: 
Function to create and add a new block to the blockchain.
### ModifyBlock: 
(Hypothetical) Function to demonstrate how a block could be modified (Note: In a real blockchain, blocks should be immutable).

## Project Structure
The project is divided into packages to organize the code effectively, ensuring clarity and modularity.

### main package: 
Contains the application entry point.

### blockchain package: 
Encapsulates the blockchain logic.

## Branching Strategy
The repository contains two main branches:
main: Serves as the primary branch where the final version of the code is maintained.
dev: Used for development purposes, where new features and fixes are implemented before being merged into the main branch.

## Git Commands Used
Below is a list of Git commands used in this project, serving as a practical reference:
### Clone the repository
- git clone git@github.com:saadbangashh/Quiz_02.git

### Create and switch to the dev branch
- git checkout -b dev

### Add changes to staging
- git add . 

### Commit changes
- git commit -m "Developing functions in GoLang"

### Push the dev branch to GitHub
- git push origin dev

### Switch to the main branch
- git checkout main

### Merge the dev branch into main
- git merge dev

### Push the merged changes to GitHub
- git push origin main
