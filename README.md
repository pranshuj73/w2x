# w2x: Write to X

`w2x` is a Go utility that allows you to tweet directly from your terminal. This tool simplifies the process of posting tweets by providing a command-line interface.

## Features

- Post tweets from your terminal.
- Appear productive to others (???).

## Installation

To install `w2x`, you need to have Go installed. Once Go is installed, you can install `w2x` using the following command:

```sh
go install github.com/pranshuj73/w2x@latest
```

## Setup

### Step 1: Sign up for a Twitter Developer Account

1. Go to the [Twitter Developer](https://developer.twitter.com/en/apply-for-access) website and sign up for a developer account.
2. Once your account is approved, log in to the [Twitter Developer](https://developer.twitter.com/en/portal/dashboard) portal.

### Step 2: Create a Twitter App

1. In the Twitter Developer portal, click on "Projects & Apps" and then "Create App".
2. Fill in the required details for your app.
3. After creating the app, go to the "Keys and tokens" tab.
4. Generate and copy the following credentials:
   - Consumer Key
   - Consumer Key Secret
   - Access Token
   - Access Token Secret

### Step 3: Configure `w2x`

Running w2x will create a configuration file at `~/.config/w2x.json` and prompt for the credentials we obtained in the previous step.

You can optionally the config file on your own or modify it anytime you wish.

Sample config file:
```
{
  "consumer_key": "your-consumer-key",
  "consumer_secret": "your-consumer-secret",
  "access_token": "your-access-token",
  "access_secret": "your-access-secret"
}
```

## Usage

To post a tweet using `w2x`, run the following command:

```sh
w2x
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or bug fixes.

## License

This project is licensed under the MIT License.
