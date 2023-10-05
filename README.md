# CredCraft

CredCraft is a powerful credential generation program that allows you to easily generate default credentials for various systems and services.

## Features

- Generate default credentials for a wide range of systems and services.
- Customize the generated credentials to match your requirements.
- Simple and intuitive command-line interface for ease of use.
- Flexible and extensible architecture to support future credential generation needs.

## Installation

To install CredCraft, you need to have Go installed on your system. Once you have Go installed, you can use the following command to install CredCraft directly:

```shell
go install github.com/tizisec/CredCraft@latest
```

Usage
CredCraft expects two input files: a wordlist file containing a list of names and a template file specifying the credential format. The program replaces placeholders in the template with names from the wordlist to generate default credentials.

shell
Copy
CredCraft -w <path-to-wordlist> -t <path-to-templates>
-w <path-to-wordlist>: Path to the file containing the list of names.
-t <path-to-templates>: Path to the file containing the credential templates.
Examples
To illustrate the usage of CredCraft, here are a few examples:

Generate default credentials for a system using a predefined template file:

shell
Copy
CredCraft -w names.txt -t template.txt


Customize the template format by using the placeholder "{}" to represent the name:
```
Copy
username: {}
password: default123
```

The program will replace "{}" with each name from the wordlist file to generate the credentials.

Contributing
Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

#License
This project is licensed under the MIT License.
