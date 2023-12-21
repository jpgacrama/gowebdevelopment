# Self-Study: Go Web Development with Angular

Welcome to my self-study repository focusing on Go Web Development with Angular. This project is a compilation of examples and source codes derived from [Packt Publishing's Go Web Development Cookbook](https://github.com/PacktPublishing/Go-Web-Development-Cookbook). The aim is to facilitate full-stack development using Go and Angular.

## Table of Contents

- [Self-Study: Go Web Development with Angular](#self-study-go-web-development-with-angular)
  - [Table of Contents](#table-of-contents)
  - [Self-Study Overview](#self-study-overview)
  - [Source Code Examples](#source-code-examples)
    - [Initializing a Go Module](#initializing-a-go-module)
    - [Fixing Workspace Errors in VSCode](#fixing-workspace-errors-in-vscode)

## Self-Study Overview

This repository serves as a personal exploration of Go Web Development with a focus on integrating Angular.

## Source Code Examples

The source codes and examples are adapted from the [Go Web Development Cookbook](https://github.com/PacktPublishing/Go-Web-Development-Cookbook) provided by Packt Publishing.

### Initializing a Go Module

To ensure proper module support for your Go project, follow these steps:

1. Navigate to your project's root directory using the terminal.
2. Run the following command to initialize a Go module:

    ```bash
    go mod init your-module-name
    ```

   Replace "your-module-name" with a suitable module name for your project.

### Fixing Workspace Errors in VSCode

If you encounter workspace errors related to gopls while using VSCode, follow these steps:

1. Open the workspace settings (`.vscode/settings.json`).
2. Ensure that the "gopls" configuration is present:

    ```json
    {
      "gopls": {
        "usePlaceholders": true,
        "experimentalWorkspaceModule": true
      }
    }
    ```

3. Restart VSCode after making these changes.

For more details on setting up your workspace with gopls, refer to the [gopls Workspace Setup documentation](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md).

Feel free to explore the examples and source codes for a hands-on experience in Go Web Development with Angular! I use VSCode as my editor of choice for this project.
