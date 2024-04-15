# Setting Environment Variables for Development
As with all ZenBusiness applications, environment variables are a necessity. Some are for your applications features,
but some are for internal packages to function properly. This service template utilizes a [package](https://github.com/joho/godotenv)
called `godotenv` to read these environment variables.

# Do not commit Secrets EVER
You can set secrets in env/local.env, but make sure to clear them before committing to main.

Change any environment variables in env/local.env or add your own .env file for dev or other environments. If you choose
to create a new .env file be sure to point the environment variable `ENV_FILE` to the location of your .env file.