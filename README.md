# Config Helper

**Config Helper** is a Go application designed to automate remote configuration tasks. It supports SSH connections to remote hosts and can execute various types of tasks such as retrieving system facts and managing directories.

## Features

- SSH connectivity to remote hosts.
- Retrieving system facts via SSH commands.
- Managing directories, files, and network configurations.
- Configurable via YAML files.
- Command-line interface for specifying configuration files.

## Installation

1. **Clone the Repository:**

   ```sh
   git clone https://github.com/yourusername/config-helper.git
   cd config-helper
   ```

2. **Build the Project:**

   Ensure you have Go installed. Then build the project using:

   ```sh
   go build -o config-helper main.go
   ```

   This will create an executable named `config-helper`.

## Usage

Run the application with the `-f` flag to specify the path to your configuration file:

```sh
./config-helper -f /path/to/your/config.yaml
```

### Configuration File

The configuration file is a YAML file that specifies the remote host, facts to retrieve, and tasks to execute. Here is an example configuration file:

```yaml
host:
  host: "example.com:2222"
  user: "user"
  keyPath: "/path/to/private/key"

facts:
  commands:
    - "cat /proc/meminfo | grep MemTotal"

tasks:
  - category: "dirs"
    type: "ensureDir"
    parameters:
      path: "/tmp/mydir"
      owner: "user"
      mode: "0755"
  - category: "files"
    type: "lineInFile"
    parameters:
      filePath: "/etc/example.conf"
      line: "new_configuration_option=value"
  - category: "files"
    type: "replace"
    parameters:
      filePath: "/etc/example.conf"
      oldPattern: "old_option"
      newPattern: "new_option"
```

### Task Types

Tasks are categorized and implemented as follows:

- **Dirs**: Manage directories on the remote host.
  - `ensureDir`: Ensures a directory exists with the specified owner and mode.

- **Files**: Manage files on the remote host.
  - `lineInFile`: Adds a line to a file.
  - `replace`: Replaces text patterns in a file.

- **Networking**: (Not yet implemented)

## Development

1. **Install Dependencies:**

   Ensure all dependencies are installed:

   ```sh
   go mod tidy
   ```

2. **Run Tests:**

   Run unit tests to ensure everything is working correctly:

   ```sh
   go test ./...
   ```

3. **Contribute:**

   Feel free to open issues or submit pull requests to contribute to the project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions or feedback, please reach out to:

- **Author:** Kevin Yu
- **Email:** kevin.yu@example.com