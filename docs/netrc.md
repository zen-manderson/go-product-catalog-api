# Configuring a .netrc file

In order to access private tooling packages within zenbusiness, Go will need to be able to privately connect to Github to satisfy the `go.mod` and `go.sum` requirements.

If you have been developing at ZenBusiness already, chances are you have already done this. You will be able to recover your token by opening your `.npmrc` file, which should be
available in your terminal root.

```shell
# Open your .npmrc file with a code editor
vim ~/.npmrc 
#nano ~/.npmrc 
#code ~/npmrc 
```

In this file, you'll have a key that will look like:
`//npm.pkg.github.com/:_authToken=ghp_....`

Copy the entire `ghp_` value.

## Setting a personal access token in .netrc
Create a `.netrc` file in the root of this project and apply the following config:

```text
machine github.com
login <your_github_username>
password <your_access_token>
```

### Generating a personal access token
If you have not generated an access token, instructions can be found in the company Onboarding Guide:
[Web Development Onboarding Guide](https://counsl.atlassian.net/wiki/spaces/ENG/pages/1668513807/Web+Development+Onboarding+Guide)
