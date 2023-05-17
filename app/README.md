
# Password Generator

This application is a simple yet powerful password generator built using Wails and React. It allows the user to generate a password of a specified length, including a mix of uppercase letters, lowercase letters, numbers, and special characters.

## Prerequisites

- [Go](https://golang.org/doc/install) >= 1.16
- [Node.js](https://nodejs.org/en/download/) >= 14
- [Wails](https://wails.app/home/) >= v2

## Installation

1. Clone the repository
```bash
git clone https://github.com/eestaniel/password_generator2.git
cd password-generator
```
2. Install the dependencies
```bash
wails setup
cd frontend
npm install
cd ..
```
3. Build the application
```bash
wails build
```
## Run the application

After building the application, you can find the executable in the `build` directory. To run the application, navigate to the `build` directory and run the executable file.

On Linux or Mac:

```bash
cd build
./password-generator
rve
```

## Usage

When you open the application, you will see a form where you can select the types of characters to include in the password and the length of the password. After selecting your preferences, click the "Generate Password" button to generate a password. The generated password will appear in the text box at the top of the form. You can then copy the password to the clipboard by clicking the "Copy" button next to the password.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
